package cadastro

import (
	"net/http"

	"github.com/profmugomes/mgcadastro/render"
)

func Lista(w http.ResponseWriter, r *http.Request) {
	render.Show(w, http.StatusOK, "lista.html", render.PageContext{
		Title: "Lista de Cadastro",
		Data:  []string{"Teste 1", "Teste 2"},
	})
}