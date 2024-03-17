package events

import (
	"database/sql"
	"fmt"
	"testing"

	hdb "github.com/ninjawarrior1337/hanamaru-go/db"
)

var db *sql.DB

const Chad AwardType = "chad_award"

func init() {
	db, _ = sql.Open("sqlite3", ":memory:")
	err := hdb.RunMigrations(db)
	if err != nil {
		fmt.Println(err)
	}
}

func TestAddPointToUser(t *testing.T) {
	guild_id := "guild_1"
	uid := "1000001"
	err := hdb.MutateAwardCount(db, guild_id, uid, Chad.MustName(), hdb.AwardIncrement)
	if err != nil {
		t.Error(err)
		return
	}

	a, _ := hdb.QueryAwardCounts(db, guild_id, uid)
	fmt.Println(a)

	if a, _ := hdb.QueryAwardCounts(db, guild_id, uid); a[Chad.MustName()] != 1 {
		t.Fail()
	}

	hdb.MutateAwardCount(db, guild_id, uid, Chad.MustName(), hdb.AwardIncrement)
	if a, _ := hdb.QueryAwardCounts(db, guild_id, uid); a[Chad.MustName()] != 2 {
		t.Fail()
	}

	hdb.MutateAwardCount(db, guild_id, uid, Chad.MustName(), hdb.AwardDecrement)
	if a, _ := hdb.QueryAwardCounts(db, guild_id, uid); a[Chad.MustName()] != 1 {
		t.Fail()
	}
	db.Close()
}

func TestGetUserAwards(t *testing.T) {
	a, _ := hdb.QueryAwardCounts(db, "guild_1", "1000001")
	fmt.Println(a)
}
