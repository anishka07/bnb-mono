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
	render.Templates(writer, "home.page.gohtml", &models.DataModel{})
}

func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Anishka Mukherjee"
	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	data := &models.DataModel{
		StringMap: stringMap,
	}
	render.Templates(writer, "about.page.gohtml", data)
}

func (m *Repository) Reservation(writer http.ResponseWriter, request *http.Request) {
	render.Templates(writer, "make-reservation.page.gohtml", &models.DataModel{})
}

func (m *Repository) Generals(writer http.ResponseWriter, request *http.Request) {
	render.Templates(writer, "generals.page.gohtml", &models.DataModel{})
}

func (m *Repository) Majors(writer http.ResponseWriter, request *http.Request) {
	render.Templates(writer, "majors.page.gohtml", &models.DataModel{})
}

func (m *Repository) SearchAvailability(writer http.ResponseWriter, request *http.Request) {
	render.Templates(writer, "search-availability.page.gohtml", &models.DataModel{})
}

func (m *Repository) Contact(writer http.ResponseWriter, request *http.Request) {
	render.Templates(writer, "contact.page.gohtml", &models.DataModel{})
}
