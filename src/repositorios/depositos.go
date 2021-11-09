package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Depositos uma structure que representa um repositório de depósitos.
type Depositos struct {
	db *sql.DB
}

// NovoRepositorioDeDepositos cria um depositório de depósitos.
func NovoRepositorioDeDepositos(db *sql.DB) *Depositos {
	return &Depositos{db}
}

// Criar insere um depósito no banco de dados.
func (repositorio Depositos) Criar(deposito modelos.Deposito) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO depositos (valorDeposito) VALUES (?)")
	if erro != nil {
		return 0, nil
	}
	defer statement.Close()
	resultado, erro := statement.Exec(deposito.ValorDeposito)
	if erro != nil {
		return 0, nil
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, nil
	}
	return uint64(ultimoIDInserido), nil
}

//BuscarDepositos mostra todos os depósitos feitos pelo usuário.
func (repositorio Depositos) BuscarDepositos() ([]modelos.Deposito, error) {
	linhas, erro := repositorio.db.Query("SELECT valorDeposito FROM depositos")
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	// Inicializamos um slice chamado "depositos" para armazenar os dados de cada depósito feito.
	var depositos []modelos.Deposito

	// O loop for (while) abaixo irá exibir todos os valorDeposito presentes na tabela depositos.
	for linhas.Next() {
		var deposito modelos.Deposito
		if erro = linhas.Scan(&deposito.ValorDeposito); erro != nil {
			return nil, erro
		}
		// Esse trecho de código fará a inserção de cada novo depósito no slice à cada interação do loop for.
		depositos = append(depositos, deposito) // Syntax append(<nome do slice>, <nome da variavel inserida>)
	}
	return depositos, nil
}

// BuscarSaldoTotal retorna a soma de todos os depósitos feitos e os exibe na moeda desejada.
func (repositorio Depositos) BuscarSaldoTotal(moeda string) (modelos.Saldo, error) {
	linhas, erro := repositorio.db.Query("SELECT valorDeposito FROM depositos")
	if erro != nil {
		return modelos.Saldo{}, erro
	}
	defer linhas.Close()
	valoresEmSlice := []float64{}
	for linhas.Next() {
		var deposito modelos.Deposito
		if erro = linhas.Scan(&deposito.ValorDeposito); erro != nil {
			return modelos.Saldo{}, erro
		}
		valoresEmSlice = append(valoresEmSlice, deposito.ValorDeposito)
	}
	var valorTotalEmBRL float64
	for i := 0; i < len(valoresEmSlice); i++ {
		valorTotalEmBRL += valoresEmSlice[i]
	}

	// Taxas e Impostos (Valores Absolutos).
	iof := 1 - 0.0638 // 6.38%
	spreadBancario := 1 - 0.04 // 4%
	
	totalTaxas := iof*spreadBancario

	// Cálculo de Conversão de moeda.
	var valores modelos.Saldo

	switch moeda {
		case "USD":
			valores.ValorMoedaUSD = valorTotalEmBRL *  0.18 * totalTaxas
		case "EUR":
			valores.ValorMoedaEUR = valorTotalEmBRL * 0.16 * totalTaxas
		case "GBP":
			valores.ValorMoedaGBP = valorTotalEmBRL * 0.13 * totalTaxas
		default: 
			valores.ValorTotal = valorTotalEmBRL
		}
	return valores, nil
}