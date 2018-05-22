package helpers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//Runn Inicia o servidor
func Runn(r *mux.Router) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatal("[ERRO RUNN] Erro ao instanciar o servidor")
	}
}
