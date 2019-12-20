/*
Package routes - Contains all of the top level routes
*/
package index

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

	json["id"] = folder
	json["uploaded"] = "n/a"
	json["expires"] = "n/a"

	if len(list) > 0 {

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

		if strings.Contains(json["name"], "jpeg") || strings.Contains(json["name"], "jpg") || strings.Contains(json["name"], "png") || strings.Contains(json["name"], "gif") || strings.Contains(json["name"], "webp") || strings.Contains(json["name"], "apng") {
			json["img"] = "https://" + ctx.Host() + "/api/file/download/" + url
			large := "1"
			ctx.ViewData("large", large)
		} else {
			json["img"] = "https://gobox.emawa.io/public/img/gobox/logo-4.png"
		}
		json["title"] = "GoBox: " + json["name"]

		json["description"] = "Download <b>" + json["name"] + "</b>. Get more free and high speed file hosting at https://gobox.dev"

		// ctx.JSON(json)

	} else {
		json["name"] = "file does not exist"
		ctx.StatusCode(400)
	}
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
