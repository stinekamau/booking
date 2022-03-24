package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"webapp/pkg/config"
)

var functions = template.FuncMap{

}

var myCache  map[string] *template.Template

var app *config.AppConfig

func NewTemplates(a *config.AppConfig){
	app = a
}

func DisplayTemplates(w http.ResponseWriter, tmpl string){
	tc := app.TemplateCache


	t, ok  := tc[tmpl]

	if !ok{
		log.Fatal("err")
	}


	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)
	if err!=nil{
		fmt.Errorf("error writing to buffer")
	}




}

func CreateCache()(map[string] *template.Template,error){

	myCache := map[string] *template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	checkError(err, "Error Pattern Matching the file")

	for _, page  := range pages{
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		checkError(err,"Error creating a new Template")

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		checkError(err, "Error checking for the base layout")

		if len(matches) > 0 {
			_, err := ts.ParseGlob("./templates/*.layout.tmpl")
			checkError(err,"Error executing the template file")
		}

		myCache[name] = ts
	}


	return myCache,nil


}

func checkError(err error, message string) {
	if err != nil{
		log.Println(fmt.Errorf(message))
	}

}
