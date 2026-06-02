# meu-primeiro-crud-go

API REST com operações de CRUD de usuários, desenvolvida com Go como parte dos meus estudos em desenvolvimento backend.

## Objetivo

Este projeto foi criado para praticar a construção de uma API em Go com separação clara de responsabilidades entre controllers, models e rotas.

## Tecnologias

- **Go** — linguagem principal
- **Gin** — framework HTTP
- **MongoDB** — banco de dados
- **Zap** — logging estruturado
- **godotenv** — gerenciamento de variáveis de ambiente

## Funcionalidades

- Criar usuário
- Buscar usuário por ID
- Atualizar dados do usuário
- Deletar usuário

## Como executar

### Pré-requisitos

- Go 1.21+
- MongoDB em execução (local ou via Docker)

### Passos

1. Clone o repositório:
   ```bash
   git clone https://github.com/Guilherme-AVVO/meu-primeiro-crud-go.git
   cd meu-primeiro-crud-go
   ```

2. Crie um arquivo `.env` com as variáveis de ambiente necessárias.

3. Execute:
   ```bash
   go run main.go
   ```

A API ficará disponível em `http://localhost:8080`.

## Estrutura

```
src/
├── configuration/   # Logger e configurações globais
├── controller/      # Handlers HTTP e definição de rotas
└── model/           # Domínio e acesso ao banco de dados
```
