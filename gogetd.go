package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"

	"golang.org/x/tools/go/packages"
)

// GoGetD represents the deprecated "go get -d" command, which use to download a given Golang module
// and let it ready to work, basically a simple "git clone" into the GOPATH.
type GoGetD struct {
	input         string   // raw package name
	module        string   // valid golang import, extracted from "input"
	repositoryURL *url.URL // module repository URL
	dir           string   // path to the module inside GOPATH
	fullDir       string   // printable directory full path
}

// ParseURL parses the input given to GoGetD in order to assert it's a valid URL, and to extract the
// golang module name given it can be employed as a valid URL.
func (g *GoGetD) ParseURL() error {
	u, err := url.Parse(g.input)
	if err != nil {
		return err
	}
	// when the first attempt does not extract scheme and hostname, we assume it's a regular "https"
	// URL and try to parse again with a more strict URL parser shaking off left over input
	// inconsistencies
	if u.Scheme == "" && u.Host == "" {
		if u, err = url.ParseRequestURI(fmt.Sprintf("https://%s", g.input)); err != nil {
			return err
		}
	}

	g.module = urlToGoModule(u)
	g.repositoryURL = u
	fmt.Printf("# Go Module: %q\n", g.module)
	return nil
}

// LookupModuleDirInGopath based on GOPATH defines the go module directory path.
func (g *GoGetD) LookupModuleDirInGopath() error {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return fmt.Errorf("GOPATH environment variable is not set")
	}

	g.dir = path.Join(gopath, path.Join("src", g.module))
	g.fullDir = strings.ReplaceAll(g.dir, gopath, "${GOPATH}")
	fmt.Printf("# Directory: %q\n", g.fullDir)
	return nil
}

// PrintChangeDir prints the full directory path with "cd" command in front.
func (g *GoGetD) PrintChangeDir() {
	fmt.Printf("cd %q\n", g.fullDir)
}

// ModuleDirExits checks if the module directory path exists.
func (g *GoGetD) ModuleDirExits() bool {
	info, err := os.Stat(g.dir)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return info != nil && info.IsDir()
}

// CloneRepository executes "git clone" into the GOPPATH based module directory.
func (g *GoGetD) CloneRepository() error {
	err := os.MkdirAll(g.dir, 0755)
	if err != nil {
		return err
	}

	gitCloneArgs := []string{"clone", g.repositoryURL.String(), g.dir}
	cmd := exec.Command("git", gitCloneArgs...)

	fmt.Println("# Cloning repository...")
	fmt.Printf("# $ %s\n", cmd.String())
	out, err := cmd.CombinedOutput()
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
		fmt.Printf("## %s\n", line)
	}
	return nil
}

// InspectModulePackage tries to load the module, inspecting if it's correctly loading.
func (g *GoGetD) InspectModulePackage() error {
	fmt.Println("# Inspecting Go package...")
	pkgs, err := packages.Load(&packages.Config{Mode: packages.NeedName, Dir: g.dir}, g.module)
	if err != nil {
		return err
	}
	if packages.PrintErrors(pkgs) > 0 {
		return fmt.Errorf("unable to load package %q", g.input)
	}
	lenPkgs := len(pkgs)
	if lenPkgs != 1 {
		return fmt.Errorf("found %d packages for module named %q", lenPkgs, g.module)
	}
	fmt.Println("# All done!")
	return nil
}

// NewGoGetD instantiate GoGetD passing the raw input.
func NewGoGetD(input string) *GoGetD {
	return &GoGetD{input: input}
}
