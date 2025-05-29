# 🏥 Sistema de Gestão Hospitalar - CESUPA

Este é um projeto de MVP (Produto Mínimo Viável) para modernizar o processo de gerenciamento de pacientes e consultas médicas no Hospital CESUPA. O sistema foi desenvolvido como parte de uma disciplina de Qualidade de Software, com foco em modularidade, testes e documentação.

---

## 🚀 Tecnologias Utilizadas

- **Go (Golang)** — Backend leve e performático
- **Gin** — Framework HTTP para Go
- **GORM** — ORM para Go (com SQLite)
- **SQLite** — Banco de dados leve e embutido
- **Swagger (swaggo)** — Documentação interativa da API
- **Testes com `testing` + `httptest`** — Cobertura de testes unitários e de integração
- **copier** — Cópia de structs para mapeamento DTO/modelo

---

## 📂 Estrutura do Projeto

```
.
├── config/               # Configuração da aplicação (variáveis, DB, etc.)
├── database/             # Inicialização do banco e seeds
├── handlers/             # Camada HTTP (controladores)
├── models/               # Modelos do banco de dados (GORM)
├── services/             # Camada de regras de negócio
├── middlewares/          # Middleware para autenticação e outros
├── docs/                 # Documentação gerada pelo Swagger
├── main.go               # Ponto de entrada da aplicação
```

---

## ✅ Casos de Uso Implementados

1. **Cadastrar paciente** (`POST /pacients`)
2. **Consultar paciente por ID** (`GET /pacients/{id}`)
3. **Atualizar paciente** (`PUT /pacients/{id}`)
4. **Inativar paciente** (`DELETE /pacients/{id}`)
5. **Agendar consulta** (`POST /pacients/{id}/appointments`)

---

## 🧪 Testes

Cada caso de uso tem:

- **Teste de unidade** com mocks
- **Teste de integração** usando banco real (SQLite)

Para rodar todos os testes:

```bash
go test ./...
```

---

## 🛠️ Como Rodar o Projeto

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/cesupa-hospital.git
cd cesupa-hospital
```

### 2. Instale as dependências

```bash
go mod tidy
```

### 3. Gere a documentação Swagger

```bash
swag init -g main.go -o docs
```

### 4. Execute a aplicação

```bash
go run main.go
```

### 5. Acesse o Swagger

Abra no navegador: [http://localhost:PORT/swagger/index.html](http://localhost:PORT/swagger/index.html)

---

## 📃 Documentação Swagger

A documentação completa da API está disponível via Swagger e descreve todos os endpoints com seus respectivos parâmetros, respostas e exemplos.

---

## 👥 Autores

- Cauã Maia de Souza Nara
- André Corrêa Sidrim
- Carlos Eduardo

---

## 📌 Observações

Este projeto é um MVP acadêmico. Algumas funcionalidades como autenticação, validação de agenda e histórico de edição podem ser expandidas em futuras versões.
