package events

import (
	"encoding/json"
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/ninjawarrior1337/hanamaru-go/framework"
	bolt "go.etcd.io/bbolt"
	"strings"
)

type AwardType string

func (a AwardType) IsValid() error {
	if len(strings.Split(string(a), "_")) == 2 {
		if strings.Split(string(a), "_")[1] == "award" {
			return nil
		}
	}
	return errors.New("invalid award type")
}

func (a AwardType) Name() (string, error) {
	if a.IsValid() != nil {
		return "", a.IsValid()
	}

	return strings.Split(string(a), "_")[0], nil
}

type Operation int

const (
	Add Operation = iota
	Subtract
	Set
)

var AwardBucket = []byte("awards")

type AwardKV map[AwardType]int

// ModifyUserAwardCount will update the award information in a DB
func ModifyUserAwardCount(db *bolt.DB, award AwardType, userID string, operation Operation, amount int) {
	db.Update(func(tx *bolt.Tx) error {
		awardB, _ := tx.CreateBucketIfNotExists(AwardBucket)
		data := awardB.Get([]byte(userID))

		var playerInfo AwardKV
		json.Unmarshal(data, &playerInfo)
		if playerInfo == nil {
			playerInfo = make(AwardKV)
		}

		switch operation {
		case Add:
			playerInfo[award] = playerInfo[award] + amount
		case Subtract:
			playerInfo[award] = playerInfo[award] - amount
		case Set:
			playerInfo[award] = amount
		}

		modifiedData, _ := json.Marshal(&playerInfo)

		awardB.Put([]byte(userID), modifiedData)
		return nil
	})
}

func GetAllAwardsInfo(db *bolt.DB) map[string]AwardKV {
	var final = make(map[string]AwardKV)
	var users = make([]string, 0)
	db.Update(func(tx *bolt.Tx) error {
		awardB, _ := tx.CreateBucketIfNotExists(AwardBucket)
		awardB.ForEach(func(k, v []byte) error {
			users = append(users, string(k))
			return nil
		})
		return nil
	})
	for _, u := range users {
		final[u] = GetUserAwards(db, u)
	}
	return final
}

func GetUserAwards(db *bolt.DB, userID string) AwardKV {
	var playerInfo AwardKV
	db.Update(func(tx *bolt.Tx) error {
		awardB, _ := tx.CreateBucketIfNotExists(AwardBucket)
		data := awardB.Get([]byte(userID))

		json.Unmarshal(data, &playerInfo)
		if playerInfo == nil {
			playerInfo = make(AwardKV)
		}

		return nil
	})
	return playerInfo
}

var AwardsAddHandler = &framework.EventListener{
	Name: "Award Add",
	HandlerConstructor: func(h *framework.Hanamaru) interface{} {
		return func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
			award := AwardType(r.Emoji.Name)
			m, _ := s.ChannelMessage(r.ChannelID, r.MessageID)
			if award.IsValid() == nil {
				ModifyUserAwardCount(h.Db, award, m.Author.ID, Add, 1)
			}
		}
	},
}

var AwardsRemoveHandler = &framework.EventListener{
	Name: "Award Remove",
	HandlerConstructor: func(h *framework.Hanamaru) interface{} {
		return func(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
			award := AwardType(r.Emoji.Name)
			m, _ := s.ChannelMessage(r.ChannelID, r.MessageID)
			if award.IsValid() == nil {
				ModifyUserAwardCount(h.Db, award, m.Author.ID, Subtract, 1)
			}
		}
	},
}
