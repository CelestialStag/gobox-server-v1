package root

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// FILEController The controller for /f
type FILEController struct {
}

func (c *FILEController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/{id:string}", "Download", downloadMiddleware)
}

// Get serves
// Method:   GET
// Resource: http://localhost:8080/f
func (c *FILEController) Get(ctx iris.Context) {

	ctx.ViewData("title", "gopy")
	ctx.View("index.pug")
}

// Download serves
// Method:   GET
// Resource: http://localhost:8080/f/{downloadid}
func (c *FILEController) Download(ctx iris.Context) {

	id := ctx.Params().GetString("id")

	ctx.ViewData("id", id)
	ctx.View("download.pug")
}

/////////////////////////////////////////////////////
//
//		Custom Middleware
//
/////////////////////////////////////////////////////

func downloadMiddleware(ctx iris.Context) {

	ctx.Next()
}
