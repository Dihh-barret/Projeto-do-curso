package main

import (
//	"fmt"
  "log"
	"net/http"
//  "strconv"
)


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