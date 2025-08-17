package tasks

import (
	"fmt"
	"github.com/WeAreTheSameBlood/malva-cli/cmd/constants"
	progress "github.com/WeAreTheSameBlood/malva-cli/cmd/subservices"
	"os"
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
	// Determine output filename: use provided or default
	output := opts.Output
	if output == "" {
		output = fmt.Sprintf(constants.CHANGE_DEFAULT_OUTPUT_NAME_PREFIX, input)
	}

	// Initialize ffmpeg arguments with basic flags and input file
	argsChange := []string{
		"-hide_banner",
		"-loglevel", "info",
		"-progress", "pipe:1",
		"-i", input,
	}

	// Remove audio track if requested
	if opts.RemoveAudio {
		argsChange = append(argsChange, "-an")
	}

	// Prepare video filters and flags: watermark overlay and resize
	var filterComplex string
	var hasResize bool = opts.ResizeHeight > constants.CHANGE_DEFAULT_RESIZE_HEIGHT ||
		opts.ResizeWidth > constants.CHANGE_DEFAULT_RESIZE_WIDTH
	var hasWatermark bool = opts.Watermark != ""
	var hasFilter bool = hasResize || hasWatermark

	if hasWatermark {
		// Check watermark file existence
		if _, err := os.Stat(opts.Watermark); err != nil {
			return fmt.Errorf("watermark file not found: %s", opts.Watermark)
		}

		// add watermark as second input
		argsChange = append(argsChange, "-i", opts.Watermark)

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

		argsChange = append(argsChange, "-filter_complex", filterComplex)

		// map filtered video
		argsChange = append(argsChange, "-map", "[v0]")

		// map audio only if not removed and not replaced
		if opts.ReplaceAudio == "" && !opts.RemoveAudio {
			argsChange = append(argsChange, "-map", "0:a")
		}
	} else if hasResize {
		var vf string
		if opts.ResizeHeight > 0 {
			vf = fmt.Sprintf("scale=-2:%d", opts.ResizeHeight)
		} else {
			vf = fmt.Sprintf("scale=%d:-2", opts.ResizeWidth)
		}
		argsChange = append(argsChange, "-vf", vf)
	}

	// Replace audio track if requested
	if opts.ReplaceAudio != "" {
		if _, err := os.Stat(opts.ReplaceAudio); err != nil {
			return fmt.Errorf("audio file not found: %s", opts.ReplaceAudio)
		}
		argsChange = append(argsChange, "-i", opts.ReplaceAudio)

		// if HAVE watermark 		--> outputs: 0=video, 1=watermark, 2=replaceAudio
		// if DO NOT HAVE watermark --> outputs: 0=video, 1=replaceAudio
		if hasWatermark {
			argsChange = append(argsChange, "-map", "2:a")
		} else {
			argsChange = append(argsChange, "-map", "1:a")
		}
	}

	// Choose codec and format: copy or re-encode with fast start
	if hasFilter || opts.ReplaceAudio != "" {
		argsChange = append(argsChange, "-c:v", "libx264")
		if opts.ReplaceAudio != "" {
			argsChange = append(argsChange, "-c:a", "aac")
		} else if !opts.RemoveAudio {
			argsChange = append(argsChange, "-c:a", "copy")
		}
	} else {
		argsChange = append(argsChange, "-c", "copy")
	}

	// Specify output file
	argsChange = append(argsChange, output)
	return progress.RunWithProgress(
		progress.OperationChange,
		input,
		argsChange,
	)
}
