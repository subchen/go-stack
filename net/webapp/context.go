package webapp

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	vars           map[string]string
	app            *Webapp
}

func newContext(w http.ResponseWriter, r *http.Request, app *Webapp) *Context {
	return &Context{
		ResponseWriter: w,
		Request:        r,
		vars:           nil,
		app:            app,
	}
}

func (ctx *Context) URL() *url.URL {
	return ctx.Request.URL
}

func (ctx *Context) Path() string {
	return ctx.Request.URL.Path
}

func (ctx *Context) QueryString() string {
	return ctx.Request.URL.RawQuery
}

func (ctx *Context) Method() string {
	return ctx.Request.Method
}

func (ctx *Context) GetHeader(name string) string {
	return ctx.Request.Header.Get(name)
}

func (ctx *Context) SetHeader(name string, value string) {
	ctx.ResponseWriter.Header().Set(name, value)
}

func (ctx *Context) Vars(name string) string {
	return ctx.vars[name]
}

func (ctx *Context) Form(name string) string {
	return ctx.Request.Form.Get(name)
}

func (ctx *Context) FormValues(name string) []string {
	return ctx.Request.Form[name]
}

func (ctx *Context) FormFile(name string) (*multipart.FileHeader, error) {
	_, f, err := ctx.Request.FormFile(name)
	return f, err
}

func (ctx *Context) MultipartForm() (*multipart.Form, error) {
	err := ctx.Request.ParseMultipartForm(32 << 20) // buffer 32M
	return ctx.Request.MultipartForm, err
}

func (ctx *Context) Cookie(name string) (*http.Cookie, error) {
	return ctx.Request.Cookie(name)
}

func (ctx *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(ctx.ResponseWriter, cookie)
}

func (ctx *Context) Cookies() []*http.Cookie {
	return ctx.Request.Cookies()
}

func (ctx *Context) JSON(code int, obj interface{}) (err error) {
	query := ctx.Request.URL.Query()

	// pretty json
	var b []byte
	if _, ok := query["pretty"]; ok {
		b, err = json.MarshalIndent(obj, "", "  ")
	} else {
		b, err = json.Marshal(obj)
	}
	if err != nil {
		return
	}

	// jsonp
	jsonp := query.Get("callback")

	resp := ctx.ResponseWriter
	if jsonp == "" {
		resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	} else {
		resp.Header().Set("Content-Type", "application/javascript; charset=UTF-8")
	}

	resp.WriteHeader(code)

	if jsonp != "" {
		if _, err = resp.Write([]byte(jsonp + "(")); err != nil {
			return
		}
	}

	if _, err = resp.Write(b); err != nil {
		return
	}

	if jsonp != "" {
		if _, err = resp.Write([]byte(");")); err != nil {
			return
		}
	}

	return
}

func (ctx *Context) Redirect(url string) {
	code := http.StatusMovedPermanently // 301
	if ctx.Request.Method != "GET" {
		code = http.StatusTemporaryRedirect // 307
	}
	http.Redirect(ctx.ResponseWriter, ctx.Request, url, code)
}

func (ctx *Context) Error(code int, err error) {
	w := ctx.ResponseWriter

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintf(w, "%+v\n", err)
}
