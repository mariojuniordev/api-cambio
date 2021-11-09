package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Depositar deposita um valor em reais.
func Depositar(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var deposito modelos.Deposito 
	if erro = json.Unmarshal(corpoRequisicao, &deposito); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = deposito.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	repositorio := repositorios.NovoRepositorioDeDepositos(db)
	deposito.ID, erro = repositorio.Criar(deposito)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusCreated, deposito)
	fmt.Printf("Foi depositado o valor de %.2f \n", deposito.ValorDeposito)
}

// BuscarDepositos mostra todos os depósitos feitos.
func BuscarDepositos(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	repositorio := repositorios.NovoRepositorioDeDepositos(db)
	depositos, erro := repositorio.BuscarDepositos()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, depositos)
}

// BuscarSaldoTotal mostra o saldo total na moeda desejada.
func BuscarSaldoTotal(w http.ResponseWriter, r *http.Request) {
	moeda := r.URL.Query().Get("moeda")
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	repositorio := repositorios.NovoRepositorioDeDepositos(db)
	saldoTotal, erro := repositorio.BuscarSaldoTotal(moeda)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, saldoTotal)
}

