package api

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

// FILEController The controller for /api
type FILEController struct{}

// GetUpload The controller for /api
func (c *FILEController) GetUpload(ctx iris.Context) {
	// create a token (optionally).

	now := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(now, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))

	// render the form with the token for any use you'd like.
	// ctx.ViewData("", token)
	// or add second argument to the `View` method.
	// Token will be passed as {{.}} in the template.
	ctx.ViewData("pageTitle", "My Upload Page")
	ctx.ViewData("token", token)
	ctx.View("upload.pug")
}

// func (c *FILEController) PostUpload(ctx iris.Context) {
// 	//
// 	// UploadFormFiles
// 	// uploads any number of incoming files ("multiple" property on the form input).
// 	//

// 	// second argument is totally optionally,
// 	// it can be used to change a file's name based on the request,
// 	// at this example we will showcase how to use it
// 	// by prefixing the uploaded file with the current user's ip.
// 	ctx.UploadFormFiles("./uploads", beforeSave)
// }

// PostUpload The controller for /api
func (c *FILEController) PostUpload(ctx iris.Context) {
	fmt.Print(ctx)
	// Get the file from the request.
	file, info, err := ctx.FormFile("uploadfile")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}

	defer file.Close()
	fname := info.Filename

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile("./data/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer out.Close()

	io.Copy(out, file)
}
