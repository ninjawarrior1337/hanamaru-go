package db

import (
	"database/sql"
	"errors"
	"fmt"
)

func MutateAddTag(db *sql.DB, tag string, guild_id string, channel_id string, message_id string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO tag(name, guild_id, channel_id, message_id) VALUES (?, ?, ?, ?)", tag, guild_id, channel_id, message_id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

type TagQueryResult struct {
	Name      string
	GuildID   string
	ChannelID string
	MessageID string
}

func QueryTagByName(db *sql.DB, name string, guild_id string) (TagQueryResult, error) {
	queryRes := db.QueryRow("SELECT name, guild_id, channel_id, message_id FROM tag WHERE guild_id = ? AND name = ?", guild_id, name)
	if queryRes.Err() != nil {
		return TagQueryResult{}, queryRes.Err()
	}
	var res TagQueryResult
	err := queryRes.Scan(&res.Name, &res.GuildID, &res.ChannelID, &res.MessageID)
	if err == sql.ErrNoRows {
		return TagQueryResult{}, errors.New("no tag found with specified name")
	}
	return res, nil
}

func QuerySearchTags(db *sql.DB, query string, guild_id string) ([]string, error) {
	var results []string
	res, err := db.Query("SELECT name FROM tag WHERE guild_id = ? AND name LIKE ?", guild_id, fmt.Sprintf("%%%v%%", query))
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var s string
		res.Scan(&s)
		results = append(results, s)
	}

	return results, nil
}

func QueryListTags(db *sql.DB, guild_id string, page int) ([]string, error) {
	const LIMIT = 20
	var results []string
	// Page here is 0 indexed
	res, err := db.Query("SELECT name FROM tag WHERE guild_id = ? LIMIT ? OFFSET ?", guild_id, LIMIT, LIMIT*page)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var s string
		res.Scan(&s)
		results = append(results, s)
	}

	return results, nil
}

func MutateRemoveTagByName(db *sql.DB, guild_id string, tag string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	r, err := tx.Exec("DELETE FROM tag WHERE name = ? AND guild_id = ?", tag, guild_id)
	if err != nil {
		return err
	}

	if n, err := r.RowsAffected(); n == 0 || err != nil {
		tx.Rollback()
		return errors.New("failed to delete tag")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
