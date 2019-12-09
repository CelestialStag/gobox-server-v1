package api

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/h2non/filetype"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// FILEController The controller for /a/f
type FileController struct{}

func (c *FileController) BeforeActivation(b mvc.BeforeActivation) {
	downloadMiddleware := func(ctx iris.Context) {

		id := ctx.Params().GetString("id")
		file := ctx.Params().GetString("file")

		dir := "./data/" + id + "/" + file
		ctx.SendFile(dir, file)
	}

	listMiddleware := func(ctx iris.Context) {

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
		json["url"] = url
		json["id"] = folder
		json["uploaded"] = "n/a"
		json["expires"] = "n/a"

		ctx.JSON(json)
	}

	b.Handle("GET", "/download/{id:string}/{file:string}", "Download", downloadMiddleware)
	b.Handle("GET", "/info/{id:string}", "Info", listMiddleware)
}

// Download the controller for /f/download
func (c *FileController) Download(id string) {}

// Info The controller for /f/info
func (c *FileController) Info(ctx iris.Context) {}

// PostUpload The controller for /f/upload
func (c *FileController) PostUpload(ctx iris.Context) {

	// Generate random hash for dir
	g, nil := GenerateRandomBytes(4)
	token := fmt.Sprintf("%x", g)
	dir := "./data/" + token

	// Get the file from the request.
	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON("{ error:" + err.Error() + "}")
		return
	}
	defer file.Close()
	fname := info.Filename

	// Create the directory
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON("{ error:" + err.Error() + "}")
			return
		}
	}

	// Create a file with the same name
	out, err := os.OpenFile(dir+"/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON("{ error:" + err.Error() + "}")
		return
	}
	defer out.Close()

	io.Copy(out, file)

	ctx.JSON(map[string]string{"url": dir, "file": fname, "hash": token}) // or myjsonStruct{hello:"json}
}

/*
GenerateRandomBytes - generates random bytes into a bye array
*/
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
