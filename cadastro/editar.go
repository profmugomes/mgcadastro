package cadastro

import (
	"net/http"

	"github.com/profmugomes/mgrender"
)

func Editar(ctx *mgrender.Context, w http.ResponseWriter, r *http.Request) {
	ctx.AddFile("layout/editar.html")

	ctx.SetData("Title", "Lista de Cadastro")

	ctx.SetData("Outro", []string{
		"Teste 1",
		"Teste 2",
	})

	ctx.SetData("Data", []map[string]string{
		{
			"Titulo": "Exemplo",
			"Idade":  "29",
		},
	})
}
