### Recursos utilizados

- Golang
- Teste Unitário
- PostgreSQL
- Docker
- DotEnv
- Domain-Drive-Design(DDD)

### Sobre o projeto

O modelo construído foi utilizando uma ideia de Domain-Drive-Design(DDD) Onde separamos os componentes da aplicação baseada em camadas.

Foi utilizado Golang como linguagem para a construção do Backend, para manutenção do banco de dados optei por usar Docker para gerir o contêiner do PostgreSQL, e para dar manutenção ao código e pensando em desenvolvimento contínuo foi implementado a uma biblioteca de migrações de banco de dados, assim ao iniciar a aplicação ele checa a última modificação feita, e caso haja uma versão nova, ele ira rodar para manter atualizado, e com suas dependências corretas.

Para a modelagem da RestAPI, a biblioteca Fiber foi escolhida, em virtude do benchmarking de tempos de resposta e requisições em relação a outras bibliotecas do Golang, também foi modelado uma rota de teste para checar se a aplicação está aceitando requisições ou se a mesma se encontra com algum problema. Assim podendo aferir se a API está rodando.

Partindo da premissa que um código pode evoluir, a utilização de uma arquitetura correta é essencial para escalabilidade e manutenibilidade do projeto. Assim todas as camadas do projeto, tendem a seguir o seu conceito e a sua ideia no mesmo.

Agrupando todas as ideias, foi escolhido utilizar variáveis de ambiente, pensando na facilidade do trabalho de equipe e gerenciamento de variáveis, como também na segurança da aplicação. Finalizando o processo de desenvolvimento do projeto, é essencial a utilização de testes para verificar se as funções estão funcionando da maneira que isoladamente devem funcionar e também como executariam em conjunto com a integração do serviço. 

Atualmente o projeto implementou as seguintes interfaces do dominio para o banco de dados:

- Adicionar uma conversão ao banco de dados
- Encontrar uma conversão a partir de um ID
- Encontrar todas as conversões realizadas

## Rotas da aplicação

    Base: "http://localhost:5000"

# Rota de segurança
  ```
  Rota: ".../ping"
  ex: "http://localhost:5000/ping"
  ```

Rota responsável apenas por realizar checar a saúde da API, aferindo se a mesma se encontra funcionando.

# Grupo de rotas "exchange"
    Rota: ".../{amount}/{from}/{to}/{rate}"  
    ex: http://localhost:5000/exchange/10/BRL/USD/4.73

Rota responsável por realizar a conversão do par de moedas e persistir no banco o registro da mesma
   ```
  Rota: ".../all"
  ex: "http://localhost:5000/exchange/all"
  ```

Rota responsável por localizar todos as conversões já realizadas dentro da API.
  ```
  Rota: ".../{id}"
  ex: "http://localhost:5000/exchange/2"
  ```

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
