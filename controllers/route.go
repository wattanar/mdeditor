package controllers

import (
  "net/http"
  "html/template"
  "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  t, _ := template.ParseFiles("views/page-index.html","views/inc-header.html","views/inc-footer.html")
  t.Execute(w, nil)
}
