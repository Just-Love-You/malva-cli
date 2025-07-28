package cmd

import (
	"github.com/spf13/cobra"

	"github.com/WeAreTheSameBlood/malva-cli/cmd/helpers"
)

// convertCmd represents the 'convert' command
var convertCmd = &cobra.Command{
	Use:   "convert <file>",
	Short: "Convert video to another format",
	Long:  `Convert a video file into a different format (e.g. GIF)`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		input := args[0]

		var opts helpers.ConvertOptions
		opts.GIF, _ = cmd.Flags().GetBool("gif")

		return helpers.ProcessConvert(input, opts)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().Bool("gif", false, "convert input video to animated GIF")
}
