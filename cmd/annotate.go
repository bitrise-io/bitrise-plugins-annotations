package cmd

import (
	"errors"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/bitrise-io/bitrise-plugins-annotations/service"
)

var (
	context string
	style   string
)

var annotateCmd = &cobra.Command{
	Use:   "annotate [markdown]",
	Short: "Annotate build",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var markdown string

		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			stdin, err := io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}

			markdown = string(stdin)
		}

		if len(args) == 1 {
			if markdown != "" {
				return errors.New("if stdin piping is used then markdown argument can't be set")
			}

			markdown = args[0]
		}

		if markdown == "" {
			return errors.New("markdown is empty")
		}

		return service.Annotate(service.Annotation{
			Context:  context,
			Markdown: markdown,
			Style:    style,
		})
	},
}

func init() {
	rootCmd.AddCommand(annotateCmd)

	annotateCmd.Flags().StringVarP(&context, "context", "c", "",
		"the context to find existing annotations and replace their content")
	annotateCmd.Flags().StringVarP(&style, "style", "s", "default",
		"the style to use for this annotation, such as default, error, warning, info")
}
