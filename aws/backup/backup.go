package backup

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/padok-team/yatas/plugins/commons"
)

func RunChecks(wa *sync.WaitGroup, s aws.Config, c *commons.Config, queue chan []commons.Check) {

	var checkConfig commons.CheckConfig
	checkConfig.Init(s, c)
	var checks []commons.Check
	plans := GetBackupPlans(s)
	backupPlanToSelections := GetBackupToSelections(s, plans)

	go commons.CheckTest(checkConfig.Wg, c, "AWS_BAC_001", CheckIfBackupPlanS3)(checkConfig, backupPlanToSelections, "AWS_BAC_001")

	go func() {
		for t := range checkConfig.Queue {
			t.EndCheck()
			checks = append(checks, t)

			checkConfig.Wg.Done()

		}
	}()

	checkConfig.Wg.Wait()

	queue <- checks
}
