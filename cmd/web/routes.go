package main

import (
	"net/http"

	"github.com/Kira1108/goweb/pkg/config"
	"github.com/Kira1108/goweb/pkg/handler"
	"github.com/bmizerany/pat"
)

func routes(ap *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))

	return mux
}
