package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ijufumi/email-service/deploy/pkg/config"
	"github.com/ijufumi/email-service/deploy/pkg/stacks"
)

type DeployStackProps struct {
	awscdk.StackProps
}

func NewDeployStack(scope constructs.Construct, id string, props *DeployStackProps) awscdk.Stack {
	var stackProps awscdk.StackProps
	if props != nil {
		stackProps = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &stackProps)

	vpc := stacks.CreateVPC(stack)
	stacks.CreateECS(stack, vpc)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewDeployStack(app, "DeployStack", &DeployStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	configuration := config.Load()
	var account = configuration.AwsAccessKeyID
	if len(account) == 0 {
		account = configuration.CdkDefaultAccount
	}
	var region = configuration.AwsRegion
	if len(region) == 0 {
		region = configuration.CdkDefaultRegion
	}
	return &awscdk.Environment{
		Account: jsii.String(account),
		Region:  jsii.String(region),
	}

}
