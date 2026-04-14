# 📦 PROJECT_STATE.md

## 🎯 Objetivo do projeto

Backend em Go com foco em aprendizado profundo e evolução até nível sênior.

O projeto evolui por fases, aumentando gradualmente a complexidade.

---

## 🧱 Arquitetura atual

Arquitetura em camadas:

- handler → HTTP (entrada/saída)
- service → regras de negócio
- repository → persistência (em memória)
- domain → entidades

### Fluxo

Request → Handler → Service → Repository → Service → Handler → Response

---

## 📁 Estrutura

cmd/api → entrypoint
internal/domain → entidades
internal/handler → HTTP
internal/service → lógica
internal/repository → armazenamento

---

## 📊 Domínio implementado

### User

- ID
- Name
- Email

### Product

- ID
- Name
- Price

### Order

- ID
- UserID
- Items []OrderItem
- Total

### OrderItem

- ProductID
- Quantity
- UnitPrice

---

## 🔌 Endpoints implementados

### Products

- GET /products
- POST /products
- GET /products/{id}
- DELETE /products/{id}

### Users

- GET /users
- POST /users
- GET /users/{id}
- DELETE /users/{id}

### Orders

- GET /orders
- POST /orders
- GET /orders/{id}

---

## 📌 Regras de negócio implementadas

### Products

- nome obrigatório
- preço > 0

### Users

- nome obrigatório
- email obrigatório

### Orders

- usuário deve existir
- pelo menos 1 item
- quantidade > 0
- produto deve existir
- preço copiado para UnitPrice
- total calculado

---

## 🧪 Estado atual

- API funcional com net/http
- Persistência em transição (em memória -> SQL)
- **Interfaces de Repositório atualizadas para retornar `error`**
- **Services e Handlers refatorados para propagar erros adequadamente**
- **Auto-Migrations configuradas com Goose (PostgreSQL)**
- Docker configurado (Postgres 16)
- Infraestrutura de conexão ao banco integrada ao `main.go`
- Fase 2 em progresso (Infra pronta, falta implementar os Repositories SQL)

---

## ⚠️ Limitações atuais

- sem banco de dados
- sem autenticação
- sem testes avançados
- sem concorrência
- sem observabilidade

---

## 🧠 Decisões importantes

- usar apenas standard library
- manter simplicidade
- evitar abstrações prematuras
- foco em aprendizado

---

## 🎯 Próximo objetivo

Evoluir para persistência real (banco de dados) mantendo:

- arquitetura atual
- separação de responsabilidades
- simplicidade do código
