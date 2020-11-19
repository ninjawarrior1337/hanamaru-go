package fun

import (
	"fmt"
	"math"

	"github.com/ninjawarrior1337/hanamaru-go/framework"
)

const SpacesMax = 32

var SineCommand = &framework.Command{
	Name: "sine",
	Help: "WaVeS",
	Exec: func(ctx *framework.Context) error {
		var output = ""
		var text = ctx.TakeRest()

		for i := 0; i <= SpacesMax*2; i++ {
			output += fmt.Sprintf("%v%v\n", computeSpaces(i, SpacesMax), text)
		}
		_, err := ctx.Reply(output)
		return err
	},
}

func computeSpaces(num int, max int) string {
	var spaces = ""
	numSpaces := int(
		float64(1) / float64(2) * (float64(-max)*
			math.Cos(
				float64(1)/float64(max)*
					math.Pi*
					float64(num)) +
			float64(max)))
	for i := 0; i < numSpaces; i++ {
		spaces += " "
	}
	return spaces
}
