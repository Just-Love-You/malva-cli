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
	// 1. Determine output filename: use provided or default
	output := opts.Output
	if output == "" {
		output = fmt.Sprintf("mod_%s", input)
	}

	// 2. Initialize ffmpeg arguments with basic flags and input file
	args := []string{"-hide_banner", "-loglevel", "error", "-i", input}

	// 3. Remove audio track if requested
	if opts.RemoveAudio {
		args = append(args, "-an")
	}

	// 4. Prepare video filters: watermark overlay and resize
	var filterComplex string
	if opts.Watermark != "" {
		// 4.1. Check watermark file existence
		if _, err := os.Stat(opts.Watermark); err != nil {
			return fmt.Errorf("watermark file not found: %s", opts.Watermark)
		}
		var baseFilter string
		if opts.ResizeHeight > 0 || opts.ResizeWidth > 0 {
			if opts.ResizeHeight > 0 {
				baseFilter = fmt.Sprintf("scale=-2:%d", opts.ResizeHeight)
			} else {
				baseFilter = fmt.Sprintf("scale=%d:-2", opts.ResizeWidth)
			}
		} else {
			baseFilter = ""
		}
		if baseFilter != "" {
			filterComplex = fmt.Sprintf("[0:v]%s[v0];[v0][1:v]overlay=10:10", baseFilter)
		} else {
			filterComplex = "[0:v][1:v]overlay=10:10[v0]"
		}
		args = append(args, "-i", opts.Watermark, "-filter_complex", filterComplex)
		args = append(args, "-map", "[v0]")
		if opts.ReplaceAudio == "" {
			args = append(args, "-map", "0:a")
		}
	} else if opts.ResizeHeight > 0 || opts.ResizeWidth > 0 {
		var vf string
		if opts.ResizeHeight > 0 {
			vf = fmt.Sprintf("scale=-2:%d", opts.ResizeHeight)
		} else {
			vf = fmt.Sprintf("scale=%d:-2", opts.ResizeWidth)
		}
		args = append(args, "-vf", vf)
	}

	// 5. Replace audio track if requested
	if opts.ReplaceAudio != "" {
		if _, err := os.Stat(opts.ReplaceAudio); err != nil {
			return fmt.Errorf("audio file not found: %s", opts.ReplaceAudio)
		}
		args = append(args,
			"-i", opts.ReplaceAudio,
			"-map", "0:v", "-map", "2:a",
		)
	}

	// 6. Choose codec and format: copy or re-encode with faststart
	hasFilter := opts.ResizeHeight > 0 || opts.ResizeWidth > 0 || opts.Watermark != ""
	if hasFilter {
		// re-encode video and handle audio
		args = append(args, "-c:v", "libx264")
		if opts.ReplaceAudio != "" {
			args = append(args, "-c:a", "aac")
		} else {
			args = append(args, "-c:a", "copy")
		}
		args = append(args, "-movflags", "+faststart")
	} else {
		// simple stream copy if no filters
		args = append(args, "-c", "copy")
		args = append(args, "-movflags", "+faststart")
	}
	// 7. Specify output file
	args = append(args, output)

	// 8. Execute ffmpeg command
	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
