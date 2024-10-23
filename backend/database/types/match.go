package types

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}

	return json.Marshal(nil)
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	var s *string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if s != nil {
		ns.Valid = true
		ns.String = *s
	} else {
		ns.Valid = false

	}

	return nil
}

type NullTime struct {
	sql.NullTime
}

func (ns NullTime) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Time)
	}

	return json.Marshal(nil)
}

type Match struct {
	ID        int        `db:"id" json:"id"`
	HomeTeam  string     `db:"hometeam" json:"home_team"`
	AwayTeam  string     `db:"awayteam" json:"away_team"`
	HomeScore int        `db:"homescore" json:"home_score"`
	AwayScore int        `db:"awayscore" json:"away_score"`
	Status    NullString `db:"status" json:"status"`
	CreatedAt NullTime   `db:"created_at" json:"created_at"`
}

type MatchCreate struct {
	HomeTeam  string     `json:"home_team"`
	AwayTeam  string     `json:"away_team"`
	HomeScore int        `json:"home_score"`
	AwayScore int        `json:"away_score"`
	Status    NullString `json:"status"`
}

type MatchUpdate struct {
	ID        int        `json:"id"`
	HomeScore int        `json:"home_score"`
	AwayScore int        `json:"away_score"`
	Status    NullString `json:"status"`
}
