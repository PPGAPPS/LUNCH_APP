package tailwindcss

import (
	"fmt"

	"github.com/gobuffalo/tags/v3"
	"github.com/gobuffalo/tags/v3/form"
	"github.com/gobuffalo/validate/v3/validators"
)

// FormFor is the FormFor version for Tailwindcss.
type FormFor struct {
	*form.FormFor
}

func (f FormFor) buildOptions(field string, opts tags.Options) tags.Options {
	opts["tags-field"] = field
	fieldName := validators.GenerateKey(field)
	if err := f.Errors.Get(fieldName); err != nil {
		opts["errors"] = err
	}

	return opts
}

// CheckboxTag adds a checkbox to a form wrapped with a form-control and a label.
func (f FormFor) CheckboxTag(field string, opts tags.Options) *tags.Tag {
	opts["class"] = fmt.Sprintf("%v %v", "focus:ring-blue-500 h-5 w-5 text-blue-600 border-gray-300 rounded", opts["class"])
	value := opts["value"]
	if value == nil {
		value = "true"
	}

	checked := opts["checked"]
	delete(opts, "checked")

	opts["type"] = "checkbox"
	opts["name"] = field
	opts["value"] = value
	chb := tags.New("input", opts)

	if checked != nil {
		chb.Checked = checked.(bool)
	}

	if opts["tag_only"] == true {
		return chb
	}

	fdiv := tags.New("div", tags.Options{
		"class": "h-5 flex items-center text-lg",
		"body":  chb,
	})

	if value == "true" {
		fdiv.Append(tags.New("input", tags.Options{
			"type":  "hidden",
			"value": "false",
			"name":  field,
		}))
	}

	ldiv := tags.New("div", tags.Options{
		"class": "ml-3 text-sm",
	})

	label := tags.New("label", tags.Options{
		"class": "font-medium text-sixth-700",
		"for":   field,
		"body":  opts["label"],
	})

	ldiv.Append(label)
	ldiv.Append(tags.New("p", tags.Options{
		"class": "text-sixth-500",
		"body":  opts["description"],
	}))

	inner := tags.New("div", tags.Options{
		"class": "relative flex items-start",
	})

	inner.Prepend(fdiv)
	inner.Append(ldiv)

	container := tags.New("div", tags.Options{
		"class": opts["containerClass"],
		"body":  inner,
	})

	return container
}

// InputTag builds an input[type=text] by default wrapped with a form-control and a label.
func (f FormFor) InputTag(field string, opts tags.Options) *tags.Tag {
	opts = f.buildOptions(field, opts)
	if opts["type"] == "hidden" {
		return f.HiddenTag(field, opts)
	}

	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.InputTag(field, opts)
	})
}

// HiddenTag adds an input[type=hidden] to the formFor.
func (f FormFor) HiddenTag(field string, opts tags.Options) *tags.Tag {
	opts = f.buildOptions(field, opts)

	return f.FormFor.HiddenTag(field, opts)
}

// FileTag adds a bootstrap input[type=file] wrapped with a form-control and a label.
func (f FormFor) FileTag(field string, opts tags.Options) *tags.Tag {
	opts = f.buildOptions(field, opts)

	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.FileTag(field, opts)
	})
}

// RadioButton adds a bootstrap input[type=radio] wrapped with a form-control and a label.
func (f FormFor) RadioButton(field string, opts tags.Options) *tags.Tag {
	// TODO: errors and field name
	value := opts["value"]
	if value == nil {
		value = "true"
	}

	opts["class"] = fmt.Sprintf("%v %v", "focus:ring-second-900 h-5 w-5 text-second-900 border-sixth-400", opts["class"])
	chb := tags.New("input", tags.Options{
		"type":  "radio",
		"name":  field,
		"value": value,
		"class": opts["class"],
	})

	if ch, ok := opts["checked"].(bool); ok {
		chb.Checked = ch
	}

	if opts["tag_only"] == true {
		return chb
	}

	fdiv := tags.New("div", tags.Options{
		"class": "flex items-center h-5",
		"body":  chb,
	})

	ldiv := tags.New("div", tags.Options{
		"class": "ml-3 text-sm",
	})

	label := tags.New("label", tags.Options{
		"class": "text-sm",
		"for":   field,
		"body":  opts["label"],
	})

	ldiv.Append(label)
	ldiv.Append(tags.New("p", tags.Options{
		"class": "text-sixth-500",
		"body":  opts["description"],
	}))

	inner := tags.New("div", tags.Options{
		"class": "relative flex items-start",
	})

	inner.Prepend(fdiv)
	inner.Append(ldiv)

	container := tags.New("div", tags.Options{
		"class": opts["containerClass"],
		"body":  inner,
	})

	return container
}

// RadioButtonTag adds a bootstrap input[type=radio] wrapped with a form-control and a label.
func (f FormFor) RadioButtonTag(field string, opts tags.Options) *tags.Tag {
	opts = f.buildOptions(field, opts)

	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.RadioButtonTag(field, opts)
	})
}

// SelectTag adds a bootstrap select tag wrapped with a form-control and a label.
func (f FormFor) SelectTag(field string, opts tags.Options) *tags.Tag {
	opts = f.buildOptions(field, opts)

	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.SelectTag(field, opts)
	})
}

// TextArea adds a bootstrap textarea tag wrapped with a form-control and a label.
func (f FormFor) TextArea(field string, opts tags.Options) *tags.Tag {
	return f.TextAreaTag(field, opts)
}

// TextAreaTag adds a bootstrap textarea tag wrapped with a form-control and a label.
func (f FormFor) TextAreaTag(field string, opts tags.Options) *tags.Tag {
	opts = f.buildOptions(field, opts)

	return divWrapper(opts, func(o tags.Options) tags.Body {
		return f.FormFor.TextAreaTag(field, opts)
	})
}

// SubmitTag returns a tag for input type submit without wrapping.
func (f FormFor) SubmitTag(value string, opts tags.Options) *tags.Tag {
	return f.FormFor.SubmitTag(value, opts)
}

// NewFormFor builds a form for a passed model.
func NewFormFor(model interface{}, opts tags.Options) *FormFor {
	return &FormFor{
		FormFor: form.NewFormFor(model, opts),
	}
}
