package render

import (
	"bytes"
	"github.com/anishka07/bnbmono/internal/config"
	"github.com/anishka07/bnbmono/internal/models"
	"github.com/justinas/nosurf"
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

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	mu.Lock()
	defer mu.Unlock()

	pages, err := filepath.Glob("templates/*.page.gohtml")
	if err != nil {
		return TemplateCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).ParseFiles(page)
		if err != nil {
			return TemplateCache, err
		}

		matches, err := filepath.Glob("templates/*.layout.gohtml")
		if err != nil {
			return TemplateCache, err
		}

		if len(matches) > 0 {
			t, err = t.ParseGlob("templates/*layout.gohtml")
			if err != nil {
				return TemplateCache, err
			}
		}
		TemplateCache[name] = t
	}
	return TemplateCache, err
}

func Templates(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
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
	td = AddDefaultData(td, r)
	err := t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}
