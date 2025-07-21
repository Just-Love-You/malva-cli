/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
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
		start, _ := cmd.Flags().GetString("start")
		finish, _ := cmd.Flags().GetString("finish")
		offAudio, _ := cmd.Flags().GetBool("off-audio")
		output := fmt.Sprintf("cut_%s", input)
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
