package fun

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"text/template"

	_ "embed"

	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

//go:embed assets/suntsu.json
var suntsuJSON []byte

var quotes []string
var quoteTempl = template.Must(template.New("suntsu").Parse(`"{{.}}" - Sun Tsu, Art of War`))

var Suntsu = &framework.Command{
	Name:               "suntsu",
	PermissionRequired: 0,
	Exec: func(ctx *framework.Context) error {
		buff := new(bytes.Buffer)
		quoteTempl.Execute(buff, quotes[rand.Intn(len(quotes))])
		ctx.Reply(buff.String())
		return nil
	},
	Setup: func() error {
		json.NewDecoder(bytes.NewReader(suntsuJSON)).Decode(&quotes)
		return nil
	},
}
