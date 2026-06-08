package cli

import (
	"encoding/json"
	"fmt"

	"github.com/vinnicostaa/gorag/internal/app"

	"github.com/spf13/cobra"
)

func newStatsCommand(root *rootOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stats",
		Short: "Print generated artifact statistics",
		Long: `Stats reads generated artifacts and prints a summary.

The first implementation prints the manifest. Later it will include counts for
notes, chunks, links, tags, aliases, diagnostics and unresolved references.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			application := app.New(app.Config{
				Verbose: root.verbose,
			})

			output, err := application.Stats(cmd.Context(), app.StatsInput{
				OutputDir: root.outputDir,
			})
			if err != nil {
				return err
			}

			encoded, err := json.MarshalIndent(output.Manifest, "", "	")
			if err != nil {
				return fmt.Errorf("encoded stats output: %w", err)
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), string(encoded))
			return err
		},
	}
	return cmd
}
