--- Metadados da Feature:
    - Nome: Melhorias em API plataforma GetFund
    - Código: FEAT01
    - Status: Review
    - Message: N/A

Recursos:
    - STATUS: Draft|ToDo|Spec|Proj|Dev|Blocked|Review|Done
    - Prefixo Feature: FEAT{0-N}
    - Prefixo Requisito Funcional: [STATUS] RF{0-N}
    - Prefixo Requisito Técnico: [RF{0-N}] RT{0-N}
    - Prefixo Task: [RT{0-N}] TASK{0-N}
    - Prefixo Task de Teste: [TASK{0-N}] TEST{0-N}
---

# 1. API para acesso a plataforma GetFund

## 2. Descrição
**Resumo:** Esta feature abrange a funcionalidade completa para melhorias na API da plataforma GetFund.

## 3. Objetivos e Métricas de Sucesso
**Objetivos de Negócio:**
- Evoluir as funcionalidades da API.
- Reduzir a complexidade de código.
- Reduzir o fluxo aplicacional.

**Métricas de Sucesso (KPIs):**
- Rapidez no acessar a plataforma.
- Segurança de dados.
- Código limpo, robusto, enxuto, escalável e técnico.

## 4. Requisitos Funcionais
- **[Todo] RF01:** Melhorar o processo de comunicação dos eventos.
- **[Draft] RF02:** ...

### 4.1. Diagramas Caso de Uso

N/A

## 5. Requisitos Técnicos
- **[RF01] RT01:** Olhar para todos os dominios e mover os eventos e seus respectivos payload para o local designado no padrão deste projeto.
- **[RF01] RT02:** Padronizar os nomes dos eventos conforme padrão do projeto.
- **[RF01] RT03:** Olhar para os testes existentes de cada dominios e garantir que passem.
- **[RF01] RT04:** Resolver os imports e verificar se há inconsitências e erros nos nomes dos testes.

- **Hipóteses:**
    - Será apenas transferência de arquivos.

- **Riscos:**
    - Inconsistêcia nos teste e falha na comunicação dos eventos.

### 5.1. Diagrama de Fluxo

N/A

## 6. Modelagem Arquitetura

### 6.1. Diagrama de Arquitetura

N/A

### 6.2. Diagrama de Evento

N/A

## 7. Tarefas

- [x] **[RF01]TASK01:** Mover `SignupStartedEvent` e seu payload para `internal/domain/auth/core/usecase/signup/event/signup_started_event_payload.go` e padronizar o nome do evento no bus para `signup.started`.
- [x] **[RF01][TASK01]TEST01:** Garantir que os testes de `signup_test.go` e `signup_started_event_handler_test.go` passem após a refatoração.
- [x] **[RF01]TASK02:** Mover `ActivateUserWithCouponConfirmedEvent` e seu payload para `internal/domain/auth/core/usecase/activate_user/event/activation_confirmed_event_payload.go` e padronizar o nome do evento no bus para `activate.user.with.coupon.confirmed`.
- [x] **[RF01][TASK02]TEST01:** Garantir que os testes de `activate_user_test.go` e `activate_user_with_coupon_confirmed_event_handler_test.go` passem após a refatoração.
- [x] **[RF01]TASK03:** Mover `RecoverPasswordStartedEvent` e seu payload para `internal/domain/auth/core/usecase/recover_password/event/recover_password_started_event_payload.go` e padronizar o nome do evento no bus para `recover.password.started`.
- [x] **[RF01][TASK03]TEST01:** Garantir que os testes de `recover_password_test.go` e `recover_password_started_event_handler_test.go` passem após a refatoração.
- [x] **[RF01]TASK04:** Mover `ValidatePrizeDrawCouponStartedEvent` e seu payload para `internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon/event/validation_started_event_payload.go` e padronizar o nome do evento no bus para `validate.prizedraw.coupon.started`.
- [x] **[RF01][TASK04]TEST01:** Garantir que os testes de `validate_prizedraw_coupon_test.go` e `validate_prizedraw_coupon_started_event_handler_test.go` passem após a refatoração.
- [x] **[RF01]TASK05:** Mover `ApplyPrizeDrawCouponStartedEvent` e seu payload para `internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon/event/apply_coupon_started_event_payload.go` e padronizar o nome do evento no bus para `apply.prizedraw.coupon.started`.
- [x] **[RF01][TASK05]TEST01:** Garantir que os testes de `apply_prizedraw_coupon_test.go` passem após a refatoração.
- [x] **[RF01]TASK06:** Mover `ApplyPrizeDrawCouponFailedEvent` e seu payload para `internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon/event/apply_coupon_failed_event_payload.go` e padronizar o nome do evento no bus para `apply.prizedraw.coupon.failed`.
- [x] **[RF01][TASK06]TEST01:** Garantir que os testes de `apply_prizedraw_coupon_test.go` passem após a refatoração.
- [x] **[RF01]TASK07:** Deletar os arquivos `internal/domain/auth/core/event/events.go` e `internal/domain/prizedraw/core/event/events.go`.
- [x] **[RF01][TASK07]TEST01:** Garantir que o projeto compila e todos os testes passam após a exclusão dos arquivos antigos.

### 7.1 Diagram de Sequência
