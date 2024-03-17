package db

import "database/sql"

type AwardOperation int

const (
	AwardIncrement AwardOperation = iota
	AwardDecrement
)

type LeaderboardEntry struct {
	AwardName string
	EarnerId  string
	Count     int
}

type AwardKV map[string]int

func QueryLeaderboard(db *sql.DB, guild_id string) ([]LeaderboardEntry, error) {
	entries := []LeaderboardEntry{}
	res, err := db.Query(`select award_name, earner_id, max(count) from award where guild_id = ? group by award_name`, guild_id)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var leaderboardEntry LeaderboardEntry
		res.Scan(&leaderboardEntry.AwardName, &leaderboardEntry.EarnerId, &leaderboardEntry.Count)
		entries = append(entries, leaderboardEntry)
	}
	res.Close()

	return entries, nil
}

func QueryAwardCounts(db *sql.DB, guild_id string, earner_id string) (AwardKV, error) {
	kv := make(AwardKV)
	res, err := db.Query(`SELECT award_name, count FROM award WHERE guild_id = ? AND earner_id = ?`, guild_id, earner_id)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var award_name string
		var count int
		res.Scan(&award_name, &count)
		kv[award_name] = count
	}
	res.Close()

	return kv, nil
}

func MutateAwardCount(db *sql.DB, guild_id string, earner_id string, award_name string, op AwardOperation) error {
	var err error

	switch op {
	case AwardIncrement:
		_, err = db.Exec(`
			INSERT INTO award(guild_id, earner_id, award_name) 
			VALUES (?, ?, ?)
			ON CONFLICT(guild_id, earner_id, award_name)
			DO UPDATE SET count = count+1
			`, guild_id, earner_id, award_name)
	case AwardDecrement:
		_, err = db.Exec(`
			INSERT INTO award(guild_id, earner_id, award_name) 
			VALUES (?, ?, ?)
			ON CONFLICT(guild_id, earner_id, award_name)
			DO UPDATE SET count = count-1
			`, guild_id, earner_id, award_name)
	}

	if err != nil {
		return err
	}

	return nil
}
