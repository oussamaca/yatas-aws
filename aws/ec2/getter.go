package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type EC2GetObjectAPI interface {
	DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

func GetEC2s(svc EC2GetObjectAPI) []types.Instance {
	input := &ec2.DescribeInstancesInput{}
	result, err := svc.DescribeInstances(context.TODO(), input)
	if err != nil {
		fmt.Println(err)
	}
	var instances []types.Instance
	for _, r := range result.Reservations {
		instances = append(instances, r.Instances...)
	}
	for {
		if result.NextToken == nil {
			break
		}
		input.NextToken = result.NextToken
		result, err = svc.DescribeInstances(context.TODO(), input)
		if err != nil {
			fmt.Println(err)
		}
		for _, r := range result.Reservations {
			instances = append(instances, r.Instances...)
		}
	}

	return instances
}

// GetInstanceArn returns the arn of an instance
func GetInstanceArn(s aws.Config, instance types.Instance) string {
	return "arn:aws:ec2:" + s.Region + "::instance/" + *instance.InstanceId
}
