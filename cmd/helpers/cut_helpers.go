package helpers

import (
	"fmt"
	"github.com/WeAreTheSameBlood/malva-cli/cmd/constants"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

// CutOptions holds parameters for the 'cut' command.
type CutOptions struct {
	Start  string
	Finish string
	Output string
}

var timeRe = regexp.MustCompile(`time=(\d+):(\d+):(\d+\.?\d*)`)

// ProcessCut builds ffmpeg arguments and executes the cut command.
func ProcessCut(input string, opts CutOptions) error {
	// determine output filename
	output := opts.Output
	if output == "" {
		output = fmt.Sprintf(constants.CUT_DEFAULT_OUTPUT_NAME_PREFIX, input)
	}

	// build ffmpeg args
	args := []string{
		"-hide_banner", "-loglevel", "error",
		"-i", input,
		"-ss", opts.Start,
		"-to", opts.Finish,
	}

	args = append(args, "-c", "copy", output)

	// execute
	cmd := exec.Command("ffmpeg", args...)
	return cmd.Run()
}

func parseFFmpegTime(hours, mins, sec string) float64 {
	hh, _ := strconv.Atoi(hours)
	mm, _ := strconv.Atoi(mins)
	ss, _ := strconv.ParseFloat(sec, 64)

	return float64(hh)*3600 + float64(mm)*60 + ss
}
