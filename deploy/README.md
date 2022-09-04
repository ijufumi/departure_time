# Deployment 

Create AWS environment using AWS CDK.

## Overall

![Overall](./docs/overall.png)

## Requirements

* awscli
* docker

## Preparation

### Precondition

* Have a AWS Account
* Have a own domain name
* Already created a certification of ACM

### Enable AWS account
If your profile is not `default`, you should execute the below command.

```bash
export AWS_PROFILE=[your profile name]
```

### Get authorization from repository

```bash
aws ecr get-login-password --region [your region] | docker login --username AWS --password-stdin [your aws account].dkr.ecr.[your region].amazonaws.com
```

### Create ECR repository

```bash
aws ecr create-repository --repository-name [your repository name]
```

### Build a image

```bash
docker build -t [your repository uri]:[tag] .
```

and push

```bash
docker build -t [your repository uri]:[tag] .
```

### Create `.env`

Crete `.env` file by copying from `.env.example`.
And then, fill out the below items.

| Item | Description |
| ---- | ---- |
| `CIDR_BLOCK` | A CIDR block of a new VPC |
| `VPC_NAME` | A name of a new VPC |
| `CLUSTER_NAME` | a name of a new ECS cluster |
| `CLUSTER_IMAGE_REPOSITORY_NAME` | The ECR repository name used by ECS |
| `CLUSTER_IMAGE_TAG` | A image tag of image used by ECS |
| `LOAD_BALANCER_DOMAIN_NAME` | The load balancer name to connect |
| `LOAD_BALANCER_CERTIFICATE_ARN` | The ARN of certification used by ALB |
| `ROUTE53_DOMAIN_NAME` | The hostzoned name |

## Initialize CDK

```bash
cdk bootstrap
```

## Confirm CDK

```bash
cdk diff
```

## Execute CDK

```bash
cdk deploy
```
