package stacks

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ijufumi/email-service/deploy/pkg/config"
)

func CreateECS(scope constructs.Construct, vpc awsec2.Vpc) {
	configuration := config.Load()
	clusterID := fmt.Sprintf("id-%s", configuration.Cluster.Name)
	clusterProps := awsecs.ClusterProps{
		ClusterName: jsii.String(configuration.Cluster.Name),
		Vpc:         vpc,
		Capacity: &awsecs.AddCapacityOptions{
			DesiredCapacity: jsii.Number(0),
			MaxCapacity:     jsii.Number(1),
			MinCapacity:     jsii.Number(0),
			InstanceType:    awsec2.InstanceType_Of(awsec2.InstanceClass_M5A, awsec2.InstanceSize_MEDIUM),
		},
	}
	ecsCluster := awsecs.NewCluster(scope, jsii.String(clusterID), &clusterProps)
	certificate := awscertificatemanager.Certificate_FromCertificateArn(scope, jsii.String(fmt.Sprintf("certificate-id-%s", configuration.Cluster.Name)), jsii.String(configuration.LoadBalancer.CertificateArn))
	ecsServiceProps := awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		Cluster:            ecsCluster,
		DesiredCount:       jsii.Number(0),
		ListenerPort:       jsii.Number(443),
		RedirectHTTP:       jsii.Bool(true),
		PublicLoadBalancer: jsii.Bool(true),
		Protocol:           awselasticloadbalancingv2.ApplicationProtocol_HTTPS,
		DomainName:         jsii.String(configuration.LoadBalancer.DomainName),
		Certificate:        certificate,
		TaskSubnets: &awsec2.SubnetSelection{
			Subnets: vpc.PrivateSubnets(),
		},
		TaskImageOptions: &awsecspatterns.ApplicationLoadBalancedTaskImageOptions{
			Image:         awsecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample"), &awsecs.RepositoryImageProps{}),
			ContainerPort: jsii.Number(8080),
		},
	}
	ecsServiceID := fmt.Sprintf("ecs-service-%s", configuration.Cluster.Name)
	awsecspatterns.NewApplicationLoadBalancedFargateService(scope, jsii.String(ecsServiceID), &ecsServiceProps)

}
