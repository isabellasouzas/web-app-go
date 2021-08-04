package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{{Nome: "Camiseta", Descricao: "amarela", Preco: 25, Quantidade: 3},
		{"tenis", "corrida", 200, 2}}

	temp.ExecuteTemplate(w, "Index", produtos)
}
