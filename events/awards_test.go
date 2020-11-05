package events

import (
	"fmt"
	"go.etcd.io/bbolt"
	"os"
	"testing"
)

var db *bbolt.DB

const Chad AwardType = "chad_award"

func init() {
	os.Remove("bbolt.test.db")
	db, _ = bbolt.Open("bbolt.test.db", 0600, nil)
}

func TestAddPointToUser(t *testing.T) {
	ModifyUserAwardCount(db, Chad, "1000001", Add, 5)

	if a := GetUserAwards(db, "1000001"); a[Chad] != 5 {
		t.Fail()
	}

	ModifyUserAwardCount(db, Chad, "1000001", Add, 1)
	if a := GetUserAwards(db, "1000001"); a[Chad] != 6 {
		t.Fail()
	}

	ModifyUserAwardCount(db, Chad, "1000001", Subtract, 1)
	if a := GetUserAwards(db, "1000001"); a[Chad] != 5 {
		t.Fail()
	}

	ModifyUserAwardCount(db, Chad, "1000001", Set, 0)
	if a := GetUserAwards(db, "1000001"); a[Chad] != 0 {
		t.Fail()
	}
}

func TestGetUserAwards(t *testing.T) {
	a := GetUserAwards(db, "1")
	fmt.Println(a)
}
