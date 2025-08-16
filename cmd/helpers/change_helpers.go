package helpers

import (
	"fmt"
	"github.com/WeAreTheSameBlood/malva-cli/cmd/constants"
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
func ProcessChange(
	input string,
	opts ChangeOptions,
) error {
	// 1. Determine output filename: use provided or default
	output := opts.Output
	if output == "" {
		output = fmt.Sprintf(constants.CHANGE_DEFAULT_OUTPUT_NAME_PREFIX, input)
	}

	// 2. Initialize ffmpeg arguments with basic flags and input file
	args := []string{"-hide_banner", "-loglevel", "error", "-i", input}

	// 3. Remove audio track if requested
	if opts.RemoveAudio {
		args = append(args, "-an")
	}

	// 4. Prepare video filters: watermark overlay and resize
	var filterComplex string
	var hasResize bool = opts.ResizeHeight > 0 || opts.ResizeWidth > 0
	var hasWatermark bool = opts.Watermark != ""
	var hasFilter bool = hasResize || hasWatermark

	if hasWatermark {
		// 4.1. Check watermark file existence
		if _, err := os.Stat(opts.Watermark); err != nil {
			return fmt.Errorf("watermark file not found: %s", opts.Watermark)
		}

		// add watermark as second input
		args = append(args, "-i", opts.Watermark)

		if hasResize {
			var baseFilter string
			if opts.ResizeHeight > 0 {
				baseFilter = fmt.Sprintf("scale=-2:%d", opts.ResizeHeight)
			} else {
				baseFilter = fmt.Sprintf("scale=%d:-2", opts.ResizeWidth)
			}
			// scale then overlay, final output labeled [v0]
			filterComplex = fmt.Sprintf("[0:v]%s[scaled];[scaled][1:v]overlay=10:10[v0]", baseFilter)
		} else {
			// just overlay, final output labeled [v0]
			filterComplex = "[0:v][1:v]overlay=10:10[v0]"
		}

		args = append(args, "-filter_complex", filterComplex)

		// map filtered video
		args = append(args, "-map", "[v0]")

		// map audio only if not removed and not replaced
		if opts.ReplaceAudio == "" && !opts.RemoveAudio {
			args = append(args, "-map", "0:a")
		}
	} else if hasResize {
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
		args = append(args, "-i", opts.ReplaceAudio)

		// if HAVE watermark 		--> outputs: 0=video, 1=watermark, 2=replaceAudio
		// if DO NOT HAVE watermark --> outputs: 0=video, 1=replaceAudio
		if hasWatermark {
			args = append(args, "-map", "2:a")
		} else {
			args = append(args, "-map", "1:a")
		}
	}

	// 6. Choose codec and format: copy or re-encode with faststart
	if hasFilter || opts.ReplaceAudio != "" {
		args = append(args, "-c:v", "libx264")
		if opts.ReplaceAudio != "" {
			args = append(args, "-c:a", "aac")
		} else if !opts.RemoveAudio {
			args = append(args, "-c:a", "copy")
		}
	} else {
		args = append(args, "-c", "copy")
	}

	// 7. Specify output file
	args = append(args, output)

	// 8. Execute ffmpeg command
	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
