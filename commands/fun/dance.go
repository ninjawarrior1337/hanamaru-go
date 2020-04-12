package fun

import (
	"encoding/json"
	"fmt"
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
		input, err := ctx.GetArgIndex(0)
		if err != nil {
			return fmt.Errorf("please type something in (if it has spaces use quotes)")
		}

		input = strings.ToUpper(input)

		targetSlice := strings.Split(input, "")
		for _, char := range targetSlice {
			ctx.Reply(dmappings[char])
		}
		return nil
	},
}
