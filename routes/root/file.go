package root

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// FILEController The controller for /api
type FILEController struct {
}

func (c *FILEController) BeforeActivation(b mvc.BeforeActivation) {
	downloadMiddleware := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /custom_path")
		ctx.ViewData("id", ctx.Params().Get("id"))
		ctx.View("download.pug")
	}
	b.Handle("GET", "/{id:string}", "Download", downloadMiddleware)
}

// CustomHandlerWithoutFollowingTheNamingGuide serves
// Method:   GET
// Resource: http://localhost:8080/custom_path
func (c *FILEController) Download(id string) string {
	return id
}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *FILEController) Get(ctx iris.Context) {

	ctx.ViewData("title", "gopy")
	ctx.View("index.pug")
}
