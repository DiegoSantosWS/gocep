package connection

import (
	"fmt"
	"log"

	"github.com/DiegoSantosWS/restcep/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Db é um ponteiro do pacote sqlx
var Db *sqlx.DB

//Connection realiza conexão com banco de dados
func Connection(conf *model.Config) (err error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Pass, conf.Host, conf.Port, conf.DB)
	Db, err = sqlx.Open("mysql", uri)
	if err != nil {
		log.Fatal("[ERRO DB] erro ao conectar com banco de dados: ", err.Error())
		return
	}
	if err = Db.Ping(); err != nil {
		log.Fatalf("Erro ao conectar com banco de dados: %s\n", err.Error())
	}
	return
}
