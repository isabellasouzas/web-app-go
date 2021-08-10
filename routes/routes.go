package routes

import (
	"net/http"

	"github/isabellasouzas/web-app-go/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
