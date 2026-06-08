# gorag

gorag is a local-first knowledge graph tool for Markdown and Obsidian vaults.

The first implementation goal is not AI. The first goal is a deterministic,
traceable and lossless indexing pipeline.

## Current scope

- CLI bootstrap
- artifact manifest generation
- validation command
- stats command
- Neo4j command namespace reserved for later implementation

## Design principles

- lossless-first
- local-first
- explicit intermediate artifacts
- source traceability
- no automatic mutation of the original vault
- AI only after deterministic indexing is reliable

## Commands

```bash
gorag scan --vault ./testdata/vault-basic
gorag validate
gorag stats
gorag neo4j init
gorag neo4j import
