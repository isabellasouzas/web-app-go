package main

import (
	"net/http"
	"text/template"

	"github/isabellasouzas/web-app-go/db"
	"github/isabellasouzas/web-app-go/models"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}
