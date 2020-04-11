package fun

import (
	"encoding/json"
	"github.com/markbates/pkger"
	"hanamaru/hanamaru"
	"strings"
)

var dmappings map[string]string

func init() {
	file, _ := pkger.Open("/assets/dance.json")
	defer file.Close()
	json.NewDecoder(file).Decode(&dmappings)
}

var Dance = &hanamaru.Command{
	Name:               "dance",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		targetSlice := strings.Split(ctx.Args[0], "")
		for _, char := range targetSlice {
			ctx.Reply(dmappings[char])
		}
		return nil
	},
}
