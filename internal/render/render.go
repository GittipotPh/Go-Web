package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"path/filepath"

	// "log"
	"net/http"

	"github.com/GittipotPh/Go-Web/internal/config"
	"github.com/GittipotPh/Go-Web/internal/models"
	"github.com/justinas/nosurf"
)

// var functions = template.FuncMap{

// }



var app *config.AppConfig



func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData , r *http.Request ) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request ,tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {

		tc = app.TemplateCache

	} else {

		tc , _ = CreateTemplateCache()
	}

	

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not find template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td , r)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}


	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template on Brownser" , err)
	}

	// parsedTemplate , _ := template.ParseFiles("./templates/" + tmpl, "./templates/baselayout.page.html")
	// err := parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("Error executing template: ", err)
	// 	return
	// }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*page.html")
	fmt.Println("line 0",pages)
	if err != nil {
		return myCache, err
	}

	for _ , page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		fmt.Println("line 1",name)
		fmt.Println("line 2",ts)
		if err != nil {
		return myCache, err
	}

	matches, err := filepath.Glob("./templates/*.layout.html")
	if err != nil {
		fmt.Println("line 3",matches)

		return myCache, err
	}

	if len(matches) > 0 {
		ts, err = ts.ParseGlob("./templates/*.layout.html")
		fmt.Println("line 4",ts)

		if err != nil {
		return myCache, err
	}
	}

	myCache[name] = ts

	fmt.Println(myCache)
}
return myCache, nil

}


// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var html *template.Template
// 	var err error

// 	// check to see if we already have a template in our cache

// 	_, inMap := tc[t]

// 	if !inMap {


// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}

// 	} else {
// 		log.Println("using cached template")
// 	}

// 	html = tc[t]

// 	err = html.Execute(w, nil)
// 	if err != nil {
// 			log.Println(err)

// 	}
// }


// func createTemplateCache(t string) error {

// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t), "./templates/baselayout.page.html",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl

// 	return nil

// }