package server

import (
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

var (
	projectId = "api-project-377888563324"
	indexHtml = "templates/index.html"
)

func InitHandlers() {
	http.HandleFunc("/", redirectHandler)
}

func getContext(r *http.Request) (context.Context) {
	ctx := appengine.NewContext(r)
	ctx = context.WithValue(ctx, "projectId", projectId)
	return ctx
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := getContext(r)
	name, err := getLinkName(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if name == "" {
		listHandler(w, r)
		return
	}
	url, err := getURL(ctx, name)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Println("name:", name, "url:", url)
	http.Redirect(w, r, url, 302)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	ctx := getContext(r)
	data, err := getListOfLinks(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tmpl, err := template.ParseFiles(indexHtml)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	tmpl.Execute(w, data)
}

