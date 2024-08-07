package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom form struct, embeds a URL.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Valid returns true if there are no errors else false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Add(field, "This field cannot be blank.")
		return false
	}
	return true
}
