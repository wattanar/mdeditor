package main

import (
  "net/http"
  "log"
  "github.com/julienschmidt/httprouter"
  c "github.com/wattanar/mdeditor/controllers"
)

func main()  {
  // Create New Router
  router := httprouter.New()
  // Router
  router.GET("/", c.Index)
  // Setup Static Path
  router.NotFound = http.FileServer(http.Dir("assets"))
  // Listening on port 8080
  log.Fatal(http.ListenAndServe(":9000", router))
}
