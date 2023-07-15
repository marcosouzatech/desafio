package banco

import (
	"api/api/src/config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// abre a conexao com o banco de dados
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		log.Fatal(erro)
	}
	if erro = db.Ping(); erro != nil {
		db.Close()
		log.Fatal(erro)
	}
	return db, nil
}
