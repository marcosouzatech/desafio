package controllers

import (
	"api/api/src/banco"
	"api/api/src/modelos"
	"api/api/src/repositorios"
	"api/api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarItem insere um usuário no banco
func CriarItem(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var item modelos.Item
	if erro = json.Unmarshal(corpoRequest, &item); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = item.Preparar("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeItens(db)
	item.ID, erro = repositorio.Criar(item)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, item)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		// Registrar o evento GET no stdout
		w.Write([]byte("Health Check OK!"))
		log.Println("INFO: Requisição processada com sucesso no endpoint /")
	}
}
func BuscarItens(w http.ResponseWriter, r *http.Request) {
	ProductOuNome := strings.ToLower(r.URL.Query().Get("item"))
	log.Println("INFO: Consulta a lista de /Itens")

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		log.Println("WARNING: Falha ao se conectar com banco de dados")
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeItens(db)
	itens, erro := repositorio.Buscar(ProductOuNome)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		log.Println("WARNING: Falha ao criar novo item")
		return
	}
	respostas.JSON(w, http.StatusOK, itens)

}

// Buscar um item salvo no banco
func BuscarItem(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	itemID, erro := strconv.ParseUint(parametros["itemId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeItens(db)
	item, erro := repositorio.BuscarPorID(itemID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, item)
}

func AtualizarItem(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	itemID, erro := strconv.ParseUint(parametros["itemId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var item modelos.Item
	if erro = json.Unmarshal(corpoRequisicao, &item); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = item.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeItens(db)
	if erro = repositorio.Atualizar(itemID, item); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar um Item"))
}
