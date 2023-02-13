package backup

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
)

func GetBackupPlans(s aws.Config) []types.BackupPlansListMember {

	svc := backup.NewFromConfig(s)

	params := &backup.ListBackupPlansInput{}
	resp, err := svc.ListBackupPlans(context.TODO(), params)
	if err != nil {
		fmt.Println(err)
	}

	return resp.BackupPlansList
}

type BackupPlanToSelection struct {
	BackupPlanName string
	Selections     []types.BackupSelection
}

func GetBackupToSelections(s aws.Config, plans []types.BackupPlansListMember) []BackupPlanToSelection {

	svc := backup.NewFromConfig(s)

	var backupPlanToSelections []BackupPlanToSelection
	for _, plan := range plans {
		params := &backup.ListBackupSelectionsInput{
			BackupPlanId: plan.BackupPlanId,
		}
		resp, err := svc.ListBackupSelections(context.TODO(), params)
		if err != nil {
			fmt.Println(err)
		}
		var selections []types.BackupSelection
		for _, s := range resp.BackupSelectionsList {
			params := &backup.GetBackupSelectionInput{
				BackupPlanId: plan.BackupPlanId,
				SelectionId:  s.SelectionId,
			}
			resp, err := svc.GetBackupSelection(context.TODO(), params)
			if err != nil {
				fmt.Println(err)
			}
			selections = append(selections, *resp.BackupSelection)
		}
		backupPlanToSelections = append(backupPlanToSelections, BackupPlanToSelection{
			BackupPlanName: *plan.BackupPlanName,
			Selections:     selections,
		})
	}
	return backupPlanToSelections
}
