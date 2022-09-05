package stacks

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
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
	}
	ecsCluster := awsecs.NewCluster(scope, jsii.String(clusterID), &clusterProps)
	certificate := awscertificatemanager.Certificate_FromCertificateArn(scope, jsii.String(fmt.Sprintf("certificate-id-%s", configuration.Cluster.Name)), jsii.String(configuration.LoadBalancer.CertificateArn))
	hostedZoneQuery := awsroute53.HostedZoneProviderProps{
		DomainName: &configuration.Route53.DomainName,
	}
	repository := awsecr.Repository_FromRepositoryName(scope, jsii.String("ecr-repository"), &configuration.Cluster.Image.RepositoryName)
	image := awsecs.ContainerImage_FromEcrRepository(repository, &configuration.Cluster.Image.Tag)
	awsAccessKey := awsssm.NewStringParameter(scope, jsii.String("aws-access-key"), &awsssm.StringParameterProps{
		ParameterName: jsii.String("app-aws-access-key"),
		StringValue:   &configuration.App.Aws.AwsAccessKeyID,
		Type:          awsssm.ParameterType_STRING,
	})
	awsSecretAccessKey := awsssm.NewStringParameter(scope, jsii.String("aws-secret-access-key"), &awsssm.StringParameterProps{
		ParameterName: jsii.String("app-aws-secret-key"),
		StringValue:   &configuration.App.Aws.AwsSecretAccessKey,
		Type:          awsssm.ParameterType_STRING,
	})
	sendGridKey := awsssm.NewStringParameter(scope, jsii.String("send-grid-key"), &awsssm.StringParameterProps{
		ParameterName: jsii.String("app-send-grid-key"),
		StringValue:   &configuration.App.SendGrid.SendGridAPIKEY,
		Type:          awsssm.ParameterType_STRING,
	})

	ecsTaskRole := awsiam.NewRole(scope, jsii.String("ecs-task-role"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("ecs-tasks.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
		InlinePolicies: &map[string]awsiam.PolicyDocument{
			"inlinePolicy1": awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
				Statements: &[]awsiam.PolicyStatement{
					awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
						Actions:   jsii.Strings("ssm:Get*", "ssm:Describe*", "ssm:List*"),
						Resources: jsii.Strings("arn:aws:ssm:*"),
					}),
				},
			}),
		},
	})
	ecsServiceProps := awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		Cluster:            ecsCluster,
		DesiredCount:       jsii.Number(1),
		ListenerPort:       jsii.Number(443),
		RedirectHTTP:       jsii.Bool(true),
		PublicLoadBalancer: jsii.Bool(true),
		Protocol:           awselasticloadbalancingv2.ApplicationProtocol_HTTPS,
		DomainName:         jsii.String(configuration.LoadBalancer.DomainName),
		DomainZone:         awsroute53.PublicHostedZone_FromLookup(scope, jsii.String(fmt.Sprintf("zone-id")), &hostedZoneQuery),
		Certificate:        certificate,
		TaskSubnets: &awsec2.SubnetSelection{
			Subnets: vpc.PrivateSubnets(),
		},
		TaskImageOptions: &awsecspatterns.ApplicationLoadBalancedTaskImageOptions{
			Image:         image,
			ContainerPort: jsii.Number(8080),
			Environment: &map[string]*string{
				"GIN_MODE":   &configuration.App.GinMode,
				"AWS_REGION": &configuration.App.Aws.AwsRegion,
			},
			Secrets: &map[string]awsecs.Secret{
				"AWS_ACCESS_KEY_ID":     awsecs.Secret_FromSsmParameter(awsAccessKey),
				"AWS_SECRET_ACCESS_KEY": awsecs.Secret_FromSsmParameter(awsSecretAccessKey),
				"SENDGRID_API_KEY":      awsecs.Secret_FromSsmParameter(sendGridKey),
			},
			TaskRole: ecsTaskRole,
		},
	}
	ecsServiceID := fmt.Sprintf("ecs-service-%s", configuration.Cluster.Name)
	awsecspatterns.NewApplicationLoadBalancedFargateService(scope, jsii.String(ecsServiceID), &ecsServiceProps)
}
