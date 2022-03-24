package handlers

import (
	"net/http"
	"webapp/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request){
	render.DisplayTemplates(w, "home.page.tmpl")

}

func About(w http.ResponseWriter, r *http.Request){
	render.DisplayTemplates(w, "about.page.tmpl")
}


