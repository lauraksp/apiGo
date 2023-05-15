package main

import "projetogo/pkg/task.go"
import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Essa é minha API de tarefas", r.URL.Path[1:])
}

// Iniciar servidor HTTP basico.
func main() {
	http.HandleFunc("/", handler){
		tasks
	}
	log.Fatal(http.ListenAndServe(":8000", nil))
}
