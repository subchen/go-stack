package webapp

import (
	"regexp"
	"strings"
)

type (
	router struct {
		root *node

		// all path routes
		routesList []*routes

		// a redefined mapping for pathVars: path -> (param -> index)
		// Example:
		//   /users/{userid}/books/{bookid}" -> { userid: 1, bookid: 3}
		//   /files/{file*}"                 -> { file: 1}
		mappingParamIndex map[string]map[string]int
	}

	node struct {
		name     string // may be "{}", "*", "", ...
		children map[string]*node
		routes   *routes
	}

	// routes is group of route, they have same path pattern
	routes struct {
		path   string // origin path pattern
		routes []*route
	}

	// route is a defined handler for path+method
	route struct {
		path    string // origin path pattern
		method  string
		handler HandlerFunc
	}
)

var (
	// validate path pattern, support {name}, {name*} as param match
	RE_PATH_PATTERN_1 = regexp.MustCompile(`^(/([a-zA-Z0-9_\-.]+|\{[A-zA-Z0-9_\-]+\*?\}))*/?$`)
	// validate {name*} must be last
	RE_PATH_PATTERN_2 = regexp.MustCompile(`^[^*]+(\*\})?$`)
)

func newRouter() *router {
	return &router{
		root: &node{
			name: "/",
		},
		routesList:        nil,
		mappingParamIndex: make(map[string]map[string]int, 16),
	}
}

// add registers a handler into router tree
func (r *router) add(method string, path string, handler HandlerFunc) {
	// 0. validate path pattern
	if !RE_PATH_PATTERN_1.MatchString(path) {
		panic("Invalid path pattern: " + path)
	}
	if !RE_PATH_PATTERN_2.MatchString(path) {
		panic("Invalid path pattern, any match must be put on last position: " + path)
	}

	names := strings.Split(path, "/")

	// 1. find node
	n := r.root
	for i := 1; i < len(names); i++ {
		name := names[i]
		if strings.Contains(name, "*") {
			name = "*" // any
		} else if strings.Contains(name, "{") {
			name = "{}" // param
		}

		nn, ok := n.children[name]
		if !ok {
			nn = &node{
				name:     name,
				children: nil,
				routes:   nil,
			}

			if n.children == nil {
				n.children = make(map[string]*node, 4)
			}
			n.children[name] = nn
		}

		n = nn
	}

	// 2. add new route to node
	route := &route{
		path:    path,
		method:  method,
		handler: handler,
	}

	if n.routes == nil {
		n.routes = &routes{
			path:   path,
			routes: nil,
		}
		r.routesList = append(r.routesList, n.routes)
	}

	n.routes.routes = append(n.routes.routes, route)

	// 3. make param mapping for path
	if _, ok := r.mappingParamIndex[path]; !ok {
		mapping := make(map[string]int, 4)
		for i, name := range names {
			if len(name) > 0 && name[0] == '{' {
				name = name[1 : len(name)-1]
				mapping[name] = i
			}
		}
		r.mappingParamIndex[path] = mapping
	}
}

// find routes for specified path, pathnames is a split for path
func (r *router) find(pathnames []string) *routes {
	if n := r.root.find(pathnames[1], pathnames[2:]); n != nil {
		return n.routes
	}
	return nil
}

// makeVars returns vars
func (r *router) makeVars(pathpattern string, pathnames []string) map[string]string {
	if mapping, ok := r.mappingParamIndex[pathpattern]; ok {
		vars := make(map[string]string, len(mapping))
		for name, index := range mapping {
			if name[len(name)-1] == '*' {
				name = name[0 : len(name)-1]
				vars[name] = strings.Join(pathnames[index:], "/")
			} else {
				vars[name] = pathnames[index]
			}
		}
		return vars
	}

	return nil
}

func (n *node) find(name string, pathnames []string) *node {
	if len(n.children) == 0 {
		return nil
	}

	// static
	if child, ok := n.children[name]; ok {
		if len(pathnames) == 0 {
			if child.routes != nil {
				return child // match
			}
		} else {
			nn := child.find(pathnames[0], pathnames[1:])
			if nn != nil {
				return nn
			}
		}
	}

	if name == "" {
		return nil
	}

	// param
	if child, ok := n.children["{}"]; ok {
		if len(pathnames) == 0 {
			if child.routes != nil {
				return child // match
			}
		} else {
			nn := child.find(pathnames[0], pathnames[1:])
			if nn != nil {
				return nn
			}
		}
	}

	// any *
	if child, ok := n.children["*"]; ok {
		return child
	}

	return nil
}

// find returns the matched route
func (r *routes) find(method string) *route {
	for _, route := range r.routes {
		if route.method == method {
			return route
		}
	}
	return nil
}

// methods returns all available methods (OPTIONS)
func (r *routes) allows() string {
	var methods []string
	for _, route := range r.routes {
		methods = append(methods, route.method)
	}
	allows := strings.Join(methods, ", ")

	if !strings.Contains(allows, "OPTIONS") {
		allows = allows + ", OPTIONS"
	}
	return allows
}
