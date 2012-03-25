package projecttest

import (
//    "fmt"
    "http"
    "template"
)

func init() {
    http.HandleFunc("/", handler)
}

var topTemplate = template.Must(template.New("top").ParseFile("projecttest/top.html"))

func handler(w http.ResponseWriter, r *http.Request) {
    if err := topTemplate.Execute(w, nil); err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError)
        return
    }
}