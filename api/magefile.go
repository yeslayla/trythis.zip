//go:build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
)

// Runs go mod download and then installs the binary.
func Build() error {
	mg.Deps(Generate)
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "build", "-o", "bin/app", ".")
}

// Run builds and runs the executable
func Run() error {
	mg.Deps(Build)
	return sh.RunV("bin/app")
}

// Test runs `go test` with coverage on all packages
func Test() error {
	fmt.Println("Running tests...")
	return sh.RunV("go", "test", "-coverprofile=coverage.out", "./...")
}

// Clean runs `go clean`, deletes old builds, and removes generated files
func Clean() error {
	if err := sh.RunV("go", "clean"); err != nil {
		return err
	}

	if err := os.RemoveAll("bin"); err != nil {
		return err
	}

	files, err := filepath.Glob("api/**.gen.go")
	if err != nil {
		return err
	}

	for _, f := range files {
		if err := os.Remove(f); err != nil {
			fmt.Printf("Removing '%s'", f)
			return err
		}
	}

	return nil
}

// Generate creates go files from the openapi file
func Generate() error {

	if _, err := os.Stat("api"); os.IsNotExist(err) {
		if err := os.Mkdir("api", os.ModePerm); err != nil {
			return err
		}
	}

	for _, v := range []string{"gorilla", "types", "spec"} {
		output, err := sh.Output("oapi-codegen", "-generate", v, "-package", "api", "openapi.yaml")
		if err != nil {
			return err
		}

		if err := os.WriteFile(fmt.Sprintf("api/%s.gen.go", v), []byte(output), 0644); err != nil {
			return err
		}
	}

	return nil
}
