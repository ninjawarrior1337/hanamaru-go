package framework

import (
	"fmt"
	"os"

	bolt "go.etcd.io/bbolt"
)

var configCommand = &Command{
	Name:               "config",
	PermissionRequired: 0,
	OwnerOnly:          true,
	Help:               "Manage the key-value store of the bot",
	Exec: func(ctx *Context) error {
		err := ctx.Hanamaru.Db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte(ctx.GuildID))
			if err != nil {
				return err
			}
			subCommand := ctx.GetArgIndexDefault(0, "get")
			switch subCommand {
			case "get":
				key := ctx.GetArgIndexDefault(1, "")
				if key == "" {
					output := ""
					b.ForEach(func(k, v []byte) error {
						output += fmt.Sprintf("**%s**: %s\n", k, v)
						return nil
					})
					if output == "" {
						ctx.Reply("No values saved on this server")
						return nil
					}
					ctx.Reply(output)
				} else {
					v := b.Get([]byte(key))
					ctx.Reply(fmt.Sprintf("%s -> %s", key, v))
				}
			case "set":
				key, err := ctx.GetArgIndex(1)
				if err != nil {
					return err
				}
				val, err := ctx.GetArgIndex(2)
				if err != nil {
					return err
				}
				err = ctx.setFromTx(tx, key, val)
				if err != nil {
					return err
				}
				ctx.Reply(fmt.Sprintf("Attempted to set: %s -> %s", key, val))
			}
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	},
}

func (h *Hanamaru) SetupDB() (err error) {
	if os.Getenv("IN_DOCKER") == "true" {
		h.Db, err = bolt.Open("/data/db.bbolt", 0666, nil)
		return
	}
	h.Db, err = bolt.Open("db.bbolt", 0666, nil)
	h.AddCommand(configCommand)
	if err != nil {
		return
	}
	return nil
}

func (c *Context) Set(key, value string) error {
	return c.Hanamaru.Db.Update(func(tx *bolt.Tx) error {
		return c.setFromTx(tx, key, value)
	})
}

func (c *Context) setFromTx(tx *bolt.Tx, key, value string) error {
	b, err := tx.CreateBucketIfNotExists([]byte(c.GuildID))
	if err != nil {
		return err
	}
	return b.Put([]byte(key), []byte(value))
}

func (c *Context) Get(key string) string {
	var vCopy []byte
	c.Hanamaru.Db.View(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(c.GuildID))
		if err != nil {
			return err
		}
		v := b.Get([]byte(key))
		copy(vCopy, v)
		return nil
	})
	return string(vCopy)
}
