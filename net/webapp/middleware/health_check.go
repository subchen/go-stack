package middleware

import (
	"fmt"
	"strings"

	"github.com/subchen/go-stack/net/webapp"
)

func HealthCheckHandler(ctx *webapp.Context) {
	accept := ctx.GetHeader("Accept")

	if strings.Contains(accept, "application/json") {
		ctx.SetHeader("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(ctx.ResponseWriter, "{\"health\": true}\n")
		return
	}

	if strings.Contains(accept, "application/xml") {
		ctx.SetHeader("Content-Type", "application/xml; charset=utf-8")
		fmt.Fprintf(ctx.ResponseWriter, "<health>true</health>\n")
		return
	}

	ctx.SetHeader("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(ctx.ResponseWriter, "health OK\n")
}
