package main

import (
	"os"

	"github.com/spf13/cobra"
)

const description = `# go-get-d

Emulates the behavior of the old "go get -d" command. It uses the provided import
module name to discover the repository URL, then downloads and builds its "main"
package.

In other words, this is the equivalent of an opinionated repository clone, given
it will always rely on "git" and a subsequent "go install".

It is still really useful for developers looking for a quick way to clone and
build a Go repository.
`

func main() {
	g := NewGoGetD()

	cmd := &cobra.Command{
		Use:          "go-get-d [flags] <import>",
		Short:        "Replaces the deprecated 'go get -d' command",
		Long:         description,
		SilenceUsage: true,

		// Ensure the CLI requires one argument, mandatory.
		Args: cobra.ExactArgs(1),

		// Validates the input and arguments.
		PreRunE: g.PreRunE,
		// Runs the primary business logic.
		RunE: g.RunE,
	}

	// Adding persistent flags to the command.
	g.PersistentFlags(cmd.PersistentFlags())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
