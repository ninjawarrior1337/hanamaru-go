package fun

import (
	"bytes"
	"encoding/json"
	"github.com/markbates/pkger"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	"html/template"
	"math/rand"
)

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
		file, err := pkger.Open("/assets/suntsu.json")
		if err != nil {
			return err
		}
		json.NewDecoder(file).Decode(&quotes)
		return nil
	},
}
