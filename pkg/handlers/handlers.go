package handlers

import (
	"github.com/anishka07/bnbmono/pkg/config"
	"github.com/anishka07/bnbmono/pkg/models"
	"net/http"

	"github.com/anishka07/bnbmono/pkg/render"
)

// Repo the repository used by the handler
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a New Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	remoteIP := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIP)
	render.Templates(writer, "home.page.html", &models.DataModel{})
}

func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Anishka Mukherjee"
	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	data := &models.DataModel{
		StringMap: stringMap,
	}
	render.Templates(writer, "about.page.html", data)
}
