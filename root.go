package main

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the root command, the main primary entrypoint.
type RootCmd struct {
	cmd     *cobra.Command // cobra instance
	inspect bool           // inspect flag
}

// Execute runs the command.
func (r *RootCmd) Execute() error {
	return r.cmd.Execute()
}

// RunE perform the core application logic.
func (r *RootCmd) RunE(_ *cobra.Command, args []string) error {
	g := NewGoGetD(args[0])

	err := g.ParseURL()
	if err != nil {
		return err
	}

	if err = g.LookupModuleDirInGopath(); err != nil {
		return err
	}
	g.PrintChangeDir()

	if !g.ModuleDirExits() {
		if err = g.CloneRepository(); err != nil {
			return err
		}
	}

	if !r.inspect {
		return nil
	}
	return g.InspectModulePackage()
}

// NewRootCmd creates a new root command, with flags and core logic.
func NewRootCmd() *RootCmd {
	r := &RootCmd{
		cmd: &cobra.Command{
			Use:          "go-get-d [import]",
			Short:        "Replaces the deprecated 'go get -d'",
			SilenceUsage: true,
			// Ensure the CLI requires one argument, mandatory.
			Args: cobra.ExactArgs(1),
		},
		inspect: false,
	}
	r.cmd.RunE = r.RunE

	f := r.cmd.PersistentFlags()
	f.BoolVarP(&r.inspect, "inspect", "i", false,
		"Inspect package, build the main executable")

	return r
}
