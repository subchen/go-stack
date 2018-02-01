package middleware

import (
	"github.com/subchen/go-stack/net/webapp"
)

const (
	HeaderAccessControlRequestMethod    = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders   = "Access-Control-Request-Headers"
	HeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	HeaderAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge           = "Access-Control-Max-Age"
)

func CORS() webapp.MiddlewareFunc {
	return func(next webapp.HandlerFunc) webapp.HandlerFunc {
		return func(ctx *webapp.Context) {
			ctx.SetHeader(HeaderAccessControlAllowOrigin, "*")
			ctx.SetHeader(HeaderAccessControlAllowMethods, webapp.DEFAULT_ALL_METHODS)
			ctx.SetHeader(HeaderAccessControlAllowCredentials, "true")

			ctx.SetHeader(HeaderAccessControlAllowHeaders, ctx.GetHeader(HeaderAccessControlRequestHeaders))

			//ctx.SetHeader(HeaderAccessControlMaxAge, "0")

			next(ctx)
		}
	}
}
