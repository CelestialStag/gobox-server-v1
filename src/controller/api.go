package controller

import (
	"github.com/kataras/iris/v12"
)

// APIController The controller for /a
type ApiController struct{}

/*
Get serves
Method:   GET
Resource: http://localhost:8080
*/
func (c *ApiController) Get(ctx iris.Context) {
	ctx.View("index.pug")
}
