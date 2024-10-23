package match

import (
	matchRepo "score/repo/match"
)

type MatchService struct {
	matchrepo *matchRepo.MatchRepository
}

func MakeMatchServ(mrepo *matchRepo.MatchRepository) *MatchService {
	return &MatchService{
		matchrepo: mrepo,
	}
}
