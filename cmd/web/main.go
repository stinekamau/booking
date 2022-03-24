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
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	server := &http.Server{
		Addr: ":8080",
		Handler: routes(&app),
	}

	server.ListenAndServe()


	//http.HandleFunc("/home",handlers.Repo.Home)
	//http.HandleFunc("/about",handlers.Repo.About)
	//http.ListenAndServe(":8080",nil)

}




