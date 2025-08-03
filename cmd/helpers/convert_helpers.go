package helpers

import (
	"fmt"
	"github.com/WeAreTheSameBlood/malva-cli/cmd/constants"
	"os/exec"
	"path/filepath"
	"strings"
)

// ConvertOptions holds flags for the 'convert' command.
type ConvertOptions struct {
	GIF   bool
	FPS   int
	Scale int
}

// ProcessConvert runs ffmpeg to convert according to opts.
func ProcessConvert(
	input string,
	opts ConvertOptions,
) error {
	ext := filepath.Ext(input)
	base := strings.TrimSuffix(input, ext)

	var output string
	var args []string

	if opts.GIF {
		// fallback defaults if user passed zero or invalid
		if opts.FPS <= 0 {
			opts.FPS = constants.CONVERT_DEFAULT_FPS
		}
		if opts.Scale <= 0 {
			opts.Scale = constants.CONVERT_DEFAULT_SCALE
		}

		output = fmt.Sprintf("%s.gif", base)

		filter := fmt.Sprintf(
			"fps=%d,scale=%d:-1:flags=lanczos",
			opts.FPS,
			opts.Scale)

		args = []string{
			"-hide_banner", "-loglevel", "error",
			"-noautorotate",
			"-i", input,
			"-vf", filter,
			"-loop", "0",
			output,
		}
	} else {
		return fmt.Errorf("no conversion format specified")
	}

	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
