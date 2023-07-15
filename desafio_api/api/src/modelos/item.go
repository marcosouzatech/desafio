package modelos

import (
	"errors"
	"strings"
	"time"
)

type Item struct {
	ID        uint64    `json:"id",omitempty"`
	Product   string    `json:"Product,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Categoria string    `json:"Categoria,omitempty"`
	Token     string    `json:"token,omitempty"`
	CriadoEm  time.Time `json:"CriadoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (item *Item) Preparar(etapa string) error {
	if erro := item.validar(etapa); erro != nil {
		return erro
	}

	item.formatar()
	return nil
}

func (item *Item) validar(etapa string) error {
	if item.Product == "" {
		return errors.New("O Product é obrigatório e não pode estar em branco")
	}
	if item.Nome == "" {
		return errors.New("O Nome é obrigatório e não pode estar em branco")
	}
	if item.Categoria == "" {
		return errors.New("A Categoria é obrigatória e não pode estar em branco")
	}
	if etapa == "cadastro" && item.Token == "" {
		return errors.New("O token é obrigatório e não pode estar em branco")
	}
	return nil
}
func (item *Item) formatar() {
	item.Product = strings.TrimSpace(item.Product)
	item.Nome = strings.TrimSpace(item.Nome)
	item.Categoria = strings.TrimSpace(item.Categoria)
}
