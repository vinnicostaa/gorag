package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newNeo4jCommand(root *rootOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "neo4j",
		Short: "Manage Neo4j schema and imports",
		Long: `Neo4j commands initialize graph constraints and import generated
gorag artifacts into a Neo4j database.

The Neo4j implementation will be added after the JSONL artifacts are stable.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(
		newNeo4jInitCommand(root),
		newNeo4jImportCommand(root),
	)

	return cmd
}

func newNeo4jInitCommand(root *rootOptions) *cobra.Command {
	return &cobra.Command{
		Use:   "Init",
		Short: "Initialize Neo4j constrainsts and indexes",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("not implemented yet: neo4j init")
		},
	}
}

func newNeo4jImportCommand(root *rootOptions) *cobra.Command {
	return &cobra.Command{
		Use:   "import",
		Short: "Import generated artifacts into Neo4j",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("not implemented yet: neo4j import")
		},
	}
}
