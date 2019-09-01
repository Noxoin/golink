package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var (
	projectId = "api-project-377888563324"
	kind      string
	tmpls     map[string]*template.Template
)

func loadEnvVars() {
	if v := os.Getenv("DS_KIND"); v != "" {
		kind = v
	}
}

func init() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/admin/", adminHandler)
	http.HandleFunc("/api/v1/", apiHandler)
	tmpls = make(map[string]*template.Template)
	tmpls["index"] = template.Must(template.ParseFiles(
		"templates/base.html", "templates/index.html"))
	tmpls["admin"] = template.Must(template.ParseFiles(
		"templates/base.html", "templates/admin.html"))

	loadEnvVars()
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		listHandler(w, r)
		return
	}
	redirectHandler(w, r)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ds, err := NewDataStore(ctx, projectId, kind)
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
	ds.IncrementCount(ctx, name)
	http.Redirect(w, r, url, 302)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ds, err := NewDataStore(ctx, projectId, kind)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	res, err := ds.GetListOfLinks(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data := &struct {
		Title   string
		Golinks []*Golink
	}{
		Title:   "index",
		Golinks: res,
	}
	tmpls["index"].ExecuteTemplate(w, "base", data)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ds, err := NewDataStore(ctx, projectId, kind)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	res, err := ds.GetListOfLinks(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data := &struct {
		Title   string
		Golinks []*Golink
	}{
		Title:   "admin",
		Golinks: res,
	}
	tmpls["admin"].ExecuteTemplate(w, "base", data)
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
		http.Error(w, "Expecting only one name and url.", 400)
		return
	}
	name, err := getLinkName("/" + names[0])
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	url := urls[0]
	ctx := r.Context()
	ds, err := NewDataStore(ctx, projectId, kind)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	_, err = ds.GetURL(ctx, name)
	if err == nil {
		http.Error(w, "Link Name Already Exists.", 400)
		return
	}
	golink := Golink{
		Name: name,
		Url:  url,
	}
	if err := ds.UpdateLink(ctx, golink); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "Ok")
}
