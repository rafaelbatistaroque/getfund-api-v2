--- Metadados da Feature:
    - Nome: Melhorias em API plataforma GetFund
    - Código: FEAT02
    - Status: Done
    - Message: A refatoração foi concluída com sucesso. Todos os testes estão passando.

Recursos:
    - STATUS: Draft|ToDo|Spec|Proj|Dev|Blocked|Review|Done
    - Prefixo Feature: FEAT{0-N}
    - Prefixo Requisito Funcional: [STATUS] RF{0-N}
    - Prefixo Requisito Técnico: [RF{0-N}] RT{0-N}
    - Prefixo Task: [RT{0-N}] TASK{0-N}
    - Prefixo Task de Teste: [TASK{0-N}] TEST{0-N}
---

# 1. API para acesso a plataforma GetFund
NOTA: Não é necessários criar diagramas para esta feature.

## 2. Descrição
**Resumo:** Esta feature abrange a funcionalidade completa para melhorias dos testes na API da plataforma GetFund.

## 3. Objetivos e Métricas de Sucesso
**Objetivos de Negócio:**
- Abranger a cobertura de testes da API.

**Métricas de Sucesso (KPIs):**
- Mitigar erros.
- Segurança das operações.
- Código limpo, robusto, enxuto, escalável e técnico.

## 4. Requisitos Funcionais
- **[Done] RF01:** Criar testes para o serviço compartilhado do pacote security.
- **[Done] RF02:** Organizar, neste documento, as referências das tarefas confome critérios da documentação.
- **[Done] RF03:** Criar teste para reposítório de autenticação
- **[Done] RF04:** Criar testes para o serviço compartilhado do pacote email
- **[Done] RF05:** Refatorar fixture que contenham método `GetHttpRequestResponse` 

### 4.1. Diagramas Caso de Uso

N/A

## 5. Requisitos Técnicos
- **[RF01] RT02:** Criar teste para TODOS os métodos público do pacote security obedecendo os padões estabelecidos na documentação.
- **[RF01] RT03:** Continuar implemetação de fallback do método `GetAuthenticatedUserByUsername`.
- **[RF01] RT04:** Garantir teste dos novos métodos do repositório `authRepositoryProxy`.
- **[RF02] RT01:** Olhar para cada requisito técnico e identificar qual tarefa está relacionada a ele e alterar o prefixo da tarefa para o RT correspondente.
- **[RF03] RT01:** Garantir testes do `UpdateUsernameHash` do repositório `authRepository`.
- **[RF04] RT01:** Refatorar/Criar/Limpar testes que estão em `test/internal/shared/mail`.
- **[RF05] RT01:** Criar um arquivo base dentro de `test/helper`que possa receber itens de reuso.
- **[RF05] RT02:** Refactorar `GetHttpRequestResponse` nas fixtures de `test/internal` para ser reutilizado.

- **Hipóteses:**
    - N/A.

- **Riscos:**
    - N/A.

### 5.1. Diagrama de Fluxo

N/A

## 6. Modelagem Arquitetura

### 6.1. Diagrama de Arquitetura

N/A

### 6.2. Diagrama de Evento

N/A

## 7. Tarefas

- [X] **[RF01]**
    - [X] **[RT02]TASK01:** Criar o arquivo de teste `test/internal/shared/security/security_test.go`.
        - [X] **[TASK01]TEST01:** Criar os testes para o método `GetRandomCode`.
        - [X] **[TASK01]TEST02:** Criar os testes para o método `Encrypt`.
        - [X] **[TASK01]TEST03:** Criar os testes para o método `Decrypt`.
        - [X] **[TASK01]TEST04:** Criar os testes para o método `HashAndMerge`.
        - [X] **[TASK01]TEST05:** Criar os testes para o método `IsMatch`.
        - [X] **[TASK01]TEST06:** Criar os testes para o método `HashWithSalt`.
        - [X] **[TASK01]TEST07:** Criar os testes para o método `Hash`.
    - [X] **[RT03]TASK02:** Implementar fallback para o método `GetAuthenticatedUserByUsername` no `authRepositoryProxy`.
        - [X] **[TASK02]TEST01:** Criar teste para o fallback do método `GetAuthenticatedUserByUsername`.
        - [X] **[TASK02]TEST02:** Criar teste para a chamada do método `HashWithSaltLegacy` no fallback.
        - [X] **[TASK02]TEST03:** Criar teste para os parâmetros do método `HashWithSaltLegacy` no fallback.
        - [X] **[TASK02]TEST04:** Criar teste para o erro de retorno do método `HashWithSaltLegacy` no fallback.
        - [X] **[TASK02]TEST05:** Criar teste para a chamada do método `GetAuthenticatedUserByUsername` no fallback.
        - [X] **[TASK02]TEST06:** Criar teste para os parâmetros do método `GetAuthenticatedUserByUsername` no fallback.
        - [X] **[TASK02]TEST07:** Criar teste para o erro de retorno do método `GetAuthenticatedUserByUsername` no fallback.
- [X] **[RF03]**
    - [X] **[RT01]TASK03:** Criar os testes para o método `UpdateUsernameHash` em `test/internal/domain/auth/adapter/repository/auth_repository_test.go`.
        - [X] **[TASK03]TEST01:** Criar teste para o sucesso na atualização do username.
        - [X] **[TASK03]TEST02:** Criar teste para o erro na atualização do username.
        - [X] **[TASK03]TEST03:** Criar teste para a atualização de um usuário inexistente.
- [X] **[RF04]**
    - [X] **[RT01]TASK04:** Refatorar/Criar/Limpar testes que estão em `test/internal/shared/mail`.
        - [X] **[TASK04]TEST01:** Criar teste para o sucesso no envio do email.
        - [X] **[TASK04]TEST02:** Criar teste para o erro no envio do email.
- [X] **[RF05]**
    - [X] **[RT01]TASK05:** Criar o arquivo `test/helper/fixture/fixture.go`.
        - [X] **[TASK05]TEST01:** Criar a função `GetHttpRequestResponse` no arquivo `test/helper/fixture/fixture.go`.
    - [X] **[RT02]TASK06:** Refatorar as fixtures para usar a nova função `GetHttpRequestResponse`.
        - [X] **[TASK06]TEST01:** Refatorar `test/internal/domain/auth/adapter/gateway/activate_user_gateway_fixture/activate_user_gateway_fixture.go`.
        - [X] **[TASK06]TEST02:** Refatorar `test/internal/domain/auth/adapter/gateway/recover_password_gateway_fixture/recover_password_gateway_fixture.go`.
        - [X] **[TASK06]TEST03:** Refatorar `test/internal/domain/auth/adapter/gateway/reset_password_gateway_fixture/reset_password_gateway_fixture.go`.
        - [X] **[TASK06]TEST04:** Refatorar `test/internal/domain/auth/adapter/gateway/signin_gateway_fixture/signin_gateway_fixture.go`.
        - [X] **[TASK06]TEST05:** Refatorar `test/internal/domain/auth/adapter/gateway/signout_gateway_fixture/signout_gateway_fixture.go`.
        - [X] **[TASK06]TEST06:** Refatorar `test/internal/domain/auth/adapter/gateway/signup_gateway_fixture/signup_gateway_fixture.go`.

### 7.1 Diagrama de Sequência

N/A