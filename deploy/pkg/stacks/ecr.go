package stacks

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ijufumi/email-service/deploy/pkg/config"
)

func CreateECR(scope constructs.Construct) {
	configuration := config.Load()
	props := awsecr.RepositoryProps{
		RepositoryName: jsii.String(configuration.Repository.Name),
	}

	repositoryID := fmt.Sprintf("id-%s", configuration.Repository.Name)
	awsecr.NewRepository(scope, jsii.String(repositoryID), &props)
}
