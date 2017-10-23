package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func InitHandlers() {
	http.HandleFunc("/", redirectHandler)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	name, err := getLinkName(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if name == "" {
		listHandler(w, r)
		return
	}
	dataObject, err := NewDataObject("api-project-377888563324", r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	url, err := dataObject.GetURL(name)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Println("name:", name, "url:", url)
	http.Redirect(w, r, url, 302)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	dataObject, err := NewDataObject("api-project-377888563324", r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data, err := dataObject.GetListOfLinks()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	/*
	for _, val := range data {
		fmt.Fprintf(w, "name: %v\nurl:%v\n\n", val.Name, val.Url)
	}
	*/
	tmpl := template.New("index.html")
	tmpl, _ = template.ParseFiles("index.html")
	tmpl.Execute(w, data)
}

