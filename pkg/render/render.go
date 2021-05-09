package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Kira1108/goweb/pkg/config"
	"github.com/Kira1108/goweb/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Can not get Template")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// CreatesTemplateCache creaes a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	// create a cache map key - string, value -  pointer to a template
	myCache := map[string]*template.Template{}

	// get all page files
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// for each page, use layout
	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("page is currently", page)

		// parse page
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// if there are some layout files,
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		// then parse layouts
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		// store the current page into cache
		myCache[name] = ts

	}

	return myCache, nil

}
