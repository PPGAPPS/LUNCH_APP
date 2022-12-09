package tailwindcss

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/tags/v3"
)

func buildOptions(opts tags.Options, err bool) {
	if opts["class"] == nil {
		opts["class"] = ""
	}

	if opts["tag_only"] != true {
		parts := []string{
			fmt.Sprint(opts["class"]),
			"focus:ring-indigo-500 focus:border-indigo-500 flex-1 block w-full rounded-md sm:text-sm border-gray-300",
		}

		opts["class"] = strings.Join(parts, " ")
	}

	if err {
		opts["class"] = strings.Join([]string{fmt.Sprint(opts["class"]), "is-invalid"}, " ")
	}
}

func divWrapper(opts tags.Options, fn func(opts tags.Options) tags.Body) *tags.Tag {
	divClass := opts["containerClass"]
	delete(opts, "containerClass")

	hasErrors := false
	errors := []string{}
	if divClass == nil {
		divClass = ""
	}

	if opts["errors"] != nil && len(opts["errors"].([]string)) > 0 {
		divClass = fmt.Sprintf("%v %v", divClass, "has-error")
		hasErrors = true
		errors = append(errors, opts["errors"].([]string)...)

		delete(opts, "errors")
	}

	div := tags.New("div", tags.Options{
		"class": divClass,
	})

	if opts["label"] == nil && opts["tags-field"] != nil {
		if tf, ok := opts["tags-field"].(string); ok {
			tf = strings.Join(strings.Split(tf, "."), " ")
			opts["label"] = flect.Titleize(tf)
		}
	}

	delete(opts, "tags-field")

	useLabel := opts["hide_label"] == nil
	if useLabel {
		div.Prepend(tags.New("label", tags.Options{
			"body":  opts["label"],
			"class": "block text-sm font-medium text-gray-700 mb-1",
		}))

		delete(opts, "hide_label")
	}

	delete(opts, "label")
	delete(opts, "hide_label")

	idiv := tags.New("div", tags.Options{
		"class": divClass,
	})

	buildOptions(opts, hasErrors)

	if opts["tag_only"] == true {
		return fn(opts).(*tags.Tag)
	}

	idiv.Append(fn(opts))
	div.Append(idiv)

	if !hasErrors {
		return div
	}

	idiv.Append(tags.New("p", tags.Options{
		"class": "mt-1 text-xs text-red-600 mb-2",
		"body":  strings.Join(errors, ". "),
	}))

	return div
}
