package cmd

import (
	"github.com/WeAreTheSameBlood/malva-cli/cmd/constants"
	"github.com/WeAreTheSameBlood/malva-cli/cmd/helpers/tasks"
	"github.com/spf13/cobra"
)

// convertCmd represents the 'convert' command
var convertCmd = &cobra.Command{
	Use:   "convert <file>",
	Short: "Convert video to another format",
	Long:  `Convert a video file into a different format (e.g. GIF).`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		input := args[0]

		var opts tasks.ConvertOptions
		opts.GIF, _ = cmd.Flags().GetBool("gif")
		opts.FPS, _ = cmd.Flags().GetInt("fps")
		opts.Scale, _ = cmd.Flags().GetInt("scale")

		return tasks.ProcessConvert(input, opts)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().Bool(
		"gif", constants.CONVERT_DEFAULT_GIF,
		"convert input video to animated GIF",
	)
	convertCmd.Flags().Int(
		"fps", constants.CONVERT_DEFAULT_FPS,
		"frames per second for GIF output",
	)
	convertCmd.Flags().Int(
		"scale", constants.CONVERT_DEFAULT_SCALE,
		"width to scale GIF to (preserves aspect ratio)",
	)
}
