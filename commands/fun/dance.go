package fun

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

//go:embed assets/dance.json
var danceBytes []byte

var dmappings map[string]string

var Dance = &framework.Command{
	Name:               "dance",
	PermissionRequired: 0,
	Exec: func(ctx *framework.Context) error {
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
	Setup: func() error {
		json.NewDecoder(bytes.NewReader(danceBytes)).Decode(&dmappings)
		return nil
	},
}
