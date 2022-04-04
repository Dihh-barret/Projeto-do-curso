package main

import (
	"fmt"
	"net/http"
  "strconv"
)

func home(rw http.ResponseWriter, r *http.Request){
  if r.URL.Path !="/"{
    http.NotFound(rw,r)
    return
  }
  rw.Write([]byte("Bem vindo ao SnipetBox."))
  
}
//https://projeto-do-curso.poiuk.repl.co/snippet?id=123
func showSnippet(rw http.ResponseWriter, r *http.Request){
  id,err := strconv.Atoi(r.URL.Query().Get("id"))//retorna valor da query string id
  if err != nil || id < 1{
    http.NotFound(rw, r)
    return
  }
fmt.Fprintf(rw, "Exibir o Snippet de ID: %d", id)  
  //rw.Write([]byte("Mostrar um snippet específico.")) 
}

func createSnippet(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    rw.Header().Set("Allow", "POST")
    http.Error(rw, "Método não permitido",http.StatusMethodNotAllowed)
    return
  }
  rw.Write([]byte("criar um snippet.")) 
}  