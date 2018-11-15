package routes

import (
	"github.com/kataras/iris"
)

// APIController The controller for /a
type APIController struct{}

/*
Get serves
Method:   GET
Resource: http://localhost:8080
*/
func (c *APIController) Get(ctx iris.Context) {
	ctx.View("index.pug")
}
