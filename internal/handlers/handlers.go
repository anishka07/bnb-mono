package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anishka07/bnbmono/internal/config"
	"github.com/anishka07/bnbmono/internal/models"
	"github.com/anishka07/bnbmono/internal/render"
	"log"
	"net/http"
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
	render.Templates(writer, request, "home.page.gohtml", &models.DataModel{})
}

func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Anishka Mukherjee"
	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	data := &models.DataModel{
		StringMap: stringMap,
	}
	render.Templates(writer, request, "about.page.gohtml", data)
}

func (m *Repository) Reservation(writer http.ResponseWriter, request *http.Request) {
	render.Templates(writer, request, "make-reservation.page.gohtml", &models.DataModel{})
}

func (m *Repository) Generals(writer http.ResponseWriter, request *http.Request) {
	render.Templates(writer, request, "generals.page.gohtml", &models.DataModel{})
}

func (m *Repository) Majors(writer http.ResponseWriter, request *http.Request) {
	render.Templates(writer, request, "majors.page.gohtml", &models.DataModel{})
}

func (m *Repository) SearchAvailability(writer http.ResponseWriter, request *http.Request) {
	render.Templates(writer, request, "search-availability.page.gohtml", &models.DataModel{})
}

func (m *Repository) PostAvailability(writer http.ResponseWriter, request *http.Request) {
	start := request.Form.Get("start")
	end := request.Form.Get("end")

	_, err := writer.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s ", start, end)))
	if err != nil {
		return
	}
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok:      false,
		Message: "available",
	}
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		return
	}
}

func (m *Repository) Contact(writer http.ResponseWriter, request *http.Request) {
	render.Templates(writer, request, "contact.page.gohtml", &models.DataModel{})
}
