package main

import (
	"log"
	"net/http"
	"webapp/pkg/config"
	"webapp/pkg/handlers"
	"webapp/pkg/render"
)

func main(){

	var app config.AppConfig
	tc, err := render.CreateCache()

	if err!=nil{
		log.Fatal("Error configuring the settings")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/home",handlers.Home)
	http.HandleFunc("/about",handlers.About)
	http.ListenAndServe(":8080",nil)

}




