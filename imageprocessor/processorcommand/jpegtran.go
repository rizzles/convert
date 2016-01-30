package processorcommand

import (
	"fmt"
	"path/filepath"
)

func Jpegtran(filename string) (string, error) {
	outfile := fmt.Sprintf("%s_opti", filename)

	args := []string{
		"-copy",
		"all",
		"-optimize",
		"-outfile",
		outfile,
		filename,
	}

	path, _ := filepath.Abs("./imageprocessor/processorcommand/jpegtran")

	err := runProcessorCommand(path, args)
	if err != nil {
		return "", err
	}

	return outfile, nil
}
