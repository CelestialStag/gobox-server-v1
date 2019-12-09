/*
Package routes - Contains all of the top level routes
*/
package index

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// FileController The controller for /
type FileController struct{}

// BeforeActivation
func (c *FileController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/{id:string}", "Download", downloadMiddleware)
}

// Get serves
// Method:   GET
// Resource: http://localhost:8080/f
func (c *FileController) Get(ctx iris.Context) {

	ctx.ViewData("title", "GoBox")
	ctx.View("index.pug")
}

// Download serves
// Method:   GET
// Resource: http://localhost:8080/f/{downloadid}
func (c *FileController) Download(ctx iris.Context) {

	id := ctx.Params().GetString("id")

	ctx.ViewData("id", id)
	ctx.View("file/index.pug")
}

/////////////////////////////////////////////////////
//
//		Custom Middleware
//
/////////////////////////////////////////////////////

func downloadMiddleware(ctx iris.Context) {

	ctx.Next()
}
