package connection

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Db é um ponteiro do pacote sqlx
var Db *sqlx.DB

//Connection realiza conexão com banco de dados
func Connection() (err error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), "3306", os.Getenv("MYSQL_DATABASE"))
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
