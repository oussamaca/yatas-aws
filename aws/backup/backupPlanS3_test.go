package backup

import (
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/padok-team/yatas/plugins/commons"
)

func TestCheckIfBackupPlanS3(t *testing.T) {
	type args struct {
		checkConfig commons.CheckConfig
		plans       []BackupPlanToSelection
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCheckIfBackupPlanS3",
			args: args{
				checkConfig: commons.CheckConfig{Queue: make(chan commons.Check, 1), Wg: &sync.WaitGroup{}},
				plans: []BackupPlanToSelection{
					{
						BackupPlanName: *aws.String("test"),
						Selections: []types.BackupSelection{
							{
								Resources: []string{"*"},
							},
						},
					},
				},
			},
		},
		{
			name: "TestCheckIfBackupPlanS3",
			args: args{
				checkConfig: commons.CheckConfig{Queue: make(chan commons.Check, 1), Wg: &sync.WaitGroup{}},
				plans: []BackupPlanToSelection{
					{
						BackupPlanName: *aws.String("test"),
						Selections: []types.BackupSelection{
							{
								Resources: []string{"arn:aws:s3:::*"},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckIfBackupPlanS3(tt.args.checkConfig, tt.args.plans, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "OK" {
						t.Errorf("CheckIfBackupPlanS3() = %v, want %v", check.Status, "OK")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}

func TestCheckIfBackupPlanS3Fail(t *testing.T) {
	type args struct {
		checkConfig commons.CheckConfig
		plans       []BackupPlanToSelection
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCheckIfBackupPlanS3",
			args: args{
				checkConfig: commons.CheckConfig{Queue: make(chan commons.Check, 1), Wg: &sync.WaitGroup{}},
				plans: []BackupPlanToSelection{
					{
						BackupPlanName: *aws.String("test"),
						Selections:     []types.BackupSelection{},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckIfBackupPlanS3(tt.args.checkConfig, tt.args.plans, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "FAIL" {
						t.Errorf("CheckIfBackupPlanS3() = %v, want %v", check.Status, "FAIL")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}
