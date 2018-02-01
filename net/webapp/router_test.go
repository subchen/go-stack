package webapp

import (
	"github.com/subchen/gstack/assert"
	"strings"
	"testing"
)

func TestRouter(t *testing.T) {
	// 1. prepare router
	r := newRouter()
	r.add("GET", "/", nil)
	r.add("GET", "/admin/", nil)
	r.add("GET", "/admin/status", nil)
	r.add("POST", "/admin/maintenance", nil)
	r.add("GET", "/users", nil)
	r.add("GET", "/users/{id}", nil)
	r.add("POST", "/users/{id}", nil)
	r.add("PUT", "/users/{id}", nil)
	r.add("PATCH", "/users/{id}", nil)
	r.add("DELETE", "/users/{id}", nil)
	r.add("GET", "/users/{id}/books", nil)
	r.add("GET", "/users/{id}/books/{book-id}", nil)
	r.add("POST", "/users/{id}/books/{book-id}", nil)
	r.add("PUT", "/users/{id}/books/{book-id}", nil)
	r.add("DELETE", "/users/{id}/books/{book-id}", nil)
	r.add("GET", "/users/{id}/files", nil)
	r.add("GET", "/users/{id}/files/{file*}", nil)
	r.add("GET", "/users/count", nil)
	r.add("GET", "/abc", nil)
	r.add("GET", "/abcd", nil)
	r.add("GET", "/abcde", nil)
	r.add("GET", "/abcdef", nil)
	r.add("GET", "/abcdefg", nil)
	r.add("GET", "/{id}/test/", nil)
	r.add("GET", "/{id}/test/books", nil)

	var routes = func(path string) *routes {
		return r.find(strings.Split(path, "/"))
	}

	var route = func(path string, method string) *route {
		routes := routes(path)
		if routes != nil {
			return routes.find(method)
		}
		return nil
	}

	// 2. testing
	assert := assert.New(t)

	assert.NotNil(route("/", "GET"))
	assert.NotNil(route("/admin/", "GET"))
	assert.NotNil(route("/admin/status", "GET"))
	assert.NotNil(route("/admin/maintenance", "POST"))
	assert.NotNil(route("/users", "GET"))
	assert.NotNil(route("/users/123", "GET"))
	assert.NotNil(route("/users/123", "DELETE"))
	assert.NotNil(route("/users/123/books", "GET"))
	assert.NotNil(route("/users/123/books/456", "GET"))
	assert.NotNil(route("/users/123/books/456", "POST"))
	assert.NotNil(route("/users/123/files", "GET"))
	assert.NotNil(route("/users/123/files/a-b-c.zip", "GET"))
	assert.NotNil(route("/users/123/files/1/2/3/a-b-c.zip", "GET"))
	assert.NotNil(route("/users/count", "GET"))
	assert.NotNil(route("/abc", "GET"))
	assert.NotNil(route("/users/test", "GET"))
	assert.NotNil(route("/users/test/books/123", "GET"))
	assert.NotNil(route("/users/test/", "GET"))
	assert.NotNil(route("/users/test/books", "GET"))

	assert.Nil(route("/admin", "GET"))
	assert.Nil(route("/admin/status/", "GET"))
	assert.Nil(route("/admin/maintenance", "GET"))
	assert.Nil(route("/users/", "GET"))
	assert.Nil(route("/users/123/", "GET"))
	assert.Nil(route("/users/123/books/", "GET"))
	assert.Nil(route("/users/123/books/456/", "GET"))
	assert.Nil(route("/users/123/files/", "GET"))
	assert.Nil(route("/users/count", "POST"))
	assert.Nil(route("/abc/def", "GET"))
	assert.Nil(route("/users/test/books/123/", "GET"))

	assert.Equal(routes("/").allows(), "GET, OPTIONS")
	assert.Equal(routes("/admin/maintenance").allows(), "POST, OPTIONS")
	assert.Equal(routes("/users/123").allows(), "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	assert.Equal(routes("/users/123/books/456").allows(), "GET, POST, PUT, DELETE, OPTIONS")

	assert.Equal(route("/users/123", "GET").path, "/users/{id}")
	assert.Equal(route("/users/count", "GET").path, "/users/count")
	assert.Equal(route("/users/test", "GET").path, "/users/{id}")
	assert.Equal(route("/users/test/books", "GET").path, "/users/{id}/books")
	assert.Equal(route("/users/test/", "GET").path, "/{id}/test/")
}

func TestMakeVars(t *testing.T) {
	// 1. prepare router
	r := newRouter()
	r.add("GET", "/users/{id}", nil)
	r.add("GET", "/users/{id}/books/{book-id}", nil)
	r.add("GET", "/users/{id}/files/{file*}", nil)

	var vars = func(path string, method string) map[string]string {
		pathnames := strings.Split(path, "/")
		routes := r.find(pathnames)
		route := routes.find(method)
		return r.makeVars(route.path, pathnames)
	}

	// 2. testing
	assert := assert.New(t)
	assert.Equal(vars("/users/123", "GET"), map[string]string{"id": "123"})
	assert.Equal(vars("/users/123/books/456", "GET"), map[string]string{"id": "123", "book-id": "456"})
	assert.Equal(vars("/users/123/files/1/2/3/a-b-c.zip", "GET"), map[string]string{"id": "123", "file": "1/2/3/a-b-c.zip"})
}
