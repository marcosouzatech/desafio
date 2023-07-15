package repositorios

import (
	"api/api/src/modelos"
	"database/sql"
	"fmt"
)

type Itens struct {
	db *sql.DB
}

// NovoRepositorioDeItem cria um reposirio de usuário
func NovoRepositorioDeItens(db *sql.DB) *Itens {
	return &Itens{db}
}

// criar insere um item no banco de dados
func (repositorio Itens) Criar(item modelos.Item) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into itens (Product, Nome, Categoria, token) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(item.Product, item.Nome, item.Categoria, item.Token)
	if erro != nil {
		return 0, erro
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

// Buscar traz todos itens ou traz com base no filtro Product ou Nome
func (repositorio Itens) Buscar(ProductOuNome string) ([]modelos.Item, error) {
	ProductOuNome = fmt.Sprintf("%%%s%%", ProductOuNome)

	linhas, erro := repositorio.db.Query(
		"select id, Product, Nome, Categoria, criadoEm from itens where Product LIKE ? or Nome LIKE ?",
		ProductOuNome, ProductOuNome,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var itens []modelos.Item

	for linhas.Next() {
		var item modelos.Item

		if erro = linhas.Scan(
			&item.ID,
			&item.Product,
			&item.Nome,
			&item.Categoria,
			&item.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		itens = append(itens, item)
	}

	return itens, nil
}

// BuscaPorID traz um item do banco de dados
func (repositorio Itens) BuscarPorID(ID uint64) (modelos.Item, error) {
	linhas, erro := repositorio.db.Query(
		"select id, Product, Nome, email, criadoEm from itens where id = ? ",
		ID,
	)
	if erro != nil {
		return modelos.Item{}, erro
	}
	defer linhas.Close()

	var item modelos.Item

	if linhas.Next() {
		if erro = linhas.Scan(
			&item.ID,
			&item.Product,
			&item.Nome,
			&item.Categoria,
			&item.CriadoEm,
		); erro != nil {
			return modelos.Item{}, erro
		}
	}
	return item, nil
}

func (repositorio Itens) Atualizar(ID uint64, item modelos.Item) error {
	statement, erro := repositorio.db.Prepare(
		"update itens set Product = ?, Nome = ?, email = ?, where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(item.Product, item.Nome, item.Categoria, ID); erro != nil {
		return erro
	}
	return nil
}
