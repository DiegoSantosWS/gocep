package controlls

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/DiegoSantosWS/restcep/helpers"
	"github.com/DiegoSantosWS/restcep/model"
	"github.com/gorilla/mux"
)

//Cep retorna os dados consultados
func Cep(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	url := "https://api.postmon.com.br/v1/cep/" + id + ""
	clients := &http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("[ERRO] Houve um erro ao criar um request no servidor: ", err.Error())
		return
	}

	request.SetBasicAuth("ws", "ws")
	resp, err := clients.Do(request)
	if err != nil {
		log.Fatal("[ERRO] Houve um erro ao abrir a pagina do servidor: ", err.Error())
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("[ERRO] Houve um erro ao ler o conteúdo da pagina do servidor. Erro: ", err.Error())
		}
		dadosCEP := &model.DadosCep{}
		err = json.Unmarshal(corpo, dadosCEP)
		if err != nil {
			log.Fatal("[ERRO] Houve um erro ao converter o retorno JSON do servidor. ERRO: ", err.Error())
		}
		helpers.VerificaDados(dadosCEP)

		w.Header().Set("Content-Type", "application/json")
		w.Write(corpo)
	}
}
