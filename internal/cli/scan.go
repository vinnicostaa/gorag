package cli

import (
	"fmt"

	"github.com/vinnicostaa/gorag/internal/app"

	"github.com/spf13/cobra"
)

func newScanCommand(root *rootOptions) *cobra.Command {
	var vaultPath string
	var includeHidden bool

	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Scan a vault and write traceable artifacts",
		Long: `Scan walks through a local Markdown or Obsidian vault and writes intermediate artifacts to the output directory.

		At this stage, scan only crates the initial manifest. The full Markdown parser, chunker and JSONL writers will be added incrementally.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if vaultPath == "" {
				return fmt.Errorf("vault path is required")
			}

			application := app.New(app.Config{
				Verbose: root.verbose,
			})

			output, err := application.Scan(cmd.Context(), app.ScanInput{
				VaultPath:     vaultPath,
				OutputDir:     root.outputDir,
				IncludeHidden: includeHidden,
			})
			if err != nil {
				return err
			}

			_, err = fmt.Fprintf(
				cmd.OutOrStdout(),
				"scan initialized\nmanifest: %s\n",
				output.ManifestPath,
			)
			return err
		},
	}

	cmd.Flags().StringVar(
		&vaultPath,
		"vault",
		"",
		"path to the Markdown or Obsidian vault",
	)

	cmd.Flags().BoolVar(
		&includeHidden,
		"incinclude-hidden",
		false,
		"include hidden files and directories",
	)

	if err := cmd.MarkFlagRequired("vault"); err != nil {
		panic(err)
	}

	return cmd
}
