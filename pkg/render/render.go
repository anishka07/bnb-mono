package render

import (
	"bytes"
	"github.com/anishka07/bnbmono/pkg/config"
	"github.com/anishka07/bnbmono/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

var (
	TemplateCache = make(map[string]*template.Template)
	mu            sync.Mutex
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.DataModel) *models.DataModel {
	return td
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	mu.Lock()
	defer mu.Unlock()

	pages, err := filepath.Glob("templates/*.page.html")
	if err != nil {
		return TemplateCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).ParseFiles(page)
		if err != nil {
			return TemplateCache, err
		}

		matches, err := filepath.Glob("templates/*.layout.html")
		if err != nil {
			return TemplateCache, err
		}

		if len(matches) > 0 {
			t, err = t.ParseGlob("templates/*layout.html")
			if err != nil {
				return TemplateCache, err
			}
		}
		TemplateCache[name] = t
	}
	return TemplateCache, err
}

func Templates(w http.ResponseWriter, tmpl string, td *models.DataModel) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache.")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}
