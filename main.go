package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "go-get-d [import]",
	Short:        "Replaces the deprecated 'go get -d'",
	PreRunE:      PreRunE,
	RunE:         RunE,
	SilenceUsage: true,
}

// PreRunE validates the arguments, one module must be informed.
func PreRunE(_ *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("go import name is required")
	}
	return nil
}

// RunE perform the primary business logic to simulate deprecated "go get -d" behavior.
func RunE(_ *cobra.Command, args []string) error {
	g := NewGoGetD(args[0])

	var err error
	if err = g.ParseURL(); err != nil {
		return err
	}
	if err = g.LookupModuleDirInGopath(); err != nil {
		return err
	}
	if !g.ModuleDirExits() {
		if err = g.CloneRepository(); err != nil {
			return err
		}
	}
	return g.InspectModulePackage()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
