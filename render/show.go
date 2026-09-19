package render

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type PageContext struct {
	Title string      // Título da aba do navegador
	Data any // Dados específicos da página (Lista, Formulários, etc.)
}

// Renderiza qualquer página aplicando o layout base e os partials (como o menu)
func Show(w http.ResponseWriter, status int, pageFile string, data any) {
	// Lista de arquivos base que TODA página precisa (layout + menu)
	files := []string{
		"layout/layout.html",
		// "layout/partials/menu.html",
		filepath.Join("layout", pageFile), // Ex: "app/cadastro/lista.html"
	}

	// Faz o parse dos arquivos necessários para essa requisição
	tpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Erro ao carregar templates: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	// Executa a estrutura base
	err = tpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, "Erro ao renderizar template: "+err.Error(), http.StatusInternalServerError)
	}
}