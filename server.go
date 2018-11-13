package main

import (
	"gopy/routes"
	"gopy/routes/api"
	"gopy/routes/root"

	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"github.com/kataras/iris/mvc"
)

func initialize() *iris.Application {

	app := iris.New()

	//////////////////////////////////////////////////////////////////
	//
	//		setup
	//
	//////////////////////////////////////////////////////////////////

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// View Engine
	app.RegisterView(iris.Pug("./views", ".pug").Reload(true))

	// Static files
	app.StaticWeb("/static", "./static")

	//////////////////////////////////////////////////////////////////
	//
	//		routes
	//
	//////////////////////////////////////////////////////////////////

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.ViewData("message", "reource could not be found!")
		ctx.ViewData("error", "ERROR 404")
		ctx.View("error/404.pug")
	})

	mvc.Configure(app.Party("/"), func(app *mvc.Application) {

		app.Party("/file/").Handle(new(root.FILEController))
		app.Party("/auth").Handle(new(api.AUTHController))
	}).Handle(new(routes.ROOTController))

	mvc.Configure(app.Party("/api/v1"), func(app *mvc.Application) {

		//app.Party("/file").Handle(new(api.FILEController))
		app.Party("/file").Handle(new(api.FILEController))
		app.Party("/auth").Handle(new(api.AUTHController))
	}).Handle(new(routes.APIController))

	return app
}

func main() {
	app := initialize()

	app.Run(iris.Addr(":8080"),
		iris.WithPostMaxMemory(32<<20),
		iris.WithoutServerError(iris.ErrServerClosed))
}
