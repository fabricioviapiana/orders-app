# 🗺️ ROADMAP.md

## 🎯 Objetivo

Guiar a evolução do projeto de forma incremental até um nível avançado de backend em Go.

---

## ✅ Fase 1 — Base (concluída)

- servidor HTTP com net/http
- CRUD básico
- estrutura de projeto
- separação em camadas

---

## ✅ Fase 1.5 — Domínio e validação (concluída)

- relações entre entidades
- regras de negócio
- orders com itens
- validações consistentes
- endpoints completos

---

## 🚧 Fase 2 — Persistência (em progresso)

Objetivo:

- substituir armazenamento em memória por banco real

Status:

- [x] Configurar Docker/Postgres
- [x] Refatorar interfaces para suportar `error`
- [x] Implementar Auto-Migrations (Goose)
- [x] Criar PostgresUserRepository
- [x] Criar PostgresProductRepository
- [x] Criar PostgresOrderRepository
- [ ] Tratamento de erros de banco

---

## 🔜 Fase 2.5 — Refinamento

- melhorar estrutura de erro
- melhorar respostas HTTP
- padronizar retornos
- refatorações leves
- melhorar testes

---

## 🔜 Fase 3 — Autenticação

- login
- hash de senha
- JWT
- middleware de autenticação
- autorização básica

---

## 🔜 Fase 4 — Concorrência e processamento assíncrono

- goroutines
- channels
- workers
- jobs assíncronos
- processamento em background

---

## 🔜 Fase 5 — Infra e observabilidade

- logs estruturados
- métricas
- tracing
- health checks

---

## 🔜 Fase 6 — Arquitetura avançada

- separação em serviços
- comunicação entre serviços
- eventos
- consistência eventual

---

## 🧭 Diretriz geral

Sempre evoluir:

- sem quebrar o que já existe
- mantendo simplicidade
- entendendo cada etapa profundamente

---

## 🎯 Regra principal

Não avançar para a próxima fase sem entender completamente a atual.
