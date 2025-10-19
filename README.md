[![go-report-card][goReportCardBadge]][goReportCard]

`go-get-d`
----------

## Abstract

This application restores the `go get -d` functionality for *modern* Go. If you're looking for a helper tool to organize Go projects under `${GOPATH}`, this is for you!

The classic `go get -d` used to clone the import repository and build the application right away, placing the executable in the `${GOPATH}/bin` directory. Nowadays, you would use `go install` if the goal is only to install the application executable.

Therefore, `go-get-d` is most useful for a developer to quickly get started with a Go repository, using the "classic" `${GOPATH}` organization style.

## Installation

Use `go install` as the following example:


```sh
go install github.com/otaviof/go-get-d@latest
```

When working on this repository you can alternatively use the `install` target, i.e.:

```sh
make install
```

The executable is placed on `${GOPATH}/bin`.

## Usage

The usage is straightforward, the only input required is the import name. For instance:

```bash
go-get-d github.com/otaviof/go-get-d
```

## Shell Integration

A practical way to clone the import repository and enter the directory, is using a shell function. The function will call `go-get-d` with the `--path` flag and change the directory for you.

### Bash/ZSH

Add the [`go_get_d` function](./go_get_d.sh) to your `.bashrc` or `.zshrc` file, alternatively, you can source the [`gogetd.sh`](./go_get_d.sh).

After including the function, or sourcing the file, you can use it like this:

```sh
go_get_d github.com/otaviof/go-get-d
```

The shell function will change into project's directory:

```
$ go_get_d github.com/otaviof/go-get-d
$ pwd
/home/otaviof/go/src/github.com/otaviof/go-get-d
```

### Eval

You can also use the `eval` command to execute the `cd` directly in your shell:

```sh
eval $(go-get-d github.com/otaviof/go-get-d)
```

## Inspect Import

You can additionally *inspect* the given import searching for the `main` package, when found it will be subject to `go build`, creating the application executable. Please consider the usage of the `--inspect` flag:

```bash
go-get-d --inspect github.com/otaviof/go-get-d
```

The `--inspect` flags makes it act as `go install`, therefore the application executable is stored on the `${GOPATH}/bin`.

[workflowsTest]: https://github.com/otaviof/go-get-d/actions/workflows/test.yaml
[workflowsTestBadge]: https://github.com/otaviof/go-get-d/actions/workflows/test.yaml/badge.svg
[goReportCard]: https://goreportcard.com/report/github.com/otaviof/go-get-d
[goReportCardBadge]: https://goreportcard.com/badge/github.com/otaviof/go-get-d
