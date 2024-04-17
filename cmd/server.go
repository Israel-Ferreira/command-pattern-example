package cmd

import (
	"log"
	"net/http"

	"github.com/Israel-Ferreira/pre-order/preorder-service/internal/handlers"
)

func StartMuxServer() {
	log.Println("Configurando e Iniciando o Servidor...")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /reserve-product", handlers.ReservePreOrderProductHandler)

	mux.HandleFunc("DELETE /reserve/{reserveId}", handlers.CancelReserveHandler)

	log.Fatalln(http.ListenAndServe(":8080", mux))
}
