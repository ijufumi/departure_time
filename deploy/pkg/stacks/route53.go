package stacks

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ijufumi/email-service/deploy/pkg/config"
)

func CreateRoute53(scope constructs.Construct, ecsService awsecspatterns.ApplicationLoadBalancedFargateService) {
	configuration := config.Load()

	cnamePropss := awsroute53.CnameRecordProps{
		Zone:       awsroute53.HostedZone_FromHostedZoneId(scope, jsii.String(fmt.Sprintf("zone-id")), jsii.String(configuration.Route53.ZoneID)),
		RecordName: ecsService.LoadBalancer().LoadBalancerDnsName(),
		DomainName: &configuration.LoadBalancer.DomainName,
	}

	awsroute53.NewCnameRecord(scope, jsii.String(fmt.Sprintf("route53-record-%s", configuration.Cluster.Name)), &cnamePropss)
}
