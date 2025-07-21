package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"time"
)

// cutCmd represents the cut command
var cutCmd = &cobra.Command{
	Use:   "cut",
	Short: "Cut a segment from a video file",
	Long:  `Cut a segment from a video file using start and finish times.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("no input file specified")
		}
		input := args[0]
		var startRaw string
		var finishRaw string
		startRaw, _ = cmd.Flags().GetString("start")
		finishRaw, _ = cmd.Flags().GetString("finish")

		var start string = startRaw
		if d, err := time.ParseDuration(startRaw); err == nil {
			start = FormatDurationFFMPEG(d)
		}

		var finish string = finishRaw
		if d, err := time.ParseDuration(finishRaw); err == nil {
			finish = FormatDurationFFMPEG(d)
		}

		var offAudio bool
		offAudio, _ = cmd.Flags().GetBool("off-audio")

		var output string = fmt.Sprintf("cut_%s", input)
		argsFfmpeg := []string{
			"-hide_banner",
			"-loglevel",
			"error",
			"-i",
			input,
			"-ss",
			start,
			"-to",
			finish,
		}

		if offAudio {
			argsFfmpeg = append(argsFfmpeg, "-an")
		}
		argsFfmpeg = append(argsFfmpeg, "-c", "copy", output)

		cmdFfmpeg := exec.Command("ffmpeg", argsFfmpeg...)
		// cmdFfmpeg.Stdout = os.Stdout
		// cmdFfmpeg.Stderr = os.Stderr
		return cmdFfmpeg.Run()
	},
}

func init() {
	rootCmd.AddCommand(cutCmd)
	cutCmd.Flags().StringP(
		"start", "s", "",
		"start time (HH:MM:SS[.ms])",
	)
	cutCmd.Flags().StringP(
		"finish", "f", "",
		"finish time (HH:MM:SS[.ms])",
	)
	cutCmd.Flags().Bool(
		"off-audio", false,
		"remove audio track",
	)
}
