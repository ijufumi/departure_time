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

Execute the below command.

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

## Points of design

### Ease to test

Ensure to write the test code easily, it uses DI to set an instance to others.


### Ease to set configuration

This service allows users to get configuration variables from `.env` file and environmental variables of OS.


### Ease to run app

On local pc, it uses `Docker Compose` to run app.

Aside from that, it provides the CDK code to create an environment on AWS.


## Points of improvement

### Select email vendor

Currently, this service sends email using AWS SES and SendGrid.

First, it sends email using AWS SES. 
If an error occurs at that time, send it using SendGrid.

So, it always uses SES unless the error occurs on SES.

It is better to be able to select email vendor when user send an email.


### Support HTML mail

Currently, it supports text email only.

However, most of emails are HTML email.

It is better to support HTML email.

### Send email to multiple email address

Now, this service sneds an email only one address.

But, there is possible that user want to send an email to multiple addresses.

It is better to send an email to multiple addresses.


### Send email as CC or BCC

This service allows users to send an email as a To address only.

It is better to change it can send an email as a CC or BCC.


### Handle error properly

The API returns error as Internal Server Error when error occurs by sending email.

However, it is better to return the error by error details.

## References

* [Go で AWS SES SDK を使ってメールを送る](https://blog.giftee.dev/2022-01-31-go-aws-ses-sdk/)
* [AWS SDK for Go API Reference](https://docs.aws.amazon.com/sdk-for-go/api/service/sesv2/)
* [Go - ドキュメント | SendGrid](https://sendgrid.kke.co.jp/docs/Integrate/Code_Examples/v3_Mail/go.html)
* [Go Bindings for AWS CDK](https://github.com/aws/aws-cdk-go)