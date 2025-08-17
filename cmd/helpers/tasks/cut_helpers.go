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
	var argsCut = append([]string{}, constants.COMMON_FFMPEG_ARGUMENTS...)
	argsCut = append(
		argsCut,
		"-i", input,
		"-ss", opts.Start,
		"-to", opts.Finish,
		"-c", "copy",
		output,
	)

	return progress.RunWithProgress(
		progress.OperationCut,
		input,
		argsCut,
	)
}
