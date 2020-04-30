// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
	"os"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

var TAGS = []string{"", "jp", "ij"}
var OSes = []string{"windows", "linux"}

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps, Test, Generate)
	fmt.Println("Building...")

	for _, cOS := range OSes {
		for _, tag := range TAGS {
			fmt.Println("Generating hanamaru OS: " + cOS + " TAG: " + tag)
			fileName := "hanamaru-" + cOS
			if tag != "" {
				fileName += "-" + tag
			}
			switch cOS {
			case "windows":
				fileName += ".exe"
			default:
				fileName += ""
			}
			env := map[string]string{
				"GOOS":   cOS,
				"GOARCH": "amd64",
			}
			cmd := sh.RunWith(env, "go", "build", "-tags", tag, `-ldflags=-s -w`, "-o", "artifacts/"+fileName)

			if err := cmd; err != nil {
				return err
			}
		}
	}
	return nil
}

func Test() error {
	fmt.Println("Running tests...")
	return sh.Run("go", "test", "./...")
}

func Generate() error {
	fmt.Println("Running Go generate...")
	return sh.Run("go", "generate")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	return sh.Run("go", "get", "github.com/markbates/pkger/cmd/pkger")
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("artifacts")
}
