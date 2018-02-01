package middleware

import (
	"fmt"
	"time"

	"github.com/subchen/go-stack/net/webapp"
)

func Logger() webapp.MiddlewareFunc {
	return func(next webapp.HandlerFunc) webapp.HandlerFunc {
		return func(ctx *webapp.Context) {
			fmt.Printf("%s %s\n", ctx.Method(), ctx.Request.URL.String())

			start := time.Now()
			next(ctx)

			fmt.Printf("time: %d\n", time.Since(start))
		}
	}
}
