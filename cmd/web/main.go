package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GittipotPh/Go-Web/internal/config"
	"github.com/GittipotPh/Go-Web/internal/handlers"
	"github.com/GittipotPh/Go-Web/internal/render"
	"github.com/alexedwards/scs/v2"
	// "path/filepath"
)

var app config.AppConfig
var session *scs.SessionManager


const portNumber = ":8080" 

func main() {

	app.InProduction = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache();
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// fmt.Println(app)
	// fmt.Println(app.TemplateCache)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About )
	
	fmt.Println((fmt.Sprintf("Starting application on port %s", portNumber)))

	// _ = http.ListenAndServe(portNumber, nil)

	serving := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = serving.ListenAndServe()
	log.Fatal(err)

}