package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/ninjawarrior1337/hanamaru-go/db"
)

type User struct {
	name string
	age  int
}

func main() {
	os.Remove("hanamaru.test.db")

	doChecks := func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	}

	sql, err := db.InitDB("hanamaru.test.db")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sql.Close()

	err = db.RunMigrations(sql)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db.MutateAwardCount(sql, "guild_1", "user_1", "test", db.AwardIncrement)
	db.MutateAwardCount(sql, "guild_1", "user_1", "test", db.AwardIncrement)
	db.MutateAwardCount(sql, "guild_1", "user_1", "test", db.AwardIncrement)
	db.MutateAwardCount(sql, "guild_1", "user_1", "test", db.AwardDecrement)
	db.MutateAwardCount(sql, "guild_1", "user_2", "test", db.AwardIncrement)
	db.MutateAwardCount(sql, "guild_1", "user_3", "test", db.AwardIncrement)

	db.MutateAwardCount(sql, "guild_1", "user_1", "test2", db.AwardIncrement)
	db.MutateAwardCount(sql, "guild_1", "user_2", "test2", db.AwardIncrement)
	db.MutateAwardCount(sql, "guild_1", "user_3", "test2", db.AwardIncrement)
	db.MutateAwardCount(sql, "guild_1", "user_1", "test2", db.AwardDecrement)
	db.MutateAwardCount(sql, "guild_1", "user_2", "test2", db.AwardIncrement)
	db.MutateAwardCount(sql, "guild_1", "user_3", "test2", db.AwardIncrement)
	db.MutateAwardCount(sql, "guild_1", "user_3", "test2", db.AwardIncrement)
	doChecks(err)

	lb, _ := db.QueryLeaderboard(sql, "guild_1")

	spew.Dump(lb)
}
