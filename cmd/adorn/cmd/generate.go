package cmd

import (
	"os"

	"github.com/mmcloughlin/adorn"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate type and adornments",
	RunE: func(cmd *cobra.Command, args []string) error {
		return generate()
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)
}

func generate() error {
	cfg := adorn.Config{
		Package:       "pkg",
		TypeName:      "Handler",
		MethodName:    "ServeHTTP",
		ArgumentTypes: []string{"http.ResponseWriter", "*http.Request"},
	}
	return adorn.Generate(cfg, os.Stdout)
}
