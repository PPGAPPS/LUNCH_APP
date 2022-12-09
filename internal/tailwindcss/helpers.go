package tailwindcss

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/helpers/paths"
	"github.com/gobuffalo/tags/v3"
)

// htmler generates HTML source
type htmler interface {
	HTML() template.HTML
}

type helperable interface {
	SetAuthenticityToken(string)
	Append(...tags.Body)
	htmler
}

func FormHelper(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	html, err := helper(opts, help, func(opts tags.Options) helperable {
		return New(opts)
	})

	return html, err
}

func FormForHelper(model interface{}, opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	if opts == nil {
		opts = tags.Options{}
	}

	if _, ok := opts["action"]; !ok {
		act, err := paths.PathFor(model)
		if err != nil {
			return "", err
		}

		opts["action"] = act
	}

	html, err := helper(opts, help, func(opts tags.Options) helperable {
		return NewFormFor(model, opts)
	})

	return html, err
}

func helper(opts tags.Options, help hctx.HelperContext, fn func(opts tags.Options) helperable) (template.HTML, error) {
	if opts == nil {
		opts = tags.Options{}
	}

	hn := "f"
	if n, ok := opts["var"]; ok {
		hn = n.(string)
		delete(opts, "var")
	}

	if opts["errors"] == nil && help.Value("errors") != nil {
		opts["errors"] = help.Value("errors")
	}

	form := fn(opts)
	if help.Value("authenticity_token") != nil && opts["method"] != "GET" {
		form.SetAuthenticityToken(fmt.Sprint(help.Value("authenticity_token")))
	}

	help.Set(hn, form)
	s, err := help.Block()
	if err != nil {
		return "", err
	}

	form.Append(s)

	return form.HTML(), nil
}
