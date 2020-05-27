package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// Context defines the context of gee
type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm returns the form of post
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query key
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status write code to Writer
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader sets header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String format  string
func (c *Context) String(code int, format string, values ...interface{}) {
	// 请求的MIME的类型（用于POST PUT）
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	// Sprintf returns the result string
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON  format json
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data format data
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	// write the data to the connection
	c.Writer.Write(data)
}

// HTML format html
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
