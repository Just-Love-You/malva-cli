package tasks

import (
	"fmt"
	"github.com/WeAreTheSameBlood/malva-cli/cmd/constants"
	progress "github.com/WeAreTheSameBlood/malva-cli/cmd/subservices"
)

// CutOptions holds parameters for the 'cut' command.
type CutOptions struct {
	Start  string
	Finish string
	Output string
}

// ProcessCut builds ffmpeg arguments and executes the cut command.
func ProcessCut(
	input string,
	opts CutOptions,
) error {
	// determine output filename
	output := opts.Output
	if output == "" {
		output = fmt.Sprintf(constants.CUT_DEFAULT_OUTPUT_NAME_PREFIX, input)
	}

	// build ffmpeg argsCut
	argsCut := []string{
		"-hide_banner",
		"-loglevel", "info",
		"-progress", "pipe:1",
		"-i", input,
		"-ss", opts.Start,
		"-to", opts.Finish,
	}

	argsCut = append(argsCut, "-c", "copy", output)

	return progress.RunWithProgress(
		progress.OperationCut,
		input,
		argsCut,
	)
}
