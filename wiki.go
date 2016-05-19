package main

import (
	//"fmt"
	"net/http"
	//"os"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	//	http.HandleFunc("/", viewHandler)
	http.Handle("/", &templateHandler{filename: "hostname.html"})
	http.ListenAndServe(":8050", nil)
}
