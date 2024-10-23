package match

import (
	matchS "score/services/match"

	"github.com/go-redis/redis/v8"
)

type MatchController struct {
	matchService *matchS.MatchService
	rClient      *redis.Client
}

func NewMatchCont(mserv *matchS.MatchService, r *redis.Client) *MatchController {
	return &MatchController{
		matchService: mserv,
		rClient:      r,
	}
}
