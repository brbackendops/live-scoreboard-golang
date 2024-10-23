package match

import (
	"fmt"
	"log"
	"score/database/types"

	"github.com/jmoiron/sqlx"
)

type MatchRepository struct {
	db *sqlx.DB
}

func NewMatchRepo(db *sqlx.DB) *MatchRepository {
	return &MatchRepository{
		db,
	}
}

func (m *MatchRepository) GetMatches() ([]types.Match, error) {
	match := types.Match{}
	matches := []types.Match{}

	rows, err := m.db.Queryx("SELECT * FROM match")
	if err != nil {
		return []types.Match{}, err
	}

	for rows.Next() {
		err := rows.StructScan(&match)
		matches = append(matches, match)
		if err != nil {
			log.Fatal(err)
		}
	}

	return matches, nil
}

func (m *MatchRepository) GetMatch(id int) (types.Match, error) {
	match := types.Match{}

	err := m.db.Get(&match, "SELECT * FROM match WHERE id=$1 LIMIT 1", id)
	if err != nil {
		return types.Match{}, nil
	}

	return match, nil
}

func (m *MatchRepository) CreateMatch(params *types.MatchCreate) (types.Match, error) {

	rows, err := m.db.NamedQuery(`
		INSERT INTO match (hometeam,awayteam,homescore,awayscore,status)
		VALUES (:hometeam,:awayteam,:homescore,:awayscore,:status) RETURNING *
	`, map[string]interface{}{
		"hometeam":  params.HomeTeam,
		"awayteam":  params.AwayTeam,
		"homescore": params.HomeScore,
		"awayscore": params.AwayScore,
		"status":    params.Status,
	})

	if err != nil {
		return types.Match{}, nil
	}

	var match types.Match

	for rows.Next() {
		err := rows.StructScan(&match)
		if err != nil {
			return types.Match{}, err
		}
	}

	return match, nil

}

func (m *MatchRepository) UpdateMatch(params types.MatchUpdate) (types.Match, error) {
	rows, err := m.db.NamedQuery(`
		UPDATE match SET 
			homescore=:homescore, 
			awayscore=:awayscore, 
			status=:status
		WHERE id=:id
		RETURNING *
	`, map[string]interface{}{
		"id":        params.ID,
		"homescore": params.HomeScore,
		"awayscore": params.AwayScore,
		"status":    params.Status,
	})

	if err != nil {
		fmt.Println(err.Error())
		return types.Match{}, nil
	}

	var match types.Match

	for rows.Next() {
		err := rows.StructScan(&match)
		if err != nil {
			return types.Match{}, err
		}
	}

	return match, nil
}
