package backup

import (
	"fmt"

	"github.com/padok-team/yatas/plugins/commons"
)

func CheckIfBackupPlanS3(checkConfig commons.CheckConfig, plans []BackupPlanToSelection, testName string) {
	var check commons.Check
	check.InitCheck("Backup plan exists for S3 Buckets", "Check if Backup plan is present for S3 Buckets", testName, []string{"Security", "Good Practice"})
	var s3BackupPlans []string
	for _, plan := range plans {
		for _, selection := range plan.Selections {
			for _, res := range selection.Resources {
				if res == "arn:aws:s3:::*" || res == "*" {
					s3BackupPlans = append(s3BackupPlans, plan.BackupPlanName)
				}
			}
		}
	}

	if len(s3BackupPlans) == 0 {
		Message := "AWS Backup does not contain a backup plan for S3 Buckets"
		result := commons.Result{Status: "FAIL", Message: Message, ResourceID: ""}
		check.AddResult(result)
	} else {
		Message := "AWS Backup contains a backup plan for S3 Buckets: " + fmt.Sprint(s3BackupPlans)
		result := commons.Result{Status: "OK", Message: Message, ResourceID: ""}
		check.AddResult(result)
	}
	checkConfig.Queue <- check
}
