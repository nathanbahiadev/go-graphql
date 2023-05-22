# Go GraphQL
Este projeto é uma aplicação Go que implementa um servidor GraphQL. Ele permite que você execute consultas e mutações em um playground GraphQL.

## Instruções de execução
Para executar o projeto, siga as etapas abaixo:

1. Clone o repositório para o seu ambiente local.
2. Execute o comando go mod tidy para baixar todas as dependências necessárias.
3. Em seguida, execute o comando go run cmd/server/server.go para iniciar o servidor.
4. O servidor irá iniciar o banco de dados e a aplicação, e um playground GraphQL estará disponível na porta 8080.

Certifique-se de ter o Go instalado em seu sistema antes de executar o projeto.

## Endpoints
- Playground GraphQL: http://localhost:8080