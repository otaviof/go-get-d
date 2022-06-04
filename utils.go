package main

import (
	"net/url"
	"path"
	"strings"
)

func urlToGoModule(u *url.URL) string {
	urlPath := strings.TrimSuffix(u.Path, ".git")
	return path.Join(u.Host, urlPath)
}
