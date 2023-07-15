package rotas

import (
	"api/api/src/controllers"
	"net/http"
)

var rotasItens = []Rota{
	{
		URI:                "/",
		Metodo:             http.MethodGet,
		Funcao:             controllers.HealthCheck,
		RequerAutenticacao: false,
	},
	{
		URI:                "/itens",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarItem,
		RequerAutenticacao: false,
	},
	{
		URI:                "/itens",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarItens,
		RequerAutenticacao: false,
	},
	{
		URI:                "/itens/{itemId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarItem,
		RequerAutenticacao: false,
	},
	{
		URI:                "/itens/{itemId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarItem,
		RequerAutenticacao: false,
	},
	{
		URI:                "/itens/{itemId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarItem,
		RequerAutenticacao: false,
	},
}
