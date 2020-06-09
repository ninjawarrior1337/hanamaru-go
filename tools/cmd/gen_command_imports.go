package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var registerCommandTemplate = template.Must(template.New("register_commands").Parse(`// Code generated by framework command linker; DO NOT EDIT.
package main

import(
	{{- range $index, $e := .}}
	"github.com/ninjawarrior1337/hanamaru-go/commands/{{$index}}"
	{{- end}}
)

func init() {
	{{- range $i, $e := .}}
	{{- range $e}}
	commands = append(commands, {{$i}}.{{.}})
	{{- end}}
	{{end}}
}
`))

func main() {
	var commandDefs = map[string][]string{}

	err := filepath.Walk("commands", func(path string, info os.FileInfo, err error) error {
		var fset token.FileSet
		file, _ := parser.ParseFile(&fset, path, nil, parser.ParseComments)
		if file != nil {
			if len(file.Comments) > 0 {
				if strings.Contains(file.Comments[0].List[0].Text, "+build") {
					return nil
				}
			}
			ast.Inspect(file, func(node ast.Node) bool {
				d, ok := node.(*ast.ValueSpec)
				if ok {
					if len(d.Values) > 0 {
						for _, vs := range d.Values {
							ur, ok := vs.(*ast.UnaryExpr)
							if ok {
								x, ok := ur.X.(*ast.CompositeLit)
								if ok {
									if strings.Contains(fmt.Sprintf("%v", x.Type), "Command") {
										fmt.Println("Adding: " + d.Names[0].String())
										commandDefs[file.Name.String()] = append(commandDefs[file.Name.String()], d.Names[0].String())
									}
								}
							}
						}
					}
				}
				return true
			})
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.OpenFile("./commands.gen.go", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	f.Truncate(0)

	registerCommandTemplate.Execute(f, commandDefs)
}
