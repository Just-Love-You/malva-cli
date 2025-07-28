package helpers

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// ConvertOptions holds flags for the 'convert' command.
type ConvertOptions struct {
	GIF bool
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
		output = fmt.Sprintf("%s.gif", base)
		args = []string{
			"-hide_banner", "-loglevel", "error",
			"-noautorotate",
			"-i", input,
			"-vf", "fps=15,scale=380:-1:flags=lanczos",
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
