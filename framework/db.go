package framework

import (
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/ninjawarrior1337/hanamaru-go/db"
)

func (h *Hanamaru) SetupDB() (err error) {
	if os.Getenv("IN_DOCKER") == "true" {
		h.Db, err = db.InitDB("/data/hanamaru.sqlite")
	} else {
		h.Db, err = db.InitDB("hanamaru.sqlite")
	}
	if err != nil {
		return
	}
	err = db.RunMigrations(h.Db)
	if err != migrate.ErrNoChange {
		return
	}
	h.AddCommand(sqlCommand)
	return nil
}

var sqlCommand = &Command{
	Name:               "sql",
	PermissionRequired: 0,
	OwnerOnly:          true,
	Help:               "Run SQL commands against the bot's DB",
	Exec: func(ctx *Context) error {
		query := ctx.TakeRest()

		var result [][]string
		rows, err := ctx.Hanamaru.Db.Query(query)
		if err != nil {
			return err
		}
		cols, err := rows.Columns()
		if err != nil {
			return err
		}
		pointers := make([]interface{}, len(cols))
		container := make([]string, len(cols))

		for i := range pointers {
			pointers[i] = &container[i]
		}

		for rows.Next() {
			rows.Scan(pointers...)
			result = append(result, container)
		}

		var s []string
		s = append(s, strings.Join(cols, "\t"))
		for _, r := range result {
			s = append(s, strings.Join(r, "\t"))
		}
		ctx.Reply(strings.Join(s, "\n"))

		return nil
	},
}
