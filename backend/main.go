package main

import (
	"log"
	"score/socket"

	"score/middlewares"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	mRoutes "score/routes"
)

func main() {

	mainRouter := router.New()

	mainRouter.GET("/", func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("welcome")
	})

	mRoutes.MatchRoutes(mainRouter)

	go socket.StartWebSocket()
	if err := fasthttp.ListenAndServe(":4000", middlewares.CorsMiddleware(mainRouter.Handler)); err != nil {
		log.Fatal(err.Error())
	}
}
