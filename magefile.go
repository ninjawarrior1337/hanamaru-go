// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
	"strings"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

var TAGS = []string{"", "jp,ij"}
var OSes = []string{"windows", "linux"}

func BuildCI() {
	mg.SerialDeps(InstallDeps, Pkger)
	build()
}

func Build() {
	mg.SerialDeps(InstallDeps, Test, Pkger)
	build()
}

// A build step that requires additional params, or platform specific steps for example
func build() error {

	fmt.Println("Building...")

	for _, cOS := range OSes {
		for _, tag := range TAGS {
			fmt.Println("Generating hanamaru: OS: " + cOS + " TAG: " + tag)
			fileName := "hanamaru-" + cOS
			if tag != "" {
				fileName += "-" + strings.ReplaceAll(tag, ",", "-")
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

func BuildDocker() error {
	mg.SerialDeps(InstallDeps, Pkger)
	fmt.Println("Building for Docker")
	return sh.RunWith(map[string]string{"CGO_ENABLED": "0"}, "go", "build", "-tags", "ij,jp", `-ldflags=-s -w`, "-o", "hanamaru")
}

func Test() error {
	fmt.Println("Running tests...")
	return sh.Run("go", "test", "./...")
}

func Pkger() error {
	mg.SerialDeps(LinkCommands)
	fmt.Println("Packaging assets with pkger...")
	err := sh.Run("pkger", "list")
	if err != nil {
		return err
	}
	return sh.Run("pkger")
}

func LinkCommands() error {
	fmt.Println("Linking Commands...")
	return sh.Run("go", "run", filepath.FromSlash("./tools/cmd/gen_command_imports.go"))
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
	os.Remove("pkged.go")
}
