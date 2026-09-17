package main

import (
	"fmt"
	"net/http"

	"github.com/profmugomes/mgcadastro/cadastro"

	// "strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	//http.ServeFile(w, r, "app/index.html")

// 	tpl := template.Must(template.ParseFiles("app/index.html"))
// 	_ = tpl.Execute(w, PageData{Title: "Home Page", User: "Teste"})
// }

// func trailingSlash(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.URL.Path != "/" && !strings.HasSuffix(r.URL.Path, "/") {
// 			http.Redirect(w, r, r.URL.Path+"/", http.StatusPermanentRedirect)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", homeHandler)
	// mux.HandleFunc()

	rt := chi.NewRouter()

	rt.Use(middleware.RedirectSlashes)
	// rt.Use(trailingSlash)

	rt.Get("/", cadastro.Lista)
	rt.Get("/editar", cadastro.Editar)

	// rt.Route("/categoria", func(rt chi.Router) {
	// 	rt.Get("/teste", func(w http.ResponseWriter, r *http.Request) {
	// 		w.Write([]byte("Teste"))
	// 	})
	// })

  	fmt.Println("Servidor iniciado em 127.0.0.1:8000")

	err := http.ListenAndServe(":8000", rt)
	if (err != nil) {
		fmt.Println("Erro ao iniciar o servidor: ", err)
		return
	}
}