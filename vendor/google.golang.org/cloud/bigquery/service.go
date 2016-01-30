// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bigquery

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
	bq "google.golang.org/api/bigquery/v2"
)

// service provides an internal abstraction to isolate the generated
// BigQuery API; most of this package uses this interface instead.
// The single implementation, *bigqueryService, contains all the knowledge
// of the generated BigQuery API.
type service interface {
	insertJob(ctx context.Context, job *bq.Job, projectId string) (*Job, error)
	jobStatus(ctx context.Context, projectId, jobID string) (*JobStatus, error)
	readTabledata(ctx context.Context, conf *readTabledataConf) (*readTabledataResult, error)
}

type bigqueryService struct {
	s *bq.Service
}

func newBigqueryService(client *http.Client) (*bigqueryService, error) {
	s, err := bq.New(client)
	if err != nil {
		return nil, fmt.Errorf("constructing bigquery client: %v", err)
	}

	return &bigqueryService{s: s}, nil
}

func (s *bigqueryService) insertJob(ctx context.Context, job *bq.Job, projectID string) (*Job, error) {
	// TODO(mcgreevy): use ctx
	res, err := s.s.Jobs.Insert(projectID, job).Do()
	if err != nil {
		return nil, err
	}
	return &Job{service: s, projectID: projectID, jobID: res.JobReference.JobId}, nil
}

type pagingConf struct {
	pageToken string

	recordsPerRequest    int64
	setRecordsPerRequest bool
}

type readTabledataConf struct {
	projectID, datasetID, tableID string
	paging                        pagingConf
}

type readTabledataResult struct {
	pageToken string
	rows      [][]Value
	totalRows int64
}

func (s *bigqueryService) readTabledata(ctx context.Context, conf *readTabledataConf) (*readTabledataResult, error) {
	list := s.s.Tabledata.List(conf.projectID, conf.datasetID, conf.tableID).
		PageToken(conf.paging.pageToken)

	if conf.paging.setRecordsPerRequest {
		list = list.MaxResults(conf.paging.recordsPerRequest)
	}

	res, err := list.Do()
	if err != nil {
		return nil, err
	}

	var rs [][]Value
	for _, r := range res.Rows {
		rs = append(rs, convertRow(r))
	}

	result := &readTabledataResult{
		pageToken: res.PageToken,
		rows:      rs,
		totalRows: res.TotalRows,
	}
	return result, nil
}

func convertRow(r *bq.TableRow) []Value {
	var values []Value
	for _, cell := range r.F {
		values = append(values, cell.V)
	}
	return values
}

func (s *bigqueryService) jobStatus(ctx context.Context, projectID, jobID string) (*JobStatus, error) {
	// TODO(mcgreevy): use ctx
	res, err := s.s.Jobs.Get(projectID, jobID).Do()
	if err != nil {
		return nil, err
	}
	return jobStatusFromProto(res.Status)
}

var stateMap = map[string]State{"PENDING": Pending, "RUNNING": Running, "DONE": Done}

func jobStatusFromProto(status *bq.JobStatus) (*JobStatus, error) {
	state, ok := stateMap[status.State]
	if !ok {
		return nil, fmt.Errorf("unexpected job state: %v", status.State)
	}

	newStatus := &JobStatus{
		State: state,
		err:   nil,
	}
	if err := errorFromErrorProto(status.ErrorResult); state == Done && err != nil {
		newStatus.err = err
	}

	for _, ep := range status.Errors {
		newStatus.Errors = append(newStatus.Errors, errorFromErrorProto(ep))
	}
	return newStatus, nil
}
