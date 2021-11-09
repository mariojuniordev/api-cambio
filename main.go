package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Carregar()
	//Encapsulamos um router dentro de uma variável qualquer "r" com inferência de tipo.
	r := router.Gerar()

	fmt.Printf("Rodando API - SISTEMA DE CÂMBIO LABSIT na porta %d", config.Porta)

	//http.ListenAndServe() sobe o servidor em uma porta desejada.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}