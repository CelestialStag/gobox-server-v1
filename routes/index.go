package routes

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

// INDEXController The controller for /api
type INDEXController struct{}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *INDEXController) Get(ctx iris.Context) {

	ctx.ViewData("pageTitle", "My Index Page")
	ctx.View("index.pug")
}

func (c *INDEXController) GetUpload(ctx iris.Context) {
	now := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(now, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))

	ctx.ViewData("pageTitle", "My Upload Page")
	ctx.ViewData("token", token)
	ctx.View("upload.pug")
}
