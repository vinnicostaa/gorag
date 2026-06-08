// Package cli wires gorag use cases to command-line commands.
package cli

import (
	"context"

	"github.com/spf13/cobra"
)

type rootOptions struct {
	outputDir string
	verbose   bool
}

// ExecuteContext builds and executes the root command using the provider context.
func ExecuteContext(ctx context.Context) error {
	cmd := NewRootCommand()
	return cmd.ExecuteContext(ctx)
}

// NewRootCommand creates the root gorag command.
func NewRootCommand() *cobra.Command {
	opts := &rootOptions{}

	cmd := &cobra.Command{
		Use:   "gorag",
		Short: "Build a traceable graph from local knowledge vaults",
		Long: `gorag indexes local Markdown and Obsidian vaults into traceable intermediate artifacts and, later, graph databases such as Neo4j.

		The first design rule is lossless-first: generated artifacts must preserve enough source information to trace every chunk back to its original file.`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.PersistentFlags().StringVar(
		&opts.outputDir,
		"out",
		".vaultgraph/out",
		"directory used for generated artifacts",
	)

	cmd.PersistentFlags().BoolVarP(
		&opts.verbose,
		"verbose",
		"v",
		false,
		"enable verbose output",
	)

	cmd.AddCommand(
		newScanCommand(opts),
		newValidateCommand(opts),
		newStatsCommand(opts),
		newNeo4jCommand(opts),
	)

	return cmd
}
