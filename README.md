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

```bash
go-get-d [import]
```

For example:

```
$ go-get-d github.com/otaviof/go-get-d
# Go Module: "github.com/otaviof/go-get-d"
# Directory: "${GOPATH}/src/github.com/otaviof/go-get-d"
# Cloning repository...
$ git clone https://github.com/otaviof/go-get-d ~/go/src/github.com/otaviof/go-get-d
Cloning into '/Users/otaviof/go/src/github.com/otaviof/go-get-d'...
warning: redirecting to https://github.com/otaviof/go-get-d/
remote: Enumerating objects: 3532, done.
remote: Counting objects: 100% (152/152), done.
remote: Compressing objects: 100% (91/91), done.
remote: Total 3532 (delta 86), reused 91 (delta 52), pack-reused 3380
Receiving objects: 100% (3532/3532), 1.40 MiB | 11.77 MiB/s, done.
Resolving deltas: 100% (2208/2208), done.
# Inspecting Go package...
# All done!
```

And then:

```
$ go-get-d github.com/otaviof/go-get-d
# Go Module: "github.com/otaviof/go-get-d"
# Directory: "${GOPATH}/src/github.com/otaviof/go-get-d"
# Inspecting Go package...
# All done!
```

[projectBadgeSVG]: https://github.com/otaviof/go-get-d/actions/workflows/test.yaml/badge.svg
[projectTestWorkflow]: https://github.com/otaviof/go-get-d/actions/workflows/test.yaml
