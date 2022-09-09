package rds

import (
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/stangirard/yatas/plugins/commons"
)

func Test_checkIfBackupEnabled(t *testing.T) {
	type args struct {
		checkConfig commons.CheckConfig
		instances   []types.DBInstance
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_checkIfBackupEnabled",
			args: args{
				checkConfig: commons.CheckConfig{
					Wg:    &sync.WaitGroup{},
					Queue: make(chan commons.Check, 1),
				},
				instances: []types.DBInstance{
					{
						DBInstanceIdentifier:  aws.String("test"),
						DBInstanceArn:         aws.String("arn:aws:rds:us-east-1:123456789012:db:test"),
						StorageEncrypted:      true,
						BackupRetentionPeriod: 7,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkIfBackupEnabled(tt.args.checkConfig, tt.args.instances, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "OK" {
						t.Errorf("CheckIfBackup() = %v, want %v", check.Status, "OK")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}

func Test_checkIfBackupEnabledFail(t *testing.T) {
	type args struct {
		checkConfig commons.CheckConfig
		instances   []types.DBInstance
		testName    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_checkIfBackupEnabled",
			args: args{
				checkConfig: commons.CheckConfig{
					Wg:    &sync.WaitGroup{},
					Queue: make(chan commons.Check, 1),
				},
				instances: []types.DBInstance{
					{
						DBInstanceIdentifier:  aws.String("test"),
						DBInstanceArn:         aws.String("arn:aws:rds:us-east-1:123456789012:db:test"),
						StorageEncrypted:      true,
						BackupRetentionPeriod: 0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkIfBackupEnabled(tt.args.checkConfig, tt.args.instances, tt.args.testName)
			tt.args.checkConfig.Wg.Add(1)
			go func() {
				for check := range tt.args.checkConfig.Queue {
					if check.Status != "FAIL" {
						t.Errorf("CheckIfBackup() = %v, want %v", check.Status, "FAIL")
					}
					tt.args.checkConfig.Wg.Done()
				}
			}()
			tt.args.checkConfig.Wg.Wait()
		})
	}
}
