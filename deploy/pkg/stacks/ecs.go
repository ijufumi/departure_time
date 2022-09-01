package stacks

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ijufumi/email-service/deploy/pkg/config"
)

func CreateECS(scope constructs.Construct, vpc awsec2.Vpc) {
	configuration := config.Load()
	clusterID := fmt.Sprintf("id-%s", configuration.Cluster.Name)
	props := awsecs.ClusterProps{
		ClusterName: jsii.String(configuration.Cluster.Name),
		Vpc:         vpc,
		Capacity: &awsecs.AddCapacityOptions{
			DesiredCapacity: jsii.Number(0),
			MaxCapacity:     jsii.Number(1),
			MinCapacity:     jsii.Number(0),
			InstanceType:    awsec2.InstanceType_Of(awsec2.InstanceClass_M5A, awsec2.InstanceSize_MEDIUM),
		},
	}
	ecsCluster := awsecs.NewCluster(scope, jsii.String(clusterID), &props)
}
