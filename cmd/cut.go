package cmd

import (
	"github.com/WeAreTheSameBlood/malva-cli/cmd/helpers"
	"github.com/WeAreTheSameBlood/malva-cli/cmd/helpers/tasks"
	"time"

	"github.com/spf13/cobra"
)

// cutCmd represents the 'cut' command
var cutCmd = &cobra.Command{
	Use:   "cut <file>",
	Short: "Cut a segment from a video file",
	Long:  `Cut a segment from a video file using start and finish times`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		input := args[0]

		// collect flags into opts
		var opts tasks.CutOptions

		// parse start (supports durations like "14s", "1m9s", "1h2m3s500ms")
		startRaw, _ := cmd.Flags().GetString("start")
		if d, err := time.ParseDuration(startRaw); err == nil {
			opts.Start = helpers.FormatDurationFFMPEG(d)
		} else {
			opts.Start = startRaw
		}

		// parse finish
		finishRaw, _ := cmd.Flags().GetString("finish")
		if d, err := time.ParseDuration(finishRaw); err == nil {
			opts.Finish = helpers.FormatDurationFFMPEG(d)
		} else {
			opts.Finish = finishRaw
		}

		// other flags
		opts.Output, _ = cmd.Flags().GetString("output")

		// execute
		return tasks.ProcessCut(input, opts)
	},
}

func init() {
	rootCmd.AddCommand(cutCmd)

	cutCmd.Flags().StringP(
		"start", "s", "",
		"start time (HH:MM:SS[.ms], e.g. 14s, 1m9s)",
	)
	cutCmd.Flags().StringP(
		"finish", "f", "",
		"finish time (HH:MM:SS[.ms])",
	)
	cutCmd.Flags().StringP(
		"output", "o", "",
		"output filename (default cut_<input>)",
	)
}
