package main

import (
	"flag"
	"fmt"
	"os"
)

type Pagamento struct {
	Id           int64
	ChaveDestino string
	Valor        float64
}

type QRPayload struct {
	TransacaoId int64 `json:"transacao_id"`
}

func main() {
	chave := flag.String("chave", "", "Chave de destino do pagamento (ex: email, cpf )")
	valor := flag.Float64("valor", 0.0, "Valor do pagamento (ex: 12.34)")
	flag.Parse()

	if *chave == "" || *valor <= 0 {
		fmt.Println("Erro: A chave e um valor positivo são obrigatórios.")
		fmt.Println("Uso: go run main.go --chave=\"sua-chave\" --valor=50.25")
		os.Exit(1)
	}

	fmt.Printf("Dados recebidos: Chave =[%s], Valor=[%.2f]\n", *chave, *valor)
}
