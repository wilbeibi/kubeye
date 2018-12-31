package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(podsCmd)
}

var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Watching Pods",
	Run: func(cmd *cobra.Command, args []string) {
		controller.Run()
	},
}
