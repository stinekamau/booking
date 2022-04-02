package main

import (
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
	"webapp/pkg/config"
	"webapp/pkg/handlers"
	"webapp/pkg/render"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// Set to true in production mode
	app.InProduction = false

	tc, err := render.CreateCache()
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// Add the session instance to the config
	app.Session = session
	
	if err != nil {
		log.Fatal("Error configuring the settings")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	server.ListenAndServe()

	//http.HandleFunc("/home",handlers.Repo.Home)
	//http.HandleFunc("/about",handlers.Repo.About)
	//http.ListenAndServe(":8080",nil)

}
