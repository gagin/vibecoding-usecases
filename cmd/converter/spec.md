# Converter Tool Specification

## Overview

The `converter` tool processes `USECASES.jsonc` to generate human-readable documentation and diagrams for vibecoding projects. It supports two commands:

- `jsonc2md`: Converts `USECASES.jsonc` to Markdown (`usecases.md`) with embedded Mermaid diagrams.
- `jsonc2mmd`: Extracts `mermaid_flow` fields into standalone Mermaid files (`diagrams.mmd`).

## Features

- **Input**: Parses `USECASES.jsonc`, handling JSONC comments.
- **Outputs**:
  - Markdown with formatted use case details and Mermaid diagrams.
  - Mermaid files for standalone diagram rendering.
- **Configuration**: Uses `converter.yaml` for output paths and options.
- **CLI**: Built with Cobra, supporting flags (e.g., `--input`, `--config`).

## Usage

```bash
# Generate Markdown
go run cmd/converter/main.go jsonc2md --input USECASES.jsonc

# Generate Mermaid
go run cmd/converter/main.go jsonc2mmd --input USECASES.jsonc

Configuration
converter.yaml:
yaml

output_md: usecases.md
output_mmd: diagrams.mmd

Implementation
input_parser.go: Parses JSONC, stripping comments, and defines use case struct.

output_md.go: Generates Markdown with Mermaid diagrams.

output_mmd.go: Outputs Mermaid files.

config.go: Loads converter.yaml with Viper.

helpers.go: Utility functions (e.g., file writing).

main.go: CLI entry point with Cobra.
