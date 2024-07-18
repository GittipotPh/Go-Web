package handlers

import (
	// "html/template"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GittipotPh/Go-Web/internal/config"
	"github.com/GittipotPh/Go-Web/internal/models"
	"github.com/GittipotPh/Go-Web/internal/render"
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


	render.RenderTemplate(w,r ,"home.page.html", data )


}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter , r *http.Request){
	data = &models.TemplateData{
		StringMap : make(map[string]string),

	}
	data.StringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	data.StringMap["remote_ip"] = remoteIP


	render.RenderTemplate(w, r ,"about.page.html", data )
	
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r ,"make-reservation.page.html", data)
	
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r ,"generals.page.html", data)
	
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r ,"majors.page.html", data)


}
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r ,"search-availability.page.html", data)


}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r ,"contact.page.html", data)
	
	
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request){
	// w.WriteHeader("200")

	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))


}

type jsonResponse struct {
	OK bool `json:"ok"`
	Message string `json:"message"`

}


func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request){

	resp := jsonResponse{
		OK: true,
		Message: "Available!" ,
	}

	out, err := json.MarshalIndent(resp,"", "   ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	// w.WriteHeader("200")


}
// addValues adds two integers and return the sum
func AddValues(x , y int) (int, error) {
	sum := x + y
	return sum, nil
}

