package banco

import (
	"api/api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// abre a conexao com o banco de dados
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
