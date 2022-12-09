package grifts

import (
	"lunch/app"

	"github.com/gobuffalo/buffalo"
)

func init() {
	bapp, err := app.New()
	if err != nil {
		panic(err)
	}
	buffalo.Grifts(bapp)

}
