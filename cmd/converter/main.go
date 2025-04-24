package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "converter"}
	var inputFile, configFile string

	var jsonc2mdCmd = &cobra.Command{
		Use:   "jsonc2md",
		Short: "Convert USECASES.jsonc to Markdown",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := LoadConfig(configFile)
			if err != nil {
				return fmt.Errorf("load config: %w", err)
			}
			useCases, err := ParseJSONC(inputFile)
			if err != nil {
				return fmt.Errorf("parse JSONC: %w", err)
			}
			return GenerateMarkdown(useCases, cfg.OutputMD)
		},
	}
	jsonc2mdCmd.Flags().StringVarP(&inputFile, "input", "i", "USECASES.jsonc", "Input JSONC file")
	jsonc2mdCmd.Flags().StringVarP(&configFile, "config", "c", "converter.yaml", "Config file")

	var jsonc2mmdCmd = &cobra.Command{
		Use:   "jsonc2mmd",
		Short: "Convert USECASES.jsonc to Mermaid",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := LoadConfig(configFile)
			if err != nil {
				return fmt.Errorf("load config: %w", err)
			}
			useCases, err := ParseJSONC(inputFile)
			if err != nil {
				return fmt.Errorf("parse JSONC: %w", err)
			}
			return GenerateMermaid(useCases, cfg.OutputMMD)
		},
	}
	jsonc2mmdCmd.Flags().StringVarP(&inputFile, "input", "i", "USECASES.jsonc", "Input JSONC file")
	jsonc2mmdCmd.Flags().StringVarP(&configFile, "config", "c", "converter.yaml", "Config file")

	rootCmd.AddCommand(jsonc2mdCmd, jsonc2mmdCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
