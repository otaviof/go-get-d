[![test][workflowsTestBadge]][workflowsTest]

`go-get-d`
----------

## Abstract

This application brings back the `go get -d` functionality to *modern* Go. If you're looking for a helper tool to organize Go projects under `${GOPATH}`, this is for you!

The classic `go get -d` use to clone the import repository and build the application right away, placing the executable on the `${GOPATH}/bin` directory. Nowadays, you adopt the `go install` if the idea is only install the application executable.

Thus `go-get-d` is most useful for a developer to quickly get started on a Go repository, using the "classic" `${GOPATH}` organization style.

## Installation

Use `go install` as the following example:


```bash
go install github.com/otaviof/go-get-d@latest
```

When working on this repository you can alternatively use the `install` target, i.e.:

```bash
make install
```

The executable is placed on `${GOPATH}/bin`.

## Usage

The usage is straightforward, the only input required is the import name. For instance:

```bash
go-get-d github.com/otaviof/go-get-d
```

## Shell Eval

A practical way to clone the import repository and enter the directory, is using `go-get-d` output as a shell `eval` expession. The shell will pick executed the uncommented commands, i.e.:

```bash
eval "$(go-get-d github.com/otaviof/go-get-d)"
```

The output produced follow the snippet below:

```
$ go-get-d github.com/otaviof/go-get-d
# Go Module: "github.com/otaviof/go-get-d"
# Directory: "${GOPATH}/src/github.com/otaviof/go-get-d"
cd "${GOPATH}/src/github.com/otaviof/go-get-d"
```

After you run the `eval`, the current directory will change accordingly:

```
$ eval "$(go-get-d github.com/otaviof/go-get-d)"
$ pwd
/home/otaviof/go/src/github.com/otaviof/go-get-d
```

## Inspect Import

You can additionally *inspect* the given import searching for the `main` package, when found it will be subject to `go build`, creating the application executable. Please consider the usage of the `--inspect` flag:

```bash
go-get-d --inspect [import]
```

For example:

```
$ go-get-d --inspect github.com/otaviof/go-get-d
# Go Module: "github.com/otaviof/go-get-d"
# Directory: "${GOPATH}/src/github.com/otaviof/go-get-d"
# Cloning repository...
$ git clone https://github.com/otaviof/go-get-d ~/go/src/github.com/otaviof/go-get-d
## Cloning into '/Users/otaviof/go/src/github.com/otaviof/go-get-d'...
# Inspecting Go package...
# All done!
```
The `--inspect` will act as `go install`, therefore the application executable is stored on the `${GOPATH}/bin`.

[workflowsTest]: https://github.com/otaviof/go-get-d/actions/workflows/test.yaml
[workflowsTestBadge]: https://github.com/otaviof/go-get-d/actions/workflows/test.yaml/badge.svg
