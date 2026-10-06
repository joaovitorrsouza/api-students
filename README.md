# API Students — Go

API REST para gerenciamento de estudantes, desenvolvida em **Go** para praticar conceitos de Back-End, organização em camadas, persistência de dados e tratamento de requisições HTTP.

## Stack

- Go 1.22
- Echo
- GORM
- SQLite
- Zerolog

## Funcionalidades

- criar estudante
- listar estudantes
- buscar estudante por ID
- atualizar estudante
- remover estudante
- filtrar estudantes por status ativo/inativo
- validar campos obrigatórios de entrada
- retornar respostas HTTP adequadas para diferentes cenários

## Endpoints

| Método | Endpoint | Descrição |
|---|---|---|
| GET | `/students` | Lista estudantes |
| GET | `/students?active=true` | Filtra estudantes ativos |
| GET | `/students?active=false` | Filtra estudantes inativos |
| GET | `/students/:id` | Busca um estudante |
| POST | `/students` | Cria um estudante |
| PUT | `/students/:id` | Atualiza um estudante |
| DELETE | `/students/:id` | Remove um estudante |

## Modelo de estudante

```json
{
  "name": "João",
  "cpf": "00000000000",
  "email": "joao@email.com",
  "age": 23,
  "active": true
}
```

## Organização do projeto

```text
.
├── api/
│   ├── api.go
│   ├── handler.go
│   └── request.go
├── db/
│   └── db.go
├── schemas/
│   └── schemas.go
├── main.go
├── go.mod
└── README.md
```

### API

Responsável pela configuração do servidor, rotas, handlers, validação das requisições e respostas HTTP.

### DB

Camada de acesso a dados utilizando GORM com SQLite.

### Schemas

Estruturas utilizadas para representação e resposta dos dados de estudantes.

## Middleware

A API utiliza os middlewares do Echo para:

- logging das requisições
- recuperação de panics

## Como executar

Com Go instalado:

```bash
go mod download
go run .
```

Por padrão, o servidor é iniciado na porta:

```text
http://localhost:8080
```

## O que este projeto demonstra

- desenvolvimento Back-End com Go
- criação de APIs REST
- CRUD
- validação de entrada
- tratamento de erros HTTP
- ORM com GORM
- persistência com SQLite
- separação de responsabilidades
- filtros via query parameters
