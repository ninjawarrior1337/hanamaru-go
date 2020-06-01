package fun

import (
	"bytes"
	"encoding/json"
	"github.com/markbates/pkger"
	"hanamaru/hanamaru"
	"html/template"
	"log"
	"math/rand"
)

var quotes []string
var quoteTempl = template.Must(template.New("suntsu").Parse(`"{{.}}" - Sun Tsu, Art of War`))

func init() {
	file, err := pkger.Open("/assets/suntsu.json")
	if err != nil {
		log.Fatalln(err)
	}
	json.NewDecoder(file).Decode(&quotes)
}

var Suntsu = &hanamaru.Command{
	Name:               "suntsu",
	PermissionRequired: 0,
	Exec: func(ctx *hanamaru.Context) error {
		buff := new(bytes.Buffer)
		quoteTempl.Execute(buff, quotes[rand.Intn(len(quotes))])
		ctx.Reply(buff.String())
		return nil
	},
}
