package handlers

import (
	"fmt"
	"net/http"
)

func InitHandlers() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/*", redirectHandler)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	name, err := getLinkName(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
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
	for _, entry := range data {
		fmt.Fprintf(w, "name: %v\nlink: %v\n\n", entry.Name, entry.Url)
	}
}

