package middlewares

import "github.com/valyala/fasthttp"

func CorsMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		// ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		next(ctx)
	}
}
