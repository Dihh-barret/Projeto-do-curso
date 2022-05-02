package main

//go run cmd/web/*
import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func(app *application) home(rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    app.notFound(rw)
    return
  } 

  files := []string{
    "./ui/html/home.page.tmpl.html",
    "./ui/html/base.layout.tmpl.html",
    "./ui/html/footer.partial.tmpl.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, nil)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

//http://localhost:4000/snippet?id=123
func(app *application) showSnippet(rw http.ResponseWriter, r *http.Request){
  id,err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    app.notFound(rw)
    return
  }
  fmt.Fprintf(rw, "Exibir o Snippet de ID: %d", id)
}

func(app *application) createSnippet(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    rw.Header().Set("Allow","POST")
    app.clientError(rw, http.StatusMethodNotAllowed)
    return
  }
  
  rw.Write([]byte("Criar novo snippet"))
}