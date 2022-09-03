# Deployment 

## Requirements

* awscli
* docker

## Setup CDK

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

### Initialize CDK

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
