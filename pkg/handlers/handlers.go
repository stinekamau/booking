package handlers

import (
	"net/http"
	"webapp/models"
	"webapp/pkg/config"
	"webapp/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		a,
	}
}

// NewHandler Sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.DisplayTemplates(w, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some logic
	stringMap := make(map[string]string)
	stringMap["zip"] = "Monty Python"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remoteIP"] = remoteIP
	render.DisplayTemplates(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
