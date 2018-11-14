package api

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// FILEController The controller for /api
type FILEController struct{}

func (c *FILEController) BeforeActivation(b mvc.BeforeActivation) {
	downloadMiddleware := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /REadME")

		id := ctx.Params().GetString("id")
		file := ctx.Params().GetString("file")

		dir := "./data/" + id + "/" + file
		ctx.SendFile(dir, file)
	}

	listMiddleware := func(ctx iris.Context) {

		ls := make(map[int]string)
		var ls2 = make([]string, 0)

		folder := ctx.Params().GetString("id")
		dir := "./data/" + folder

		f, err := os.Open(dir)
		if err != nil {
		}

		list, err := f.Readdir(-1)
		f.Close()

		sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })

		for i, file := range list {
			ls2 = append(ls2, file.Name())
			ls[i] = file.Name()
		}

		ctx.JSON(ls2) // or myjsonStruct{hello:"json}
	}

	b.Handle("GET", "/download/{id:string}/{file:string}", "Download", downloadMiddleware)
	b.Handle("GET", "/list/{id:string}", "List", listMiddleware)
}

// CustomHandlerWithoutFollowingTheNamingGuide serves
// Method:   GET
// Resource: http://localhost:8080/custom_path
func (c *FILEController) Download(id string) {}

// PostUpload The controller for /api
func (c *FILEController) List(ctx iris.Context) {}

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

type UploadResponce struct {
	File string `json:"file"`
	URL  string `json:"url"`
}

// PostUpload The controller for /api
func (c *FILEController) PostUpload(ctx iris.Context) {

	g, nil := GenerateRandomBytes(4)
	token := fmt.Sprintf("%x", g)
	dir := "./data/" + token

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}

	ctx.UploadFormFiles(dir)
	ctx.JSON(map[string]string{"url": dir, "hash": token}) // or myjsonStruct{hello:"json}
}

func saveUploadedFile(fh *multipart.FileHeader, destDirectory string) (int64, error) {
	src, err := fh.Open()
	if err != nil {
		return 0, err
	}
	defer src.Close()

	out, err := os.OpenFile(filepath.Join(destDirectory, fh.Filename),
		os.O_WRONLY|os.O_CREATE, os.FileMode(0666))

	if err != nil {
		return 0, err
	}
	defer out.Close()

	return io.Copy(out, src)
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
