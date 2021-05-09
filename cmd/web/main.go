package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kira1108/goweb/pkg/config"
	"github.com/Kira1108/goweb/pkg/handler"
	"github.com/Kira1108/goweb/pkg/render"
)

const portNumber = ":8080"

func main() {

	// create and loading template cache.
	// associate TemplateCache to AppConfig
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true

	// create repo, associalte with AppConfig - app
	repo := handler.NewRepo(&app)

	// pass repo back to handler
	handler.NewHandlers(repo)

	// pass app to render
	render.NewTemplates(&app)

	// http.HandleFunc("/", handler.Repo.Home)
	// http.HandleFunc("/about", handler.Repo.About)

	fmt.Println("Starting web application on", portNumber)
	// http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
