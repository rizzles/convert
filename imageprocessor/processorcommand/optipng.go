package processorcommand

import (
	"fmt"
	"path/filepath"
)

func Optipng(filename string) (string, error) {
	outfile := fmt.Sprintf("%s_opi", filename)

	args := []string{
		"-fix",
		"-out",
		outfile,
		filename,
	}

	path, _ := filepath.Abs("./imageprocessor/processorcommand/optipng")

	err := runProcessorCommand(path, args)
	if err != nil {
		return "", err
	}

	return outfile, nil
}
