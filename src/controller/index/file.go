/*
Package routes - Contains all of the top level routes
*/
package index

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/h2non/filetype"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// FileController The controller for /
type FileController struct{}

// BeforeActivation
func (c *FileController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/{id:string}", "Download", downloadMiddleware)
	// b.Handle("GET", "/{id:string}", "Get", downloadMiddleware)
	// b.Handle("GET", "/", "Info", listMiddleware)
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

	json := make(map[string]string) // return json

	folder := ctx.Params().GetString("id")
	dir := "./data/" + folder

	f, err := os.Open(dir)
	if err != nil {
	}

	list, err := f.Readdir(-1)
	f.Close()

	file := dir + "/" + list[0].Name()
	url := folder + "/" + list[0].Name()

	buf, _ := ioutil.ReadFile(file)
	kind, unknown := filetype.Match(buf)
	if unknown != nil {
		fmt.Printf("Unknown: %s", unknown)
	}

	json["name"] = list[0].Name()
	json["size"] = fmt.Sprintf("%.0f", float64(list[0].Size())) // fmt.Sprintf("%.2f", float64(list[0].Size())/1e+6)
	json["type"] = kind.MIME.Value
	json["url"] = "/file/" + url
	json["id"] = folder
	json["uploaded"] = "n/a"
	json["expires"] = "n/a"

	json["img"] = "https://" + ctx.Host() + "/api/file/download/" + url
	json["title"] = "GoBox: " + json["name"]
	json["description"] = "Download file " + json["name"] + ". Get more free and private and high speed file hosting at gobox.dev"

	// ctx.JSON(json)

	id := ctx.Params().GetString("id")

	ctx.ViewData("id", id)
	ctx.ViewData("data", json)
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
