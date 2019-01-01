package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wilbeibi/kubeye/pkg/controller"
)

func init() {
	rootCmd.AddCommand(podsCmd)
}

var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Watching Pods",
	Run: func(cmd *cobra.Command, args []string) {
		stopChan := make(chan struct{})
		controller := controller.NewController()
		controller.Run(stopChan)
	},
}
