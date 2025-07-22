package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

// ChangeOptions holds parameters for the 'change' command.
type ChangeOptions struct {
	RemoveAudio  bool
	Watermark    string
	ReplaceAudio string
	Output       string
	ResizeHeight int
	ResizeWidth  int
}

// ProcessChange builds ffmpeg arguments and executes the command based on ChangeOptions.
func ProcessChange(input string, opts ChangeOptions) error {
	// determine output filename
	output := opts.Output
	if output == "" {
		output = fmt.Sprintf("mod_%s", input)
	}

	// base args
	args := []string{"-hide_banner", "-loglevel", "error", "-i", input}

	// remove audio
	if opts.RemoveAudio {
		args = append(args, "-an")
	}

	// resize
	if opts.ResizeHeight > 0 || opts.ResizeWidth > 0 {
		var filter string
		if opts.ResizeHeight > 0 {
			filter = fmt.Sprintf("scale=-2:%d", opts.ResizeHeight)
		} else {
			filter = fmt.Sprintf("scale=%d:-2", opts.ResizeWidth)
		}
		args = append(args, "-vf", filter)
	}

	// watermark
	if opts.Watermark != "" {
		if _, err := os.Stat(opts.Watermark); err != nil {
			return fmt.Errorf("watermark file not found: %s", opts.Watermark)
		}
		args = append(args,
			"-i", opts.Watermark,
			"-filter_complex", fmt.Sprintf("overlay=10:10"),
		)
	}

	// replace audio
	if opts.ReplaceAudio != "" {
		if _, err := os.Stat(opts.ReplaceAudio); err != nil {
			return fmt.Errorf("audio file not found: %s", opts.ReplaceAudio)
		}
		args = append(args,
			"-i", opts.ReplaceAudio,
			"-map", "0:v", "-map", "2:a",
		)
	}

	// final copy and output
	args = append(args, "-c", "copy", output)

	// execute ffmpeg
	cmd := exec.Command("ffmpeg", args...)
	return cmd.Run()
}
