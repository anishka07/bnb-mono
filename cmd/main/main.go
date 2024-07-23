package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/anishka07/bnbmono/pkg/config"
	"github.com/anishka07/bnbmono/pkg/handlers"
	"github.com/anishka07/bnbmono/pkg/render"
	"log"
	"net/http"
	"time"
)

const port = ":9000"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	var srv *http.Server
	srv = &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	log.Println("Listening on " + port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}