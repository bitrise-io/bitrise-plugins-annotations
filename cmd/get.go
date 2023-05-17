package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/bitrise-io/bitrise-plugins-annotations/service"
)

var getCmd = &cobra.Command{
	Use:   "get context",
	Short: "Get build annotation by context",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		annotation, err := service.Get(args[0])
		if err != nil {
			return err
		}

		fmt.Println(annotation)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
