package main

import (
    "html/template"
    "net/http"

    "github.com/Caronise/lenslocked.com/views"

    "github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    if err := homeView.Template.Execute(w, nil)
    if err != nil {
        panic(err)
    }
}

func contact (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    if err := contactView.Template.Execute(w, nil)
    if err != nil {
        panic(err)
    }
}


func main() {
    homeView = views.NewView("views/home.gohtml")
    contactView = views.NewView("views/contact.gohtml")

    r := mux.NewRouter()
    r.HandleFunc("/", home)
    r.HandleFunc("/contact", contact)
    http.ListenAndServe(":3000", r)
}
