# Gerador de QR Code de Pagamento em Go

Este projeto é uma CLI em Go que simula a criação de um QR Code de pagamento. Ele recebe os dados de um pagamento, registra a transação em um banco de dados PostgreSQL e gera uma imagem de QR Code contendo o ID único da transação.

## Pré-requisitos
- Go (1.18+)
- Docker e Docker Compose

## Como Rodar
1. Inicie o banco de dados PostgreSQL com Docker:

   docker run --name meu-postgres -e POSTGRES_PASSWORD=minhasenha -e POSTGRES_USER=meuuser -e POSTGRES_DB=pagamentos -p 5432:5432 -d postgres`

1. Instale as dependências Go:Bash
    
    `go mod tidy`
    
2. Execute o programa:Bash
    
    `go run main.go --chave="seu-email@provedor.com" --valor=19.99`
    

Isso irá criar um arquivo `pagamento.png` no diretório do projeto.
