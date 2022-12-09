package app

import (
	"lunch/app/actions/first_actions"
	"lunch/public"
	"net/http"

	"github.com/gobuffalo/buffalo"
	csrf "github.com/gobuffalo/mw-csrf"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
)

func setRoutes(app *buffalo.App) {
	app.Use(forceSSL())
	app.Use(paramlogger.ParameterLogger)
	app.Use(csrf.New)

	// Wraps each request in a transaction.
	//   c.Value("tx").(*pop.Connection)
	// Remove to disable this.
	// app.Use(popmw.Transaction(models.DB))
	// Setup and use translations:
	// app.Use(translations())

	app.GET("/", first_actions.HomeHandler)
	app.ServeFiles("/", http.FS(public.FS()))
}
