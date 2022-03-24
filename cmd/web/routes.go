package main

import (
	"github.com/bmizerany/pat"
	"net/http"
	"webapp/pkg/config"
	"webapp/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler{

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about",http.HandlerFunc(handlers.Repo.Home))

	return mux

}
