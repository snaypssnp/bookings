package handlers

import (
	"fmt"
	"github.com/snaypssnp/bookings/pkg/config"
	"github.com/snaypssnp/bookings/pkg/models"
	"github.com/snaypssnp/bookings/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	fmt.Println(remoteIP)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	stringMap["test"] = "Hello, again."
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{StringMap: stringMap})
}
