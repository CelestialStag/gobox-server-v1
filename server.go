/*
server.go
@package main
contains the main entry point for the program
*/
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

/*
initialize - creates a iris application
@returns a new iris application
*/
func initialize() *iris.Application {

	app := iris.New()

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// View Engine
	app.RegisterView(iris.Pug("./views", ".pug").Reload(true))

	// Static files
	app.StaticWeb("/static", "./static")

	// Routes
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.ViewData("message", "This page does not exist!")
		ctx.ViewData("error", "ERROR 404")
		ctx.View("error/404.pug")
	})

	mvc.Configure(app.Party("/"), func(app *mvc.Application) {

		app.Party("/f").Handle(new(root.FILEController))
		app.Party("/a").Handle(new(api.AUTHController))
	}).Handle(new(routes.ROOTController))

	mvc.Configure(app.Party("/api/v1"), func(app *mvc.Application) {

		//app.Party("/file").Handle(new(api.FILEController))
		app.Party("/f").Handle(new(api.FILEController))
		app.Party("/a").Handle(new(api.AUTHController))
	}).Handle(new(routes.APIController))

	return app
}

/*
main - Gets the iris object and runs it
*/
func main() {
	app := initialize()

	app.Run(iris.Addr(":4040"),
		iris.WithPostMaxMemory(5e+9),
		iris.WithoutServerError(iris.ErrServerClosed))
}
