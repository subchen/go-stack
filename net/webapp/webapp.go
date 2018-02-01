package webapp

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/subchen/go-log"
)

type (
	HandlerFunc func(ctx *Context)

	MiddlewareFunc func(next HandlerFunc) HandlerFunc

	Webapp struct {
		prefix        string
		middlewarePre []MiddlewareFunc
		middleware    []MiddlewareFunc
		router        *router
		handler       HandlerFunc

		// Enables automatic redirection if the current route can't be matched but a
		// handler for the path with (without) the trailing slash exists.
		// For example if /foo/ is requested but a route only exists for /foo, the
		// client is redirected to /foo with http status code 301 for GET requests
		// and 307 for all other request methods.
		RedirectTrailingSlash bool
	}
)

const (
	DEFAULT_ALL_METHODS = "GET,POST,PUT,PATCH,DELETE"
)

func New(prefix string) *Webapp {
	if prefix == "" || prefix == "/" {
		prefix = ""
	} else {
		if !strings.HasPrefix(prefix, "/") {
			panic("prefix should be start with '/'")
		}
		if strings.HasSuffix(prefix, "/") {
			panic("prefix should NOT be end with '/'")
		}
	}

	app := &Webapp{}
	app.prefix = prefix
	app.router = newRouter()
	app.RedirectTrailingSlash = true
	return app
}

func (app *Webapp) UsePre(middleware ...MiddlewareFunc) {
	app.middlewarePre = append(app.middlewarePre, middleware...)
}

func (app *Webapp) Use(middleware ...MiddlewareFunc) {
	app.middleware = append(app.middleware, middleware...)
}

func (app *Webapp) GET(path string, handler HandlerFunc) {
	app.add("GET", path, handler)
}

func (app *Webapp) POST(path string, handler HandlerFunc) {
	app.add("POST", path, handler)
}

func (app *Webapp) PUT(path string, handler HandlerFunc) {
	app.add("PUT", path, handler)
}

func (app *Webapp) PATCH(path string, handler HandlerFunc) {
	app.add("PATCH", path, handler)
}

func (app *Webapp) DELETE(path string, handler HandlerFunc) {
	app.add("DELETE", path, handler)
}

func (app *Webapp) OPTIONS(path string, handler HandlerFunc) {
	app.add("OPTIONS", path, handler)
}

func (app *Webapp) HEAD(path string, handler HandlerFunc) {
	app.add("HEAD", path, handler)
}

func (app *Webapp) CONNECT(path string, handler HandlerFunc) {
	app.add("CONNECT", path, handler)
}

func (app *Webapp) TRACE(path string, handler HandlerFunc) {
	app.add("TRACE", path, handler)
}

func (app *Webapp) Handle(methods string, path string, handler HandlerFunc) {
	if methods == "*" {
		methods = DEFAULT_ALL_METHODS
	}
	for _, method := range strings.Split(methods, ",") {
		app.add(strings.TrimSpace(method), path, handler)
	}
}

func (app *Webapp) Group(path string) *Group {
	return newGroup(app.prefix+path, app.middleware, app.router)
}

func (app *Webapp) add(method string, path string, handler HandlerFunc) {
	// make handler chain
	for i := len(app.middleware) - 1; i >= 0; i-- {
		handler = app.middleware[i](handler)
	}
	app.router.add(method, app.prefix+path, handler)
}

// Routes returns all register route path
func (app *Webapp) Routes() []string {
	var paths []string
	for _, routes := range app.router.routesList {
		paths = append(paths, routes.path)
	}
	return paths
}

// ServeHTTP implements http.Handler
func (app *Webapp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(w, r, app)

	// handle 500 panic
	defer func() {
		if err := recover(); err != nil {
			ctx.Error(http.StatusInternalServerError, fmt.Errorf("panic: %v", err))
		}
	}()

	// it make middleware and handler as a single chain handler
	// and cached in app.handler
	// it will be built on first request
	if app.handler == nil {
		// make chain for pre middleware
		handler := app.serveHttpRequestHandler()
		for i := len(app.middlewarePre) - 1; i >= 0; i-- {
			handler = app.middlewarePre[i](handler)
		}
		app.handler = handler
	}

	// execute middleware and handler
	app.handler(ctx)
}

func (app *Webapp) serveHttpRequestHandler() HandlerFunc {
	return func(ctx *Context) {
		pathnames := strings.Split(ctx.Path(), "/")

		// 1. get routes by path
		routes := app.router.find(pathnames)
		if routes == nil {
			// try to fix url and redirect
			if app.RedirectTrailingSlash && ctx.Path() != "/" {
				last := len(pathnames) - 1
				if pathnames[last] == "" {
					pathnames = pathnames[0:last]
				} else {
					pathnames = append(pathnames, "")
				}
				routes = app.router.find(pathnames)
				if routes != nil {
					// redirect with query string (trim slash redirect)
					ctx.Request.URL.Path = strings.Join(pathnames, "/")
					ctx.Redirect(ctx.Request.URL.String())
					return
				}
			}

			ctx.Error(http.StatusNotFound, fmt.Errorf("Request not found: %s", ctx.Path()))
			return
		}

		// 2. get route handler
		route := routes.find(ctx.Method())
		if route == nil {
			ctx.ResponseWriter.Header().Set("Allow", routes.allows())
			if ctx.Method() == "OPTIONS" {
				ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
			} else {
				ctx.ResponseWriter.WriteHeader(http.StatusMethodNotAllowed)
			}
			return
		}
		log.Debug("Found route: %v", route.path)

		// 3. extra vars param from path
		ctx.vars = app.router.makeVars(route.path, pathnames)

		// 4. execute middleware and handler
		route.handler(ctx)
	}
}

func (app *Webapp) Run(addr string) error {
	return http.ListenAndServe(addr, app)
}

func (app *Webapp) RunTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, app)
}
