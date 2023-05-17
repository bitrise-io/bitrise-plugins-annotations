package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bitrise-io/bitrise-plugins-annotations/service"
)

var deleteCmd = &cobra.Command{
	Use:   "delete context",
	Short: "Delete build annotation by context",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.Delete(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
