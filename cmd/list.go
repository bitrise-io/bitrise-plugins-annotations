package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/bitrise-io/bitrise-plugins-annotations/service"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List build annotations",
	RunE: func(cmd *cobra.Command, args []string) error {
		list, err := service.List()
		if err != nil {
			return err
		}

		fmt.Println(list)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
