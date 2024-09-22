# Task Manager

Um sistema de gerenciamento de tarefas simples e eficiente desenvolvido em Go. Este projeto permite que os usuários se registrem, façam login e gerenciem suas tarefas de forma prática. Utiliza um banco de dados SQLite para persistência de dados.

## Funcionalidades

- Registro e autenticação de usuários
- Criação, leitura e gerenciamento de tarefas
- Persistência de dados usando SQLite
- Servidor HTTP para manipulação de requisições

## Estrutura do Projeto

- `main.go`: Ponto de entrada do aplicativo.
- `routes.go`: Definição das rotas e manipuladores.
- `models.go`: Modelos de dados para usuários e tarefas.
- `database.go`: Inicialização e configuração do banco de dados.
- `handlers.go`: Implementação das lógicas de negócio.
- `go.mod`: Gerenciamento de dependências.

## Como Executar

1. Clone o repositório.
2. Navegue até o diretório do projeto.
3. Execute `go mod tidy` para instalar as dependências.
4. Execute `go run main.go` para iniciar o servidor.

## Tecnologias Utilizadas

- Go
- SQLite

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.

## Licença

Este projeto está licenciado sob a MIT License.
