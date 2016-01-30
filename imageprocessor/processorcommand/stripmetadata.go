package processorcommand

import (
	"path/filepath"
)

func StripMetadata(filename string) error {
	args := []string{
		"-all=",
		"--icc_profile:all",
		"-overwrite_original",
		filename,
	}

	path, _ := filepath.Abs("./imageprocessor/processorcommand/exiftool")

	err := runProcessorCommand(path, args)
	if err != nil {
		return err
	}

	return nil
}
