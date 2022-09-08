package lambda

import (
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/stangirard/yatas/config"
)

func TestCheckIfLambdaPrivate(t *testing.T) {
	type args struct {
		checkConfig config.CheckConfig
		lambdas     []types.FunctionConfiguration
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCheckIfLambdaPrivate",
			args: args{
				checkConfig: config.CheckConfig{
					Wg:    &sync.WaitGroup{},
					Queue: make(chan config.Check, 1),
				},
				lambdas: []types.FunctionConfiguration{
					{
						FunctionName: aws.String("test"),
						FunctionArn:  aws.String("arn:aws:lambda:us-east-1:123456789012:function:test"),
						VpcConfig: &types.VpcConfigResponse{
							VpcId: aws.String("vpc-123456789012"),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckIfLambdaPrivate(tt.args.checkConfig, tt.args.lambdas, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "OK" {
						t.Errorf("CheckifLambdaPrivate() = %v, want %v", check.Status, "OK")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}

func TestCheckIfLambdaPrivateFail(t *testing.T) {
	type args struct {
		checkConfig config.CheckConfig
		lambdas     []types.FunctionConfiguration
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCheckIfLambdaPrivate",
			args: args{
				checkConfig: config.CheckConfig{
					Wg:    &sync.WaitGroup{},
					Queue: make(chan config.Check, 1),
				},
				lambdas: []types.FunctionConfiguration{
					{
						FunctionName: aws.String("test"),
						FunctionArn:  aws.String("arn:aws:lambda:us-east-1:123456789012:function:test"),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckIfLambdaPrivate(tt.args.checkConfig, tt.args.lambdas, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "FAIL" {
						t.Errorf("CheckifLambdaPrivate() = %v, want %v", check.Status, "FAIL")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}