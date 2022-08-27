# Email Service

This code is implementation of [Coding Challenge](https://github.com/ArentInc/coding-challenge-tools/blob/master/coding_challenge.md).

## Package structure

Basically, the package structure follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

## How to run app

### for local

#### Prepare `.env` file

Copy `.env.example` to `.env` and edit it.
Should input these values

* `AWS_ACCESS_KEY_ID`
* `AWS_SECRET_KEY`
* `AWS_SES_ENDPOINT`
* `AWS_REGION`

#### Execute 

```bash
COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILCDKIT=1 docker-compose up --build
```

## How to execute test

Execute all tests.

```bash
go test ./...
```