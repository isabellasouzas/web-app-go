package main

import (
	"net/http"

	"github/isabellasouzas/web-app-go/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
