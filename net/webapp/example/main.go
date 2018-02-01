package main

import (
	"fmt"

	"github.com/subchen/go-log"
	"github.com/subchen/go-stack/net/webapp"
	"github.com/subchen/go-stack/net/webapp/middleware"
)

func userGroupHandles(g *webapp.Group) {
	handler := func(ctx *webapp.Context) {
		fmt.Fprintf(ctx.ResponseWriter, "userid = %s\n", ctx.Vars("id"))
	}

	g.Handle("GET", "", handler)
	g.Handle("GET", "/{id}", handler)
	g.Handle("POST", "/{id}", handler)
	g.Handle("PUT", "/{id}", handler)
	g.Handle("PATCH", "/{id}", handler)
	g.Handle("DELETE", "/{id}", handler)
	g.Handle("GET", "/{id}/profiles", handler)
}

func main() {

	app := webapp.New("/v2")

	app.UsePre(middleware.Logger())

	app.Use(middleware.CORS())

	app.Use(WrapMiddleware(func(ctx *webapp.Context) {
		log.Info("Middleware")
	}))
	/*
		app.GET("/health", handler)
		app.POST("/stats", handler)

		app.Handle("*", "/stats", handler)
		app.Handle("POST,PUT", "/stats", handler)

		g := webapp.Group("/admin")
		g.Use(...)
		g.GET(...)

		app.Group("/users").Apply(userGroupHandles)
	*/

	app.GET("/health", middleware.HealthCheckHandler)
	app.POST("/health", middleware.HealthCheckHandler)

	app.GET("/ping", middleware.HealthCheckHandler)
	app.GET("/ping/{id}/{name}", func(ctx *webapp.Context) {
		fmt.Fprintf(ctx.ResponseWriter, "id = %s\n", ctx.Vars("id"))
		fmt.Fprintf(ctx.ResponseWriter, "name = %s\n", ctx.Vars("name"))
	})

	app.Group("/users").Configure(userGroupHandles)

	fmt.Println(app.Routes())

	fmt.Println("Listening http://127.0.0.1:8080/v2/")
	app.Run(":8080")
}

func WrapMiddleware(handler webapp.HandlerFunc) webapp.MiddlewareFunc {
	return func(next webapp.HandlerFunc) webapp.HandlerFunc {
		return func(ctx *webapp.Context) {
			handler(ctx)
			next(ctx)
		}
	}
}
