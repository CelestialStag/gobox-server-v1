package routes

import (
	"github.com/kataras/iris"
)

// ROOTController The controller for /api
type ROOTController struct {
}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *ROOTController) Get(ctx iris.Context) {

	ctx.ViewData("title", "gopy")
	ctx.View("index.pug")
}
