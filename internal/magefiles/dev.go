//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	"github.com/GuanceCloud/guance-scenarios/internal/chore"
)

type Dev mg.Namespace

// Lint run the linter
func (ns Dev) Lint() error {
	argList := [][]string{
		{"golangci-lint", "run", "./..."},
		{"markdownlint", "-i", "docs/references", "-f", "."},
		{"gofumpt", "-l", "-e", "."},
	}
	return chore.BatchRunV(argList)
}

// Fmt run the formatter
func (ns Dev) Fmt() error {
	argList := [][]string{
		{"golangci-lint", "run", "--fix", "./..."},
		{"gofumpt", "-l", "-w", "."},
		{"goimports", "-w", "."},
		{"prettier", "-w", "**/*.md"},
	}
	return chore.BatchRunV(argList)
}

// D2 build svg from d2 files
func (ns Dev) D2() error {
	files, err := chore.ListFileByExt(".", "d2")
	if err != nil {
		return err
	}
	var mErr error
	for _, d2File := range files {
		outFile := d2File[:len(d2File)-3] + ".svg"
		if err := sh.RunV("d2", "--sketch", "-t", "0", d2File, outFile); err != nil {
			mErr = fmt.Errorf("d2 svg %s: %w", d2File, err)
		}
	}
	return mErr
}
