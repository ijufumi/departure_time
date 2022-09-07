# Email Service

This code is implementation of [Coding Challenge](https://github.com/ArentInc/coding-challenge-tools/blob/master/coding_challenge.md).

## Package structure

Basically, the package structure follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

```
.
├── api                     # API Specification
├── cmd                     # Entrypoint of launching app
├── deploy                  # The code to create environment
├── internal                # Code for internal use
│   ├── app                 # Code for app specific
│   │   ├── container       # Create DI container
│   │   ├── http            # Code for API specific
│   │   │   ├── hander      # Request handler
│   │   │   ├── request     # Data object for reuqest
│   │   │   └── response    # Data object for response
│   │   └── service         # Business logic
│   └── pkg                 # Common codes
├── web                     # Frontend codes
├── .env.example
├── docker-compose.yaml
├── Dockerfile
├── LICENSE
└── README.md

```

## How to run app
### Common
#### Prepare mail services
* Amazon SES
  * [Amazon Simple Email Service を設定する](https://docs.aws.amazon.com/ja_jp/ses/latest/dg/setting-up.html)
* SendGrid
  * [SMTPメールの構築](https://sendgrid.kke.co.jp/docs/API_Reference/SMTP_API/building_an_smtp_email.html)

### for local
#### Prepare `.env` file

Copy `.env.example` to `.env` and edit it.
Should input these values

* `MAIL_FROM_ADDRESS`
* `AWS_ACCESS_KEY_ID`
* `AWS_SECRET_KEY`
* `AWS_SES_ENDPOINT`
* `AWS_REGION`
* `SENDGRID_API_KEY`

#### Execute 

```bash
COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILCDKIT=1 docker-compose up --build
```

And then, you aceess to [localhost:8080](http://localhost:8080)

### for AWS
#### Create environment
See [deploy](./deploy)


## How to execute test

Execute all tests.

```bash
go test ./...
```