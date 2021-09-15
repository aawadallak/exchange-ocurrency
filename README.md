### Recursos utilizados

- Golang
- Teste Unitário
- PostgreSQL
- Docker
- DotEnv
- Domain-Drive-Design(DDD)

### Sobre o projeto

O módelo construido foi utilizando uma ideia de Domain-Drive-Design(DDD)
Onde separamos os componenetes da aplicação baseada em camadas.

Foi utilizado Golang como linguagem para a construção do Backend, para manutenção
do banco de dados optei por usar Docker para gerir o container do PostgreSQL, e
para dar manutenção ao código e pensando em desenvolvimento contínuo foi implementado
a uma biblíoteca de migrações de banco de dados, assim ao iniciar a aplicação ele checa
a última modíficação do banco, e caso haja uma versão nova, ele ira rodar para manter o
banco atualizado, e com suas dependências corretas.

Para a modelagem da RestAPI, a biblioteca Fiber foi escolhida, em virtude do benchmarking
de tempos de resposta e requisições em relação a outras bibliotecas do Golang, uma rota
de teste foi implementada para checar se a aplicação está rodando, denomidanada como uma
rota de "Ping", onde roda no root da aplicação.

Partindo da premissa que um código pode evoluir, a utilização do DDD é essencial,
uma arquitetura é essencial para escalabilidade e manutenibilidade do projeto.
Assim toda as camadas do projeto, tendem a seguir o seu conceito e a sua ídeia dentro do projeto.

Agrupando todas as ideias, foi escolhido utilizar variáveis de ambiente, a fim de um trabalho em equipe
e segurança da aplicação com uma maior tranquilidade. E afim de não deixar que as aplicações possuam falhas
foi implementado teste unitário nas funções que são essenciais em todo o fluxo do aplicação.

Atualmente o projeto implementou as seguintes interfaces do dominio para o banco de dados:

- Adicionar uma conversão ao banco de dados
- Encontrar uma conversão a partir de um ID
- Encontrar todas as conversões realizadas

## Rotas da aplicação

    Base: "http://localhost:5000"

# Rota de segurança
```
- ".../ping"  <br> ex: "http://localhost:5000/ping"
  ```

Rota responsável apenas por realizar checar a saúde da API, aferindo se a mesma se encontra funcionando.

# Grupo de rotas "exchange"
  - ".../{amount}/{from}/{to}/{rate}"  
    ex: "http://localhost:5000/exchange/10/BRL/USD/4.73"

Rota responsável por realizar a conversão do par de moedas e persistir no banco o registro da mesma.

- ".../all"
  ex: "http://localhost:5000/exchange/all"

Rota responsável por localizar todos as conversões já realizadas dentro da API.

- ".../{id}"
  ex: "http://localhost:5000/exchange/2"

Rota responsável por localizar no banco de dados um registro que possua o ID informado pelo usuário.

### Estrutura do projeto

```
┣ 📂 app
┃ ┗ 📂 http
┃ ┃ ┗ 📂 controllers
┣ 📂 cmd
┣ 📂 config
┣ 📂 database
┃ ┣ 📂 migrations
┣ 📂 domain
┃ ┗ 📂 ocurrency
┣ 📂 dto
┣ 📂 infra
┃ ┗ 📂 repository
┣ 📂 server
┃ ┣ 📂 middlewares
┃ ┣ 📂 routes
┣ 📂 usecases
┣ 📜 .env
┣ 📜 Makefile
┣ 📜 README.md
┣ 📜 docker-compose.yaml
┣ 📜 go.mod
┗ 📜 go.sum
```
