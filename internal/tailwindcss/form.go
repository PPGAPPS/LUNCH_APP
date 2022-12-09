package tailwindcss

import (
	"github.com/gobuffalo/tags/v3"
	"github.com/gobuffalo/tags/v3/form"
)

// Form is the bootstrap version of Form
type Form struct {
	*form.Form
}

// New creates a bootstrap Form from passed options
func New(opts tags.Options) *Form {
	return &Form{form.New(opts)}
}

// CheckboxTag builds a bootstrap checkbox with passed options
func (f Form) CheckboxTag(opts tags.Options) *tags.Tag {
	return f.Form.CheckboxTag(opts)
}

// InputTag builds a bootstrap input[type=text] with passed options
func (f Form) InputTag(opts tags.Options) *tags.Tag {
	return f.Form.InputTag(opts)
}

// FileTag builds a bootstrap input[type=file] with passed options
func (f Form) FileTag(opts tags.Options) *tags.Tag {
	return f.Form.FileTag(opts)
}

// RadioButton builds a bootstrap input[type=radio] with passed options
func (f Form) RadioButton(opts tags.Options) *tags.Tag {
	return f.RadioButtonTag(opts)
}

// RadioButtonTag builds a bootstrap input[type=radio] with passed options
func (f Form) RadioButtonTag(opts tags.Options) *tags.Tag {
	return f.Form.RadioButtonTag(opts)
}

// SelectTag builds a bootstrap select with passed options
func (f Form) SelectTag(opts tags.Options) *tags.Tag {
	return f.Form.SelectTag(opts)
}

// TextArea builds a bootstrap textarea with passed options
func (f Form) TextArea(opts tags.Options) *tags.Tag {
	return f.TextAreaTag(opts)
}

// TextAreaTag builds a bootstrap textarea with passed options
func (f Form) TextAreaTag(opts tags.Options) *tags.Tag {
	return f.Form.TextAreaTag(opts)
}

// HiddenTag adds a hidden input to the form
func (f Form) HiddenTag(opts tags.Options) *tags.Tag {
	return f.Form.HiddenTag(opts)
}
