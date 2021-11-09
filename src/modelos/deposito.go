package modelos

import "errors"

// Deposito é a structure que contém o ID e o valor de cada depósito.
type Deposito struct {
	ID uint64 `json:"id,omitempty"`
	ValorDeposito float64 `json:"valorDeposito,omitempty"`
}

// Preparar prepara o depósito para ser enviado ao banco de dados.
func (deposito *Deposito) Preparar() error {
	if erro := deposito.validar(); erro != nil {
		return erro
	}
	return nil
}

//Validar valida o depósito caso tenha sido feito com sucesso.
func (deposito *Deposito) validar() error {
	if deposito.ValorDeposito <= 0 {
		return errors.New("o valor do depósito deve ser positivo e maior que zero")
	}
	return nil
}