package models

import "github.com/anishka07/bnbmono/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      interface{}
	CSRFToken string
	Warning   string
	Error     string
	Form      *forms.Form
}
