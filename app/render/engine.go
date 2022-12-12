package render

import (
	"lunch/app/templates"
	"lunch/public"

	"github.com/gobuffalo/buffalo/render"
)

var Engine = render.New(render.Options{
	HTMLLayout:  "application.plush.html",
	TemplatesFS: templates.FS(),
	AssetsFS:    public.FS(),
	//Helpers:     helpers.All,
})
