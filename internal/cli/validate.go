package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vinnicostaa/gorag/internal/app"
)

func newValidateCommand(root *rootOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate generated gorag artifacts",
		Long: `Validate checks whether generated artifacts exist and are internally
consistent.

The first implementation validates only the manifest. Later it will verify
coverage, chunks, links, tags and Neo4j import readiness.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			application := app.New(app.Config{
				Verbose: root.verbose,
			})

			output, err := application.Validate(cmd.Context(), app.ValidateInput{
				OutputDir: root.outputDir,
			})
			if err != nil {
				return err
			}

			_, err = fmt.Fprintf(
				cmd.OutOrStdout(),
				"validation ok\nmanifest: %s\n",
				output.ManifestPath,
			)
			return err
		},
	}

	return cmd
}
