package main

import (
	
	
	"github.com/gorilla/mux"
	"html/template"
	"fmt"
	"log"
	"net/http"
	
)


var templates *template.Template

//var jsontemplates *template.Template
func init() {
	templates = template.Must(template.ParseGlob("html/*.html"))
}
func handleTemplateError(w http.ResponseWriter, templateName string, data interface{}) {
	if err := templates.ExecuteTemplate(w, templateName, data); err != nil {
		fmt.Printf("Error was %v\n", err)
		PrintAndWriteError(w, http.StatusInternalServerError, "Template error: %v", err)
		return
	}
}
func PrintAndWriteError(w http.ResponseWriter, StatusCode int, Msg string, Err error) {
	if Err != nil {
		w.WriteHeader(StatusCode)
		return
	}
	w.WriteHeader(StatusCode)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	//css

	fs := http.FileServer(http.Dir("./bootstrap/"))
	router.PathPrefix("/bootstrap/").Handler(http.StripPrefix("/bootstrap/", fs))
	//
	router.HandleFunc("/",index)
	router.HandleFunc("/mineralzonesfile",readfile)
	router.HandleFunc("/file",readfile1)
	log.Fatal(http.ListenAndServe(":7000", router))
}

func index(w http.ResponseWriter, r * http.Request){
	_ = r
	handleTemplateError(w, "test2.html", nil)
}
func readfile (w http.ResponseWriter , r * http.Request){
	//_ = r
	//handleTemplateError(w, "graph.html", nil)

	//for minerals 
	temp, err := template.ParseGlob("geojson/*.geojson")
	if err != nil {
		log.Fatalln(err)
	}
	if err := temp.ExecuteTemplate(w, "kenyaMinerals.geojson", nil); err != nil {
		fmt.Printf("Error was %v\n", err)
		PrintAndWriteError(w, http.StatusInternalServerError, "Template error: %v", err)
		return
	}
}
func readfile1 (w http.ResponseWriter , r * http.Request){
	//_ = r
	//handleTemplateError(w, "graph.html", nil)
	temp, err := template.ParseGlob("geojson/*.geojson")
	if err != nil {
		log.Fatalln(err)
	}
	if err := temp.ExecuteTemplate(w, "kenya.geojson", nil); err != nil {
		fmt.Printf("Error was %v\n", err)
		PrintAndWriteError(w, http.StatusInternalServerError, "Template error: %v", err)
		return
	}
}
