package main

import (
	"mineralsInKenya/gocode"
	
	"github.com/gorilla/mux"
	"html/template"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"
	
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
var combined gocode.Combined
func main() {
	//call poly to run
	//combineddata:=gocode.Poly()
	//combined=combineddata
	CallCombined()
	router := mux.NewRouter().StrictSlash(true)
	//css

	fs := http.FileServer(http.Dir("./bootstrap/"))
	router.PathPrefix("/bootstrap/").Handler(http.StripPrefix("/bootstrap/", fs))
	//
	router.HandleFunc("/",index)
	router.HandleFunc("/table",table)
	router.HandleFunc("/mineralzonesfile",readfile)
	router.HandleFunc("/file",readfile1)
	log.Fatal(http.ListenAndServe(":7000", router))
}

func index(w http.ResponseWriter, r * http.Request){
	_ = r
	handleTemplateError(w, "test2.html", nil)
}
func table(w http.ResponseWriter, r * http.Request){
	_ = r
	handleTemplateError(w, "index.html", combined)
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


func CallCombined() {
	jsonFile, err := os.Open("minerals.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened Minerals.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &combined)

}
