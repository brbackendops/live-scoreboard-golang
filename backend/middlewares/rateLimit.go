package middlewares

import (
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(rate.Every(time.Second), 100)

func RateLimit(next fasthttp.RequestHandler) fasthttp.RequestHandler {

	return func(ctx *fasthttp.RequestCtx) {
		if !limiter.Allow() {
			ctx.Error("you have ", http.StatusBadRequest)
			return
		}
		next(ctx)
	}

}
