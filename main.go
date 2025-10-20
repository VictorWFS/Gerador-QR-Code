package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Pagamento struct {
	Id           int64
	ChaveDestino string
	Valor        float64
}

type QRPayload struct {
	TransacaoId int64 `json:"transacao_id"`
}

func initDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	statement := `
		CREATE TABLE IF NOT EXISTS pagamentos (
			id SERIAL PRIMARY KEY,
			chave TEXT NOT NULL,
			valor NUMERIC(10,2) NOT NULL
		);
	`
	_, err = db.Exec(statement)
	if err != nil {
		return nil, err
	}
	fmt.Println("Conexão com PostgreeSQL estabelecida e tabela pronta.")
	return db, nil
}

func salvarPagamento(db *sql.DB, p Pagamento) (int64, error) {
	var id int64
	query := "INSERT INTO pagamentos (chave, valor) VALUES ($1, $2) RETURNING id"
	err := db.QueryRow(query, p.ChaveDestino, p.Valor).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func main() {
	connStr := "postgres://postgres:1234@localhost:5432/pagamentos?sslmode=disable"
	db, err := initDB(connStr)
	if err != nil {
		fmt.Println("Erro ao inicializar o banco de dados: ", err)
		os.Exit(1)
	}
	defer db.Close()

	chave := flag.String("chave", "", "Chave de destino do pagamento (ex: email, cpf )")
	valor := flag.Float64("valor", 0.0, "Valor do pagamento (ex: 12.34)")
	flag.Parse()

	if *chave == "" || *valor <= 0 {
		fmt.Println("Erro: A chave e um valor positivo são obrigatórios.")
		fmt.Println("Uso: go run main.go --chave=\"sua-chave\" --valor=50.25")
		os.Exit(1)
	}

	novoPagamento := Pagamento{
		ChaveDestino: *chave,
		Valor:        *valor,
	}

	id, err := salvarPagamento(db, novoPagamento)
	if err != nil {
		fmt.Println("Erro ao salvar pagamento no banco: ", err)
		os.Exit(1)
	}
	fmt.Printf("Pagamento registrado no banco com ID: %d\n", id)
	fmt.Printf("Dados recebidos: Chave =[%s], Valor=[%.2f]\n", *chave, *valor)
}
