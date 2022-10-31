`go-get-d`
----------

[![CI][projectBadgeSVG]](https://github.com/otaviof/go-get-d/actions/workflows/test.yaml)

## Abstract

This application brings back the `go get -d` functionality to modern Go. If you're looking for a helper tool to organize Go projects under `${GOPATH}`, this is for you!

## Installing

The best way to install is through `go install`, as the example below:


```bash
go install github.com/otaviof/go-get-d@latest
```

## Usage

The usage is straight forward, the only input required is the import name.

## Shell Eval

A practical way to `git clone` and enter the repository directory is using `go-get-d` output as a shell `eval` expession. The shell will pick up the uncommented `cd` command and run immediately, for instance:

```bash
eval "$(go-get-d github.com/otaviof/go-get-d)"
```

Producing the following outcome:

```
$ eval "$(go-get-d github.com/otaviof/go-get-d)"
$ pwd
/home/otaviof/go/src/github.com/otaviof/go-get-d
```

## Inspect Import

You can additionally inspect the import searching for `main` package to `go build` the project executable. To inspect the import use `go-get-d --inspect`, i.e.:


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

[projectBadgeSVG]: https://github.com/otaviof/go-get-d/actions/workflows/test.yaml/badge.svg
[projectTestWorkflow]: https://github.com/otaviof/go-get-d/actions/workflows/test.yaml
