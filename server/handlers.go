package server

import (
	"fmt"
	"html/template"
	"net/http"

	"google.golang.org/appengine"
)

var (
	projectId = "api-project-377888563324"
	indexHtml = "templates/index.html"
)

func init() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api/v1/", apiHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		listHandler(w, r)
		return
	}
	redirectHandler(w, r)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ds, err := NewDataStore(ctx, projectId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	name, err := getLinkName(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	url, err := ds.GetURL(ctx, name)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Println("name:", name, "url:", url)
	http.Redirect(w, r, url, 302)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ds, err := NewDataStore(ctx, projectId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data, err := ds.GetListOfLinks(ctx)
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

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		updateHandler(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	names := r.Form["name"]
	urls := r.Form["url"]
	if len(names) != 1 || len(urls) != 1 {
		http.Error(w, "Invalid Request", 400)
		return
	}
	name, err := getLinkName(names[0])
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if name == "" {
		http.Error(w, "Invalid Request", 400)
		return
	}
	url := urls[0]
	ctx := appengine.NewContext(r)
	ds, err := NewDataStore(ctx, projectId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	golink := Golink{
		Name: name,
		Url: url,
	}
	if err := ds.UpdateLink(ctx, golink); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "Ok")
}

