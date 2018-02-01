package webapp

import (
	"strings"
)

type (
	Group struct {
		prefix     string
		middleware []MiddlewareFunc
		router     *router
	}
)

func newGroup(prefix string, middleware []MiddlewareFunc, router *router) *Group {
	return &Group{
		prefix:     prefix,
		middleware: append([]MiddlewareFunc{}, middleware...), // copied
		router:     router,
	}
}

func (g *Group) Configure(fn func(*Group)) {
	fn(g)
}

func (g *Group) Use(middleware ...MiddlewareFunc) {
	g.middleware = append(g.middleware, middleware...)
}

func (g *Group) GET(path string, handler HandlerFunc) {
	g.add("GET", path, handler)
}

func (g *Group) POST(path string, handler HandlerFunc) {
	g.add("POST", path, handler)
}

func (g *Group) PUT(path string, handler HandlerFunc) {
	g.add("PUT", path, handler)
}

func (g *Group) PATCH(path string, handler HandlerFunc) {
	g.add("PATCH", path, handler)
}

func (g *Group) DELETE(path string, handler HandlerFunc) {
	g.add("DELETE", path, handler)
}

func (g *Group) OPTIONS(path string, handler HandlerFunc) {
	g.add("OPTIONS", path, handler)
}

func (g *Group) HEAD(path string, handler HandlerFunc) {
	g.add("HEAD", path, handler)
}

func (g *Group) CONNECT(path string, handler HandlerFunc) {
	g.add("CONNECT", path, handler)
}

func (g *Group) TRACE(path string, handler HandlerFunc) {
	g.add("TRACE", path, handler)
}

func (g *Group) Handle(methods string, path string, handler HandlerFunc) {
	if methods == "*" {
		methods = DEFAULT_ALL_METHODS
	}
	for _, method := range strings.Split(methods, ",") {
		g.add(strings.TrimSpace(method), path, handler)
	}
}

func (g *Group) Group(path string) *Group {
	return newGroup(g.prefix+path, g.middleware, g.router)
}

func (g *Group) add(method string, path string, handler HandlerFunc) {
	// make handler chain
	for i := len(g.middleware) - 1; i >= 0; i-- {
		handler = g.middleware[i](handler)
	}
	g.router.add(method, g.prefix+path, handler)
}
