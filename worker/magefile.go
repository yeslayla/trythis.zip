//go:build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
)

// Runs go mod download and then installs the binary.
func Build() error {
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "build", "-o", "bin/worker", ".")
}

// Run builds and runs the executable
func Run() error {
	mg.Deps(Build)
	return sh.RunV("bin/worker")
}

// Test runs `go test` with coverage on all packages
func Test() error {
	fmt.Println("Running tests...")
	return sh.RunV("go", "test", "-coverprofile=coverage.out", "./...")
}

// Clean runs `go clean` and deletes old builds
func Clean() error {
	if err := sh.RunV("go", "clean"); err != nil {
		return err
	}

	if err := os.RemoveAll("bin"); err != nil {
		return err
	}

	return nil
}
