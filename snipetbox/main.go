package main

import (
	"log"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request){
  if r.URL.Path !="/"{
    http.NotFound(rw,r)
    return
  }
  rw.Write([]byte("Bem vindo ao SnipetBox."))
  
}
func showSnippet(rw http.ResponseWriter, r *http.Request){
  rw.Write([]byte("Mostrar um snippet específico.")) 
}
func createSnippet(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    rw.Header().Set("Allow", "POST")
    http.Error(rw, "Método não permitido",http.StatusMethodNotAllowed)
    return
  }
  rw.Write([]byte("criar um snippet.")) 
}  

//curl -i -X POST https://localhost:4000/snippet/create
func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", home )//criar rota
  mux.HandleFunc("/snippet", showSnippet)//criar rota
  mux.HandleFunc("/snippet/create", createSnippet)//criar rota
  
  log.Println("Inicializando servidor na porta: 4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}



//3 pilares 1-router(acessas paginas)2-handler (para onde vai a chamada) 3-servidor(recebe URL)