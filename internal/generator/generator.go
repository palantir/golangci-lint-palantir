// Copyright 2026 Palantir Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate go run $GOFILE

package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"

	"golang.org/x/mod/modfile"
)

const (
	golangCILintModule = "github.com/golangci/golangci-lint/v2"
	goModPath          = "../../go.mod"
	distConfigPath     = "../../godel/config/dist-plugin.yml"
)

var golangCILintVersion = regexp.MustCompile(`(?m)(type: golangci-lint\n(?:[ \t].*\n)*?[ \t]+golangci-lint-version: )[^\s]+`)

// Updates the "golangci-lint-version:" specified in godel/config/dist-plugin.yml to match the version of the
// github.com/golangci/golangci-lint/v2 module specified in the go.mod file of this repository.
func main() {
	goModBytes, err := os.ReadFile(goModPath)
	if err != nil {
		panic(fmt.Errorf("read go.mod: %w", err))
	}

	goMod, err := modfile.Parse(goModPath, goModBytes, nil)
	if err != nil {
		panic(fmt.Errorf("parse go.mod: %w", err))
	}

	version := ""
	for _, requirement := range goMod.Require {
		if requirement.Mod.Path == golangCILintModule {
			version = requirement.Mod.Version
			break
		}
	}
	if version == "" {
		panic(fmt.Errorf("module %q is not required by go.mod", golangCILintModule))
	}

	distConfig, err := os.ReadFile(distConfigPath)
	if err != nil {
		panic(fmt.Errorf("read dist configuration: %w", err))
	}
	if matches := golangCILintVersion.FindAllIndex(distConfig, -1); len(matches) != 1 {
		panic(fmt.Errorf("expected exactly one golangci-lint version in %s, found %d", distConfigPath, len(matches)))
	}

	updated := golangCILintVersion.ReplaceAll(distConfig, []byte("${1}"+version))
	if bytes.Equal(updated, distConfig) {
		return
	}
	if err := os.WriteFile(distConfigPath, updated, 0o644); err != nil {
		panic(fmt.Errorf("write dist configuration: %w", err))
	}
}
