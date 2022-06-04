package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const selfModuleName = "github.com/otaviof/go-get-d"

func TestGoGetD_ParseURL(t *testing.T) {
	tests := []struct {
		name    string // test case short description
		input   string // raw input given to GoGetD
		module  string // expected module name parsed
		wantErr bool   // error is expected on the test case
	}{{
		name:    "Regular GitHub HTTP repository",
		input:   "https://github.com/otaviof/go-get-d.git",
		module:  selfModuleName,
		wantErr: false,
	}, {
		name:    "Go Module",
		input:   selfModuleName,
		module:  selfModuleName,
		wantErr: false,
	}, {
		name:    "Invalid URL",
		input:   "something invalid...",
		wantErr: true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGoGetD(tt.input)
			if err := g.ParseURL(); (err != nil) != tt.wantErr {
				t.Errorf("GoGetD.ParseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if tt.module != g.module {
					t.Errorf("GoGetD.ParseURL() module name mismatch = %q, want %q",
						g.module, tt.module)
				}
				if g.repositoryURL == nil {
					t.Error("GoGetD.ParseURL() g.repositoryURL is nil")
				}
			}
		})
	}
}

func TestGoGetD_ModuleDir(t *testing.T) {
	g := NewGoGetD(selfModuleName)
	g.module = selfModuleName

	t.Run("LookupModuleDirInGopath", func(t *testing.T) {
		err := g.LookupModuleDirInGopath()
		require.NoError(t, err)
		require.NotEmpty(t, g.dir)
	})

	t.Run("ModuleDirExists", func(t *testing.T) {
		exists := g.ModuleDirExits()
		require.True(t, exists)
	})
}
