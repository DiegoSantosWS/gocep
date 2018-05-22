#CONSULTAR CEP
API PARA CONSULTAR CEP, RETORNA DADOS DE ENDEREÇO DO CEP SOLICITADO

# FORMA DE USO

``` bash
$ git clone -u git@github.com:DiegoSantosWS/gocep.git
```
Crie um banco de dados para salava os dados caso queira salvar os dados consultados, caso não queira criar um banco, edite o aquivo


* [v1](https://github.com/DiegoSantosWS/gocep/blob/master/controlls/v1.go#L47) - comente a linha 47 conforme o exemplo.

No arquivo main.go comente a linha [L24](https://github.com/DiegoSantosWS/gocep/blob/master/main.go#L24) a L28 para retirar a conexão com banco de dados.


basta comentar a linha 47

```go 
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
        //helpers.VerificaDados(dadosCEP) comente esta linha

        w.Header().Set("Content-Type", "application/json")
        w.Write(corpo)
    }
}
```


# Acesso

Pode acessar via terminal com CRUL como mostra do exemplo abaixo ou via browser
``` bash
$ curl GET http://localhost:5000/v1/cep/30626680
```