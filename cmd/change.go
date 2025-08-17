package cmd

import (
	"github.com/WeAreTheSameBlood/malva-cli/cmd/constants"
	"github.com/WeAreTheSameBlood/malva-cli/cmd/helpers/tasks"
	"github.com/spf13/cobra"
)

var changeCmd = &cobra.Command{
	Use:   "change <file>",
	Short: "Modify video without cutting",
	Long:  `Apply transformations like removing audio track to a video`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		input := args[0]
		var opts tasks.ChangeOptions
		opts.RemoveAudio, _ = cmd.Flags().GetBool("remove-audio")
		opts.Watermark, _ = cmd.Flags().GetString("watermark")
		opts.ReplaceAudio, _ = cmd.Flags().GetString("replace-audio")
		opts.Output, _ = cmd.Flags().GetString("output")
		opts.ResizeHeight, _ = cmd.Flags().GetInt("resize-height")
		opts.ResizeWidth, _ = cmd.Flags().GetInt("resize-width")

		return tasks.ProcessChange(input, opts)
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)
	changeCmd.Flags().Bool(
		"remove-audio", constants.CHANGE_REMOVE_AUDIO,
		"remove audio track",
	)
	changeCmd.Flags().String(
		"watermark", "",
		"path to PNG watermark image",
	)
	changeCmd.Flags().String(
		"replace-audio", "",
		"path to audio file to replace",
	)
	changeCmd.Flags().StringP(
		"output", "o",
		"", "output filename (default mod_<input>)",
	)
	changeCmd.Flags().Int(
		"resize-height", constants.CHANGE_DEFAULT_RESIZE_HEIGHT,
		"resize video to this height, preserving aspect ratio",
	)
	changeCmd.Flags().Int(
		"resize-width", constants.CHANGE_DEFAULT_RESIZE_WIDTH,
		"resize video to this width, preserving aspect ratio",
	)
}
