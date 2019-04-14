package controllers

import (
	"fmt"
	"net/http"
	"parallelSearch/config"
	"parallelSearch/googleMapPlace"
	"regexp"
	"text/template"
)

type Place struct {
	Name     string
	Park     []byte
	Izakaya  []byte
	Yakiniku []byte
}

// templateをキャシュする
var templates = template.Must(template.ParseFiles("app/views/form.html", "app/views/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Place) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request, name string) {
	p := &Place{}
	renderTemplate(w, "form", p)
}

func execHandler(w http.ResponseWriter, r *http.Request, name string) {
	placedata := r.FormValue("name")
	yakinikuc := make(chan []byte)
	izakayac := make(chan []byte)
	parkc := make(chan []byte)
	go googleMapPlace.GooglePlaces("焼肉", placedata, yakinikuc)
	go googleMapPlace.GooglePlaces("居酒屋", placedata, izakayac)
	go googleMapPlace.GooglePlaces("公園", placedata, parkc)
	yakiniku := <-yakinikuc
	izakaya := <-izakayac
	park := <-parkc
	p := &Place{Name: placedata, Park: park, Izakaya: izakaya, Yakiniku: yakiniku}
	renderTemplate(w, "view", p)
}

var validPath = regexp.MustCompile("^/(form|exec)/$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
		}
		fmt.Println(m)
		fn(w, r, m[1])
	}
}

func StartWebServer() error {
	http.HandleFunc("/form/", makeHandler(formHandler))
	http.HandleFunc("/exec/", makeHandler(execHandler))
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
