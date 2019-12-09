/*
Package routes - Contains all of the top level routes
*/
package controller

import (
	"github.com/kataras/iris/v12"
)

// ROOTController The controller for /
type IndexController struct{}

func (c *IndexController) Get(ctx iris.Context) {
	ctx.ViewData("title", "GoBox")
	ctx.View("index.pug")
}
