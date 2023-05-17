package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bitrise-io/bitrise-plugins-annotations/service"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete all build annotations",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.DeleteAll()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
