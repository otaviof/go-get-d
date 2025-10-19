package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/tools/go/packages"
)

// GoGetD represents the deprecated "go get -d" command, which use to download a
// given Golang module and let it ready to work, basically a simple "git clone"
// into the GOPATH.
type GoGetD struct {
	inspect bool   // builds the target executable
	path    bool   // prints out the module path in the system
	verbose bool   // prints out verbose information
	input   string // raw package name

	fullPath       string   // full path to the module
	relativeGoPath string   // relative path to GOPATH
	module         string   // valid golang import name
	repositoryURL  *url.URL // module repository URL
}

// PersistentFlags decorates the flag set with application flags.
func (g *GoGetD) PersistentFlags(p *pflag.FlagSet) {
	p.BoolVarP(&g.inspect, "inspect", "i", false,
		"Inspect package, build the main executable")
	p.BoolVarP(&g.path, "path", "p", false,
		"Prints only the path to the module directory")
	p.BoolVarP(&g.verbose, "verbose", "v", false,
		"Add verbose information")
}

// logger prints the information when flag path is disabled.
func (g *GoGetD) logger(format string, a ...any) {
	if g.path {
		return
	}
	fmt.Printf(format, a...)
}

// parseURL parses the input given to GoGetD in order to assert it's a valid URL,
// and to extract the golang module name given it can be employed as a valid URL.
func (g *GoGetD) parseURL() error {
	g.logger("# Parsing repository for %q...\n", g.input)
	u, err := url.Parse(g.input)
	if err != nil {
		return err
	}
	// When the first attempt does not extract scheme and hostname, we assume it's
	// a regular "https" URL and try to parse again with a more strict URL parser
	// shaking off left over input inconsistencies
	if u.Scheme == "" && u.Host == "" {
		u, err = url.ParseRequestURI(fmt.Sprintf("https://%s", g.input))
		if err != nil {
			return err
		}
	}

	g.module = urlToGoModule(u)
	g.logger("# Package module: %q\n", g.module)
	g.repositoryURL = u
	g.logger("# Repository URL: %q\n", g.repositoryURL)
	return nil
}

// lookupModuleDirInGopath based on GOPATH defines the go module directory path.
func (g *GoGetD) lookupModuleDirInGopath() error {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return fmt.Errorf("GOPATH environment variable is not set")
	}

	g.fullPath = path.Join(gopath, path.Join("src", g.module))
	g.relativeGoPath = strings.ReplaceAll(g.fullPath, gopath, "${GOPATH}")

	g.logger("# Package directory full path: %q\n", g.relativeGoPath)
	return nil
}

// moduleDirExits checks if the module directory path exists.
func (g *GoGetD) moduleDirExits() bool {
	info, err := os.Stat(g.fullPath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return info != nil && info.IsDir()
}

// cloneRepository executes "git clone" into the GOPATH based module directory.
func (g *GoGetD) cloneRepository(ctx context.Context) error {
	err := os.MkdirAll(g.fullPath, 0o755)
	if err != nil {
		return err
	}

	/* #nosec G204 */
	git := exec.CommandContext(
		ctx,
		"git",
		"clone",
		"--depth",
		"1",
		g.repositoryURL.String(),
		g.fullPath,
	)
	g.logger("# $ %s %s\n", git.Path, strings.Join(git.Args, " "))
	out, err := git.CombinedOutput()
	if err != nil {
		return err
	}

	bytesReader := bytes.NewReader(out)
	bufferReader := bufio.NewReader(bytesReader)
	for {
		line, _, err := bufferReader.ReadLine()
		if err == io.EOF {
			break
		}
		g.logger("## %s\n", line)
	}
	return nil
}

// inspectModulePackage tries to load the module, inspecting if it's correctly
// loading. Which means building the "main" package.
func (g *GoGetD) inspectModulePackage() error {
	g.logger("# Inspecting Go %q package...\n", g.fullPath)
	pkgs, err := packages.Load(
		&packages.Config{Mode: packages.NeedName, Dir: g.fullPath},
		g.module,
	)
	if err != nil {
		return err
	}
	if packages.PrintErrors(pkgs) > 0 {
		return fmt.Errorf("unable to load package %q", g.input)
	}
	lenPkgs := len(pkgs)
	if lenPkgs != 1 {
		return fmt.Errorf("found %d packages for module named %q",
			lenPkgs, g.module)
	}
	return nil
}

// PreRunE parse the flags and ensures the input module name is informed.
func (g *GoGetD) PreRunE(_ *cobra.Command, args []string) error {
	if g.path && g.verbose {
		return fmt.Errorf("--path and --verbose flag are incompatible")
	}
	if len(args) != 1 {
		return fmt.Errorf("expected exactly one argument, got %d", len(args))
	}
	g.input = args[0]
	return nil
}

// RunE runs the main application logic.
func (g *GoGetD) RunE(cmd *cobra.Command, _ []string) error {
	if g.input == "" {
		return fmt.Errorf("no input module name is informed")
	}

	err := g.parseURL()
	if err != nil {
		return err
	}

	if err = g.lookupModuleDirInGopath(); err != nil {
		return err
	}

	if g.path {
		fmt.Println(g.fullPath)
	} else {
		fmt.Printf("cd %q\n", g.relativeGoPath)
	}

	if !g.moduleDirExits() {
		if err = g.cloneRepository(cmd.Context()); err != nil {
			return err
		}
	}

	if g.inspect {
		if err = g.inspectModulePackage(); err != nil {
			return err
		}
	}

	g.logger("# All done!\n")
	return nil
}

// NewGoGetD instantiate GoGetD passing the raw input.
func NewGoGetD() *GoGetD {
	return &GoGetD{}
}
