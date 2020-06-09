// +build ij

package events

import "github.com/ninjawarrior1337/hanamaru-go/util"

func init() {
	expansionMap["🇮🇳"] = append([]string{"🆗"}, util.MustMapToEmoji("shiraz")...)
	expansionMap["🇻🇳"] = append([]string{"🆗"}, util.MustMapToEmoji("ethan")...)
	expansionMap["🇵🇭"] = append([]string{"🆗"}, util.MustMapToEmoji("aidan")...)
	expansionMap["🇵🇪"] = append([]string{"🆗"}, util.MustMapToEmoji("jony")...)
	expansionMap["🇯🇵"] = append([]string{"🆗"}, util.MustMapToEmoji("tyler")...)
	expansionMap["🇮🇹"] = append([]string{"🆗"}, util.MustMapToEmoji("jaxon")...)
	expansionMap["⚫"] = append([]string{"🆗"}, util.MustMapToEmoji("bishop")...)
}
