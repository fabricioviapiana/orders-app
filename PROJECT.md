# 📦 PROJECT.md

## 🎯 Objetivo do projeto

Este é um projeto de backend desenvolvido em Go com o objetivo de:

- aprender a linguagem de forma profunda
- evoluir até um nível de desenvolvimento sênior
- praticar construção de sistemas reais
- entender arquitetura de backend na prática

O foco não é velocidade, e sim:

- clareza
- simplicidade
- boas decisões técnicas
- evolução incremental

---

## 🧱 Arquitetura atual

O projeto segue uma arquitetura em camadas:

- **handler** → camada HTTP (entrada/saída)
- **service** → regras de negócio
- **repository** → acesso a dados (em memória)
- **domain** → entidades do sistema

### Fluxo de execução

Request → Handler → Service → Repository → Service → Handler → Response

---

## 📁 Estrutura do projeto

```text
cmd/api            → ponto de entrada da aplicação
internal/domain    → entidades
internal/handler   → handlers HTTP
internal/service   → lógica de negócio
internal/repository→ persistência em memória
```

---

## 📊 Domínio atual

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

## 📌 Regras de negócio atuais

### Produtos

- nome não pode ser vazio
- preço deve ser maior que zero

### Usuários

- nome não pode ser vazio
- email não pode ser vazio

### Pedidos

Ao criar um pedido:

- usuário deve existir
- deve haver pelo menos um item
- quantidade deve ser maior que zero
- produto deve existir
- preço do produto deve ser copiado para `UnitPrice`
- total deve ser calculado com base nos itens

---

## 🔌 Endpoints atuais

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

## 🧪 Estado atual

- API HTTP funcional com `net/http`
- persistência em memória
- separação em camadas implementada
- validações básicas presentes
- domínio inicial definido

---

## 🚧 Próximos passos (curto prazo)

- garantir consistência nos endpoints
- melhorar validações no service
- tratar erros de forma mais consistente
- adicionar testes unitários para services
- evitar duplicação de código

---

## 🧠 Decisões de projeto

- usar apenas standard library (sem frameworks)
- manter arquitetura simples e explícita
- evitar abstrações prematuras
- priorizar entendimento sobre performance inicial
- evoluir o sistema em fases

---

## 📈 Evolução planejada

O projeto será evoluído em fases:

- Fase 1 → base HTTP + CRUD simples
- Fase 1.5 → relações + validações + testes
- Fase 2 → banco de dados + persistência real
- Fase 3 → autenticação e autorização
- Fase 4 → concorrência e processamento assíncrono
- Fase 5 → observabilidade e produção

---

## 🎯 Critério de qualidade

Uma funcionalidade só é considerada concluída quando:

- funciona corretamente
- possui validações adequadas
- está no lugar certo da arquitetura
- não introduz complexidade desnecessária
- é compreensível ao ler o código

---

## 🧭 Filosofia do projeto

Este projeto é um ambiente de aprendizado intencional.

Decisões devem priorizar:

- aprendizado de fundamentos
- código simples e claro
- evolução consistente
- entendimento profundo

Velocidade é secundária.
