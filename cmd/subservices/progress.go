package subservices

import (
	"bufio"
	"fmt"
	"github.com/WeAreTheSameBlood/malva-cli/cmd/helpers"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var timeRe = regexp.MustCompile(`time=(\d+):(\d+):(\d+\.?\d*)`)

// OperationType describes the name to show in progress bar + additional info
type OperationType string

const (
	OperationCut     OperationType = "Cutting"
	OperationChange  OperationType = "Changing"
	OperationConvert OperationType = "Converting"
)

// RunWithProgress executes `ffmpeg` with given (completed) args and show progress bar
func RunWithProgress(
	operation OperationType,
	input string,
	ffmpegArgs []string,
) error {
	// Probe duration
	var probeCmd = exec.Command("ffprobe",
		"-v", "error",
		"-show_entries", "format=duration",
		"-of", "default=nokey=1:noprint_wrappers=1",
		input,
	)
	var out, err = probeCmd.Output()

	if err != nil {
		return fmt.Errorf("ffprobe error: %w", err)
	}

	var totalSeconds, _ = strconv.ParseFloat(strings.TrimSpace(string(out)), 64)

	// Start ffmpeg
	var resultCommand = exec.Command("ffmpeg", ffmpegArgs...)
	var progressPipe, _ = resultCommand.StdoutPipe()

	if runError := resultCommand.Start(); runError != nil {
		return runError
	}

	// Progress loop
	var startTime = time.Now()
	var progressScanner = bufio.NewScanner(progressPipe)
	var lastPercentage = -1

	for progressScanner.Scan() {
		var scannedLine = progressScanner.Text()
		var matchParts = timeRe.FindStringSubmatch(scannedLine)
		if matchParts != nil {
			// parse current ffmpeg time into seconds
			var currentSeconds = helpers.ParseFFmpegTime(
				matchParts[1], // hours
				matchParts[2], // minutes
				matchParts[3], // seconds.milliseconds
			)
			// compute raw percentage round to int
			var ratio = currentSeconds / totalSeconds
			var rawPercentage = int(ratio*100 + 0.5)

			// round to nearest --> 5%
			var roundedPercentage = (rawPercentage + 2) / 5 * 5
			if roundedPercentage > 100 {
				roundedPercentage = 100
			}

			if roundedPercentage != lastPercentage {
				lastPercentage = roundedPercentage
				var filledBars = roundedPercentage * 20 / 100
				var emptyBars = 20 - filledBars

				fmt.Printf(
					"\r[%s%s] %3d%% %s %s",
					strings.Repeat("|", filledBars),
					strings.Repeat("-", emptyBars),
					roundedPercentage,
					time.Since(startTime).Truncate(time.Second),
					operation,
				)
			}
		}
	}

	// Finalise processing results
	fmt.Printf(
		"\r[%s] 100%% %s %s\n",
		strings.Repeat("|", 20),
		time.Since(startTime).Truncate(time.Second),
		operation,
	)

	// Print the output data
	var outputFilename = ffmpegArgs[len(ffmpegArgs)-1]
	var absolutePath, absErr = filepath.Abs(outputFilename)

	if absErr != nil {
		absolutePath = outputFilename
	}

	fmt.Printf(
		"Saved as %s\nPath --> %s\n",
		outputFilename,
		absolutePath,
	)

	return resultCommand.Wait()
}
