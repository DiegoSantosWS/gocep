package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/DiegoSantosWS/restcep/connection"
	"github.com/DiegoSantosWS/restcep/model"
	"github.com/DiegoSantosWS/restcep/routers"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	yamlFile, err := ioutil.ReadFile("db.yml")
	if err != nil {
		log.Fatal(err.Error())
	}
	conf := &model.Config{}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = connection.Connection(conf)
	if err != nil {
		fmt.Println("Erro ao abrir conexão com banco ", err.Error())
		return
	}
	routers.Routers()
}
