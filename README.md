<p align="center" style="margin-top: 40px; margin-bottom: 40px">
  <img src="https://unico.io/wp-content/uploads/2021/09/logo-unico.svg" height="80">
  <h1 align="center">TESTE DE API</h1>
  <p align="center">
    <!-- <a href="https://github.com/booscaaa/go-api-test-unico/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/booscaaa/go-api-test-unico.svg?style=for-the-badge"></a> -->
    <!-- <a href="https://github.com/booscaaa/go-api-test-unico/actions/workflows/test.yaml"><img alt="Test status" src="https://img.shields.io/github/workflow/status/booscaaa/go-api-test-unico/Test?label=TESTS&style=for-the-badge"></a>
    <a href="https://codecov.io/gh/booscaaa/go-api-test-unico"><img alt="Coverage" src="https://img.shields.io/codecov/c/github/booscaaa/go-api-test-unico/master.svg?style=for-the-badge"></a> -->
  </p>
</p>


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

- ENDPOINT_API: http://localhost:3000
- ENDPOINT_PGADMIN: http://localhost:5050

<br><br>

## Credênciais para desenvolvimento
### PGADMIN
- login: admin@admin.com
- senha: admin

<br><br>


## Bibliotecas

- [Viper](https://github.com/spf13/viper)
- [Migrate](github.com/golang-migrate/migrate)
- [Cobra](github.com/spf13/cobra)
- [Gin](github.com/gin-gonic/gin) 
