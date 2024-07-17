package handlers

import (
	// "html/template"
	"net/http"

	"github.com/GittipotPh/Go-Web/pkg/models"
	"github.com/GittipotPh/Go-Web/pkg/config"
	"github.com/GittipotPh/Go-Web/pkg/render"
	// "fmt"
)



var Repo *Repository

var data *models.TemplateData


type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// every type *Repository has access to the Handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){

	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	data = &models.TemplateData{
		StringMap: make(map[string]string),
	}

	data.StringMap["testHome"] = "Hello, world!"


	render.RenderTemplate(w,"home.page.html", data )


}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter , r *http.Request){
	data = &models.TemplateData{
		StringMap : make(map[string]string),

	}
	data.StringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	data.StringMap["remote_ip"] = remoteIP


	render.RenderTemplate(w, "about.page.html", data )



	
}

// addValues adds two integers and return the sum
func AddValues(x , y int) (int, error) {
	sum := x + y
	return sum, nil
}

