package hanamaru

import (
	bolt "go.etcd.io/bbolt"
	"os"
)

func (h *Hanamaru) SetupDB() (err error) {
	if os.Getenv("IN_DOCKER") == "true" {
		h.db, err = bolt.Open("/data/db.bbolt", 0666, nil)
		if err != nil {
			return
		}
	}
	h.db, err = bolt.Open("db.bbolt", 0666, nil)
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
