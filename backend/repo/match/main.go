package match

import (
	"score/database/types"
)

type MatchRepo interface {
	GetMatches() ([]types.Match, error)
	GetMatch(int) (types.Match, error)
	CreateMatch(types.MatchCreate) (types.Match, error)
	UpdateMatch(types.MatchUpdate) (types.Match, error)
}
