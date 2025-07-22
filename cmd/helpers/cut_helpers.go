package helpers

import (
	"fmt"
	"os/exec"
	"time"
)

// CutOptions holds parameters for the 'cut' command.
type CutOptions struct {
	Start  string
	Finish string
	Output string
}

// ProcessCut builds ffmpeg arguments and executes the cut command.
func ProcessCut(input string, opts CutOptions) error {
	// determine output filename
	output := opts.Output
	if output == "" {
		output = fmt.Sprintf("cut_%s", input)
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

// FormatDurationFFMPEG converts a time.Duration to "HH:MM:SS.ms" format for ffmpeg
func FormatDurationFFMPEG(d time.Duration) string {
	totalMillis := d.Milliseconds()
	h := totalMillis / 3600000
	m := (totalMillis % 3600000) / 60000
	s := (totalMillis % 60000) / 1000
	ms := totalMillis % 1000
	return fmt.Sprintf("%02d:%02d:%02d.%03d", h, m, s, ms)
}
