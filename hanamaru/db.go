package hanamaru

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
	"os"
)

var configCommand = &Command{
	Name:               "config",
	PermissionRequired: 0,
	OwnerOnly:          false,
	Help:               "",
	Exec: func(ctx *Context) error {
		err := ctx.Hanamaru.db.Update(func(tx *bolt.Tx) error {
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
				ctx.Set(key, val)
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
		h.db, err = bolt.Open("/data/db.bbolt", 0666, nil)
		if err != nil {
			return
		}
	}
	h.db, err = bolt.Open("db.bbolt", 0666, nil)
	h.AddCommand(configCommand)
	if err != nil {
		return
	}
	return nil
}

func (c *Context) Set(key, value string) {
	c.Hanamaru.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(c.GuildID))
		if err != nil {
			return err
		}
		return b.Put([]byte(key), []byte(value))
	})
}

func (c *Context) Get(key string) string {
	var vCopy []byte
	c.Hanamaru.db.View(func(tx *bolt.Tx) error {
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
