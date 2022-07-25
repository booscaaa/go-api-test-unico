<p align="center">
  <br>
  <br>
  <br>
  <br>
  <br>
  <img src="https://unico.io/wp-content/uploads/2021/09/logo-unico.svg" height="80">
  <br>
  <br>
  <br>
  <br>
  <br>
  <br>
  <h1 align="center">TESTE DE API</h1>
  <p align="center">
    <a href="https://github.com/booscaaa/go-api-test-unico/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/booscaaa/go-api-test-unico.svg?style=for-the-badge"></a>
    <a href="https://github.com/booscaaa/go-api-test-unico/actions/workflows/test.yaml"><img alt="Test status" src="https://img.shields.io/github/workflow/status/booscaaa/go-api-test-unico/Tests?label=TESTS&style=for-the-badge"></a>
    <a href="https://codecov.io/gh/booscaaa/go-api-test-unico"><img alt="Coverage" src="https://img.shields.io/codecov/c/github/booscaaa/go-api-test-unico/master.svg?style=for-the-badge"></a>
    <a href="https://api-unico-test.herokuapp.com/swagger/index.html"><img alt="Swagger docs" src="https://img.shields.io/badge/api docs-swagger-purple?style=for-the-badge"></a>
  </p>
</p>
<br>
<br>
<br>
<br>


## Pré-requisitos

Para rodar o projeto se faz necessário a instalação de alguns recursos.

- A SDK do Go [nesse link](https://golang.org/).
- O Docker [nesse link](https://docs.docker.com/).

<br>
<br>

### Baixando

```bash
$ git clone git@github.com:booscaaa/go-api-test-unico.git
$ cd go-api-test-unico
```

<br>

### Rodando

```bash
$ docker-compose up --build -d
```

### Importando os dados do CSV do teste

```bash
$ docker exec api go run ./adapter/cli/main.go importer
```

- ENDPOINT_API: http://localhost:3000
- ENDPOINT_API_DOCS: http://localhost:3000/swagger/index.html
- ENDPOINT_PGADMIN: http://localhost:5050

- METODOS:

  <a href="https://insomnia.rest/run?label=API%20UNICO%20V1&uri=https://api-unico-test.herokuapp.com/swagger/doc.json"><img alt="Swagger docs" src="https://insomnia.rest/images/run.svg" alt="Run in Insomnia"></a>  

<br><br>

## Credênciais para desenvolvimento
### PGADMIN
- login: admin@admin.com
- senha: admin

<br><br>


## Bibliotecas

- [Viper](github.com/spf13/viper)
- [Migrate](github.com/golang-migrate/migrate)
- [Cobra](github.com/spf13/cobra)
- [Gin](github.com/gin-gonic/gin) 
- [SQLMock](github.com/DATA-DOG/go-sqlmock)
- [GoPaginate](github.com/booscaaa/go-paginate)
- [Faker](github.com/bxcodec/faker) 
- [Validator](github.com/go-playground/validator) 
- [Mock](github.com/golang/mock)
- [Sqlx](github.com/jmoiron/sqlx)
- [Swag](github.com/swaggo/swag)
- [Zap](go.uber.org/zap)