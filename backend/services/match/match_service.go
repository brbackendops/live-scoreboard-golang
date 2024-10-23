package match

import (
	"score/database/types"
)

func (ms *MatchService) GetMatchService(id int) (types.Match, error) {
	return ms.matchrepo.GetMatch(id)
}

func (ms *MatchService) GetMatchesService() ([]types.Match, error) {
	return ms.matchrepo.GetMatches()
}

func (ms *MatchService) CreateMatchService(params *types.MatchCreate) (types.Match, error) {
	return ms.matchrepo.CreateMatch(params)
}

func (ms *MatchService) UpdateMatchService(params *types.MatchUpdate) (types.Match, error) {
	return ms.matchrepo.UpdateMatch(*params)
}
