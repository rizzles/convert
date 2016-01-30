package processorcommand

import (
	"fmt"
	"path/filepath"

	"../thumbType"
)

func ConvertToJpeg(filename string) (string, error) {
	outfile := fmt.Sprintf("%s_jpg", filename)

	args := []string{
		filename,
		"-flatten",
		"JPEG:" + outfile,
	}

	err := runProcessorCommand(GetExecPath(), args)
	if err != nil {
		return "", err
	}

	return outfile, nil
}

func FixOrientation(filename string) (string, error) {
	outfile := fmt.Sprintf("%s_ort", filename)

	args := []string{
		filename,
		"-auto-orient",
		outfile,
	}
	err := runProcessorCommand(GetExecPath(), args)
	if err != nil {
		return "", err
	}

	return outfile, nil
}

func Quality(filename string, quality int) (string, error) {
	outfile := fmt.Sprintf("%s_q", filename)

	args := []string{
		filename,
		"-quality",
		fmt.Sprintf("%d", quality),
		"-density",
		"72x72",
		outfile,
	}

	err := runProcessorCommand(GetExecPath(), args)
	if err != nil {
		return "", err
	}

	return outfile, nil
}

func ResizePercent(filename string, percent int) (string, error) {
	outfile := fmt.Sprintf("%s_rp", filename)

	args := []string{
		filename,
		"-resize",
		fmt.Sprintf("%d%%", percent),
		outfile,
	}

	err := runProcessorCommand(GetExecPath(), args)
	if err != nil {
		return "", err
	}

	return outfile, nil
}

func SquareThumb(filename, name string, size int, format thumbType.ThumbType) (string, error) {
	outfile := fmt.Sprintf("%s_%s", filename, name)

	args := []string{
		fmt.Sprintf("%s[0]", filename),
		"-quality",
		"94",
		"-resize",
		fmt.Sprintf("%dx%d^", size, size),
		"-gravity",
		"center",
		"-crop",
		fmt.Sprintf("%dx%d+0+0", size, size),
		"-density",
		"72x72",
		"-unsharp",
		"0.5",
		fmt.Sprintf("%s:%s", format.ToString(), outfile),
	}

	err := runProcessorCommand(GetExecPath(), args)
	if err != nil {
		return "", err
	}

	return outfile, nil
}

func Thumb(filename, name string, width, height int, format thumbType.ThumbType) (string, error) {
	outfile := fmt.Sprintf("%s_%s", filename, name)

	args := []string{
		fmt.Sprintf("%s[0]", filename),
		"-quality",
		"83",
		"-resize",
		fmt.Sprintf("%dx%d>", width, height),
		"-density",
		"72x72",
		fmt.Sprintf("%s:%s", format.ToString(), outfile),
	}

	err := runProcessorCommand(GetExecPath(), args)
	if err != nil {
		return "", err
	}

	return outfile, nil
}

func CircleThumb(filename, name string, width int, format thumbType.ThumbType) (string, error) {
	outfile := fmt.Sprintf("%s_%s", filename, name)

	filename, err := SquareThumb(filename, name, width, format)
	if err != nil {
		return "", err
	}

	args := []string{
		"-size",
		fmt.Sprintf("%dx%d", width, width),
		"xc:none",
		"-fill",
		filename,
		"-quality",
		"83",
		"-density",
		"72x72",
		"-draw",
		fmt.Sprintf("circle %d,%d %d,1", width/2, width/2, width/2),
		fmt.Sprintf("PNG:%s", outfile),
	}

	err = runProcessorCommand(GetExecPath(), args)
	if err != nil {
		return "", err
	}

	return outfile, nil
}

func CustomThumb(filename, name string, width, height int, cropGravity string, cropWidth, cropHeight, quality int, format thumbType.ThumbType) (string, error) {
	outfile := fmt.Sprintf("%s_%s", filename, name)

	args := []string{
		fmt.Sprintf("%s[0]", filename),
		"-resize",
		fmt.Sprintf("%dx%d^", width, height),
		"-density",
		"72x72",
	}

	if quality != -1 {
		args = append(args,
			"-quality",
			fmt.Sprintf("%d", quality),
		)
	}

	if cropGravity != "" {
		args = append(args,
			"-gravity",
			fmt.Sprintf("%s", cropGravity),
			"-crop",
			fmt.Sprintf("%dx%d+0+0", cropWidth, cropHeight),
		)
	}

	args = append(args, fmt.Sprintf("%s:%s", format.ToString(), outfile))
	err := runProcessorCommand(GetExecPath(), args)
	if err != nil {
		return "", err
	}

	return outfile, nil
}

func Full(filename string, name string, format thumbType.ThumbType) (string, error) {
	outfile := fmt.Sprintf("%s_%s", filename, name)

	args := []string{
		fmt.Sprintf("%s[0]", filename),
		"-quality",
		"83",
		"-density",
		"72x72",
		fmt.Sprintf("%s:%s", format.ToString(), outfile),
	}

	err := runProcessorCommand(GetExecPath(), args)
	if err != nil {
		return "", err
	}

	return outfile, nil
}

func GetExecPath() string {
	path, _ := filepath.Abs("./imageprocessor/processorcommand/convert")
	return path
}