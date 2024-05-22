## 🎉 Sistema de Gestão de Estoque - API RESTful 🎉

```
 ⚠️ Esse projeto é apenas um exemplo e não deve ser utilizado em produção
     sem as devidas modificações e testes.
```

Este repositório contém o código fonte para uma API RESTful que está sendo desenvolvida em Go para gerenciar o estoque de uma loja online. 🚀

### **Funcionalidades Principais:**

- 📦 **Gerenciamento de Produtos:**

  - Cadastro de produtos, incluindo nome, descrição, preço, quantidade em estoque, imagem, categoria e marca.
  - 🔄 Atualização de informações de produtos existentes.
  - 🗑️ Remoção de produtos.
  - 🔍 Busca e filtragem de produtos por critérios como nome, categoria, marca, preço, etc.

- 📦 **Gerenciamento de Estoque:**

  - Cadastro de estoques, incluindo nome, localização e capacidade.
  - ➕ Adição e remoção de produtos de estoques.
  - 👀 Visualização da quantidade de produtos em cada estoque.

- 🛒 **Gerenciamento de Pedidos:**

  - Cadastro de pedidos, incluindo cliente, produtos, data do pedido, status do pedido, valor total, etc.
  - 🔎 Visualização de pedidos por cliente, status, data, etc.
  - 🔄 Atualização do status do pedido.
  - ❌ Cancelamento de pedidos.

- 👤 **Gerenciamento de Clientes:**

  - Cadastro de clientes, incluindo nome, email, senha, endereço e telefone.
  - 🔐 Autenticação de clientes via JWT (JSON Web Token).
  - 👀 Visualização de informações do cliente.
  - 🔄 Atualização de informações do cliente.

- 💳 **Gerenciamento de Pagamentos:**

  - Processamento de pagamentos via API de pagamento (integração não implementada neste exemplo).
  - 📝 Registro de informações de pagamento, incluindo tipo, data, valor e status.

- 🚚 **Gerenciamento de Entregas:**

  - 📝 Registro de informações de entrega, incluindo tipo, data, status e transportadora.
  - 🗺️ Rastreamento de entregas.

- 🎁 **Gerenciamento de Cupons:**

  - Criação de cupons promocionais, incluindo código, desconto e data de validade.
  - 🎟️ Aplicação de cupons a pedidos.

- 💬 **Feedback de Clientes:**
  - 🗣️ Coleta de feedback de clientes sobre pedidos, incluindo avaliação e comentários.

### 🛠️ **Tecnologias:**

- **Go:** Linguagem de programação utilizada para desenvolver a API.
- **PostgreSQL:** Banco de dados relacional usado para armazenar os dados do sistema.
- **GORM:** ORM (Object-Relational Mapping) para facilitar a interação com o banco de dados.
- **Gin:** Framework web para construir a API RESTful.
- **JWT:** JSON Web Token para autenticação de usuários.

### ⚙️ **Instalação e Execução:**

1. **Clone o repositório:** `gh repo clone becardine/gestock-api`
2. **Configure as variáveis de ambiente:** Crie um arquivo `.env` na raiz do projeto e configure as variáveis de ambiente necessárias (consulte o arquivo `.env.example`).
3. **Instale as dependências:** `go mod tidy`
4. **Execute a aplicação:** `go run cmd/server/main.go`

### 📑 **Documentação:**

A documentação da API está disponível em formato OpenAPI/Swagger (acessível via `/docs/index.html` após o servidor iniciar).

### **Contribuições:**

1. Abra um issue para reportar um problema ou sugerir uma nova funcionalidade.
2. Faça um fork do repositório.
3. Implemente sua alteração.
4. Abra um Pull Request para enviar sua contribuição.

### **Licença:**

Este projeto está licenciado sob a licença MIT.
