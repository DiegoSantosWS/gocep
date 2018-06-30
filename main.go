package main

import (
	"fmt"

	"github.com/DiegoSantosWS/restcep/connection"
	"github.com/DiegoSantosWS/restcep/routers"
)

//main start program
func main() {

	err := connection.Connection()
	if err != nil {
		fmt.Println("Erro ao abrir conexão com banco ", err.Error())
		return
	}
	routers.Routers()
}
