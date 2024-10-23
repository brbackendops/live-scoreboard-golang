package routes

import (
	matchCont "score/controller/match"
	sdb "score/database"
	redisclient "score/redisClient"
	matchRepo "score/repo/match"
	matchServ "score/services/match"

	"github.com/fasthttp/router"
)

func MatchRoutes(r *router.Router) {
	m := r.Group("/match")

	db := sdb.Connect()
	mrepo := matchRepo.NewMatchRepo(db)
	mserv := matchServ.MakeMatchServ(mrepo)

	rclient := redisclient.RedisClient()
	mcont := matchCont.NewMatchCont(mserv, rclient)

	m.GET("/", mcont.GetMatchesController)
	m.GET("/{id}", mcont.GetMatchController)
	m.POST("/create", mcont.CreateMatchController)
	m.PUT("/update", mcont.UpdateMatchController)
}
