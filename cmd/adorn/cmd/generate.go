package cmd

import (
	"os"

	"github.com/mmcloughlin/adorn"
	"github.com/spf13/cobra"
)

var cfg adorn.Config

func init() {
	generateCmd.Flags().StringVarP(&cfg.Package, "package", "p", "", "package name")
	generateCmd.Flags().StringVarP(&cfg.TypeName, "type", "t", "", "type name")
	generateCmd.Flags().StringVarP(&cfg.MethodName, "method", "m", "", "method name")
	generateCmd.Flags().StringSliceVarP(&cfg.ArgumentTypes, "args", "a", nil, "argument types")
	generateCmd.Flags().StringSliceVarP(&cfg.ReturnTypes, "return", "r", nil, "return types")

	RootCmd.AddCommand(generateCmd)
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate type and adornments",
	RunE: func(cmd *cobra.Command, args []string) error {
		return adorn.Generate(cfg, os.Stdout)
	},
}
