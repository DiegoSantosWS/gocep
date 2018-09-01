package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DiegoSantosWS/restcep/controlls"

	"github.com/DiegoSantosWS/restcep/helpers"
	"github.com/gorilla/mux"
)

//Routers cria as routas usadas
func Routers() {
	r := mux.NewRouter()
	s := r.PathPrefix("/v1").Subrouter()

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/v1/cep/", 301)
	})

	s.HandleFunc("/cep/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Entre com número cep para ser consultado"))
	})

	s.HandleFunc("/cep/{id}", controlls.Cep)

	s.HandleFunc("/validatecpf/{number}", func(w http.ResponseWriter, r *http.Request) {
		cod := mux.Vars(r)

		var retornoJS []helpers.Message
		m := helpers.Message{}
		if r.Method == "GET" {
			msg, err := helpers.ValidateCPF(cod["number"])
			if err != nil {
				m.Msg = msg
				m.Err = true
			} else {
				m.Msg = msg
			}
			retornoJS = append(retornoJS, m)

			retornoJSON, err := json.Marshal(m)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(retornoJSON)
		}
	})

	s.HandleFunc("/validatecnpj/{number}", func(w http.ResponseWriter, r *http.Request) {
		cod := mux.Vars(r)

		var retornoJS []helpers.Message
		m := helpers.Message{}
		if r.Method == "GET" {
			msg, err := helpers.ValidateCNPJ(cod["number"])
			if err != nil {
				m.Msg = msg
				m.Err = true
			} else {
				m.Msg = msg
			}
			retornoJS = append(retornoJS, m)

			retornoJSON, err := json.Marshal(m)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(retornoJSON)
		}
	})

	helpers.Runn(s)
}
