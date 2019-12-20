package main

import (
	"github.com/CeruleanSong/gopy-server/src/controller"
	"github.com/CeruleanSong/gopy-server/src/controller/api"
	"github.com/CeruleanSong/gopy-server/src/controller/index"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func initialize() *iris.Application {

	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	/* Middleware */
	app.Use(recover.New())
	app.Use(logger.New())

	// Static files
	app.HandleDir("/public", "res/public", iris.DirOptions{
		Gzip: false,
	})

	/* View Engine */
	app.RegisterView(iris.Pug("res/view", ".pug").Reload(true))

	/* Routes */
	// ROOT
	mvc.Configure(app.Party("/"), func(app *mvc.Application) {

		app.Party("/file").Handle(new(index.FileController))
	}).Handle(new(controller.IndexController))

	// API
	mvc.Configure(app.Party("/api"), func(app *mvc.Application) {

		app.Party("/file").Handle(new(api.FileController))
	}).Handle(new(controller.ApiController))

	return app
}

/*
main - Gets the iris object and runs it
*/
func main() {
	app := initialize()

	app.Run(iris.Addr(":7070"),
		iris.WithPostMaxMemory(1.2e+9),
		iris.WithoutServerError(iris.ErrServerClosed))
}
