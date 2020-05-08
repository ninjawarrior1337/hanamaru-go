package info

import (
	"bytes"
	"fmt"
	"hanamaru/hanamaru"
	"html/template"
)

var InviteTemplate = template.Must(template.New("invite").Parse(`https://discordapp.com/oauth2/authorize?client_id={{.}}&scope=bot&permissions=8`))

var Invite = &hanamaru.Command{
	Name:               "invite",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *hanamaru.Context) error {
		var inviteString bytes.Buffer
		InviteTemplate.Execute(&inviteString, ctx.State.User.ID)
		ctx.Reply(fmt.Sprintf("Have fun with this: <%v>", inviteString.String()))
		return nil
	},
}
