package api

import "github.com/kataras/iris/mvc"

// AUTHController The controller for /api
type AUTHController struct{}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *AUTHController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome to /v1/file</h1>",
	}
}
