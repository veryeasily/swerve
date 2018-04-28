package main

import (
  "fmt"
  "net/http"
  "log"
  "io"
  "github.com/gorilla/mux"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
  fmt.Printf("Hello world!\n")
  r := mux.NewRouter()
  r.HandleFunc("/{name}", Chain(HomeHandler, SetMIMEType("application/json")))
  log.Fatal(http.ListenAndServe(":8080", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  w.WriteHeader(http.StatusOK)
  CreateDb(vars["name"])
  io.WriteString(w, `{"hi": "`+vars["name"]+`"}`)
}
