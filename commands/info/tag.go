package info

import (
	"errors"
	"fmt"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	bolt "go.etcd.io/bbolt"
)

var Tag = &framework.Command{
	Name:               "tag",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "Let's you tag pieces of text: <get|add|delete> <key> <value>",
	Exec: func(ctx *framework.Context) error {
		db := ctx.Hanamaru.Db

		return db.Update(func(tx *bolt.Tx) error {
			b, _ := tx.CreateBucketIfNotExists([]byte(ctx.GuildID))
			tb, _ := b.CreateBucketIfNotExists([]byte("tags"))
			tagOp := ctx.GetArgIndexDefault(0, "get")
			switch tagOp {
			case "get":
				key := ctx.GetArgIndexDefault(1, "")
				if key == "" {
					output := "```"
					tb.ForEach(func(k, v []byte) error {
						output += fmt.Sprintf("%s: %s\n", k, v)
						return nil
					})
					output += "```"
					if output == "" {
						return errors.New("no tags saved")
					}
					ctx.Reply(output)
				} else {
					v := tb.Get([]byte(key))
					ctx.Reply(fmt.Sprintf("%s -> %s", key, v))
				}
			case "delete":
				label, err := ctx.GetArgIndex(0)
				if err != nil {
					return err
				}
				err = tb.Delete([]byte(label))
				if err != nil {
					return err
				}
			default:
				label, err := ctx.GetArgIndex(0)
				if err != nil {
					return err
				}
				val, err := ctx.GetArgIndex(1)
				if err != nil {
					return err
				}
				err = tb.Put([]byte(label), []byte(val))
				if err != nil {
					return err
				}
				ctx.Reply(fmt.Sprintf("Set: %s -> %s", label, val))
			}
			return nil
		})
	},
}
