package handlers

import (
	"fmt"
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
	dataObject, err:= NewDataObject("api-project-377888563324")
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

