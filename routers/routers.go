package routers

import (
	"net/http"

	"github.com/DiegoSantosWS/restcep/controlls"

	"github.com/DiegoSantosWS/restcep/helpers"
	"github.com/gorilla/mux"
)

//Routers cria as routas usadas
func Routers() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá vamos consultar cep?"))
	})

	r.HandleFunc("/v1/cep/{id}", controlls.Cep)
	helpers.Runn(r)
}
