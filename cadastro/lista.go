package cadastro

import (
	"net/http"

	"github.com/profmugomes/mgcadastro/render"
)

func Lista(w http.ResponseWriter, r *http.Request) {
	render.Show(w, http.StatusOK, "lista.html", map[string]any{
		"Title": "Lista de Cadastro",
		"Outro": []string{"Teste 1", "Teste 2"},
		"Data":  []map[string]string{{"Titulo": "Exemplo", "Idade": "29"}},
	})
}