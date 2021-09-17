## Technologies
- Golang
- Unit Tests
- PostgreSQL
- Docker
- Environmental Variables
- Domain Drive Design (DDD)

## About the project


The model built was using an idea of ​​Domain Drive Design (DDD). Where we separate the application components based on layers.

**Golang** was used as the language to build the Backend, for database maintenance I chose to use **Docker** to manage the **PostgreSQL** container, and to maintain the code and also thinking about continuous development, a **database migration** library was implemented , so when starting the application it checks the last modification made, and if there is a new version, it will run to keep it updated, and with its correct dependencies.

For the RestAPI modeling, the **Fiber** library was chosen, due to the benchmarking of response times and requests against other Golang libraries, a test route was also modeled to check if the application is accepting requests or if it is encounters some problem. Thus being able to check if the API is running.

Based on the premise that code can evolve, the use of a correct **architecture** is essential for project scalability and maintainability. So all layers of the project tend to follow your concept and your idea in it.

Grouping all the ideas, it was chosen to use environment variables, thinking about the ease of teamwork and variable management, as well as the security of the application. At the end of the project development process, it is essential to **use tests** to verify that the functions are working the way they should work separately and also how they would perform in conjunction with the service integration.

## Routes
  ```
Method: GET
Path: http://localhost:5000/ping
Check if application is running
  ```
```
Method: POST
Path: http://localhost:5000/exchange/{amount}/{from}/{to}/{rate}
Example: http://localhost:5000/exchange/10/BRL/USD/4.73
Make ocurrency conversion and save on database
```
```
Method: GET
Path: http://localhost:5000/exchange/all
Return all currency conversion
  ```

  ```
Method: GET
Path: http://localhost:5000/exchange/:id
Return currency conversion from informed id
  ```
## Project architecture

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
