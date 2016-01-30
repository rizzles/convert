package imagestore

import (
	/*
	"io"
	"io/ioutil"
	"log"
	"os"
	*/

	"golang.org/x/net/context"
	//"google.golang.org/cloud/storage"
)

type GCSImageStore struct {
	ctx            context.Context
	bucketName     string
	storeRoot      string
	namePathMapper *NamePathMapper
}

func NewGCSImageStore(ctx context.Context, bucket string, root string, mapper *NamePathMapper) *GCSImageStore {
	return &GCSImageStore{
		ctx:            ctx,
		bucketName:     bucket,
		storeRoot:      root,
		namePathMapper: mapper,
	}
}
/*
func (this *GCSImageStore) Exists(obj *StoreObject) (bool, error) {
	_, err := storage.StatObject(this.ctx, this.bucketName, this.toPath(obj))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *GCSImageStore) Save(src string, obj *StoreObject) (*StoreObject, error) {
	srcFd, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer srcFd.Close()

	data, err := ioutil.ReadAll(srcFd)
	if err != nil {
		log.Printf("error on read file: %s", err)
		return nil, err
	}

	wc := storage.NewWriter(this.ctx, this.bucketName, this.toPath(obj))
	wc.ContentType = obj.MimeType
	if _, err := wc.Write(data); err != nil {
		log.Printf("error on write data: %s", err)
		return nil, err
	}
	if err := wc.Close(); err != nil {
		log.Printf("error on close writer: %s", err)
		return nil, err
	}

	obj.Url = "https://storage.googleapis.com/" + this.bucketName + "/" + this.toPath(obj)
	return obj, nil
}

func (this *GCSImageStore) Get(obj *StoreObject) (io.ReadCloser, error) {
	reader, err := storage.NewReader(this.ctx, this.bucketName, this.toPath(obj))
	if err != nil {
		log.Printf("error on read file: %s", err)
		return nil, err
	}

	return reader, nil
}

func (this *GCSImageStore) String() string {
	return "GCSStore"
}

func (this *GCSImageStore) toPath(obj *StoreObject) string {
	if this.storeRoot != "" {
		return this.storeRoot + "/" + this.namePathMapper.mapToPath(obj)
	}
	return this.namePathMapper.mapToPath(obj)
}
*/