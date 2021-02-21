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
	mg.SerialDeps(Generate)
	build()
}

func Build() {
	mg.SerialDeps(Test, Generate)
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
				"GOOS":        cOS,
				"GOARCH":      "amd64",
				"CGO_ENABLED": "0",
			}
			cmd := sh.RunWith(env, "go",
				"build",
				"-tags",
				tag,
				"-ldflags="+flags(),
				"-o",
				"artifacts/"+fileName,
			)

			if err := cmd; err != nil {
				return err
			}
		}
	}
	return nil
}

func BuildDocker() error {
	mg.SerialDeps(Generate)
	fmt.Println("Building for Docker")
	flags := []string{
		"build",
		"-a",
		"-tags",
		"jp,ij",
		`-ldflags=-s -w`,
		"-o",
		"hanamaru",
	}
	return sh.RunWith(map[string]string{"CGO_ENABLED": "0"}, "go", flags...)
}

func Test() error {
	fmt.Println("Running tests...")
	return sh.Run("go", "test", "./...")
}

func Generate() error {
	mg.SerialDeps(linkCommands)
	return nil
}

func linkCommands() error {
	fmt.Println("Linking Commands...")
	return sh.Run("go", "run", filepath.FromSlash("./tools/cmd/gen_command_imports.go"))
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("artifacts")
	os.Remove("pkged.go")
}

func flags() string {
	commitHash := os.Getenv("GITHUB_SHA")
	buildDate, _ := sh.OutputWith(map[string]string{
		"TZ": "America/Los_Angeles",
	}, "date")
	return fmt.Sprintf(`-s -w -X "github.com/ninjawarrior1337/hanamaru-go/commands/info.CommitHash=%v" -X "github.com/ninjawarrior1337/hanamaru-go/commands/info.BuildDate=%v"`, commitHash, buildDate)
}
