# User API

User API é uma aplicação Back-end desenvolvida em Go (Golang) utilizando Gin, PostgreSQL, GORM, JWT Authentication e Swagger para documentação.

O projeto implementa autenticação baseada em token JWT, arquitetura em camadas (Handler → Service → Repository) e organização modular utilizando internal/.


### Instalação/Execução:
```bash

# Instalar dependências
go mod tidy

# Inicie o servidor (localhost)
go run ./cmd/api

```

### Estrutura:  

```text
cmd/
└── api/
    └── main.go

internal/
├── database/
│   └── postgres.go
├── middleware/
│   └── auth.go
├── router/
│   └── router.go
├── user/
│   ├── model.go
│   ├── dto.go
│   ├── repository.go
│   ├── service.go
│   └── handler.go
├── utils/
│   └── jwt.go

docs/ (gerado pelo Swagger)

go.mod
.env
docker-compose.yml

```

### Endpoints:
<img width="1457" height="507" alt="image" src="https://github.com/user-attachments/assets/9b33ceca-e512-441b-8669-065267e0cf7c" />




