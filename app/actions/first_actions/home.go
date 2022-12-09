package first_actions

import (
	"lunch/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

var r = render.Engine

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}
