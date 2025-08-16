package subservices

import "regexp"

var timeRe = regexp.MustCompile(`time=(\d+):(\d+):(\d+\.?\d*)`)

// OperationType describes the name to show in progress bar + additional info
type OperationType string

const (
	OpCut     OperationType = "Cutting"
	OpChange  OperationType = "Changing"
	OpConvert OperationType = "Converting"
)

// RunWithProgress executes `ffmpeg` with given (completed) args and show progress bar
func RunWithProgress(
	operation OperationType,
	input string,
	ffmpegArgs []string,
) error {
	return nil
}
