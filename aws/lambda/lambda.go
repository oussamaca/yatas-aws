package lambda

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/padok-team/yatas/plugins/commons"
)

func RunChecks(wa *sync.WaitGroup, s aws.Config, c *commons.Config, queue chan []commons.Check) {

	var checkConfig commons.CheckConfig
	checkConfig.Init(s, c)
	var checks []commons.Check
	lambdas := GetLambdas(s)

	go commons.CheckTest(checkConfig.Wg, c, "AWS_LMD_001", CheckIfLambdaPrivate)(checkConfig, lambdas, "AWS_LMD_001")
	go commons.CheckTest(checkConfig.Wg, c, "AWS_LMD_002", CheckIfLambdaInSecurityGroup)(checkConfig, lambdas, "AWS_LMD_002")
	go commons.CheckTest(checkConfig.Wg, c, "AWS_LMD_003", CheckIfLambdaNoErrors)(checkConfig, lambdas, "AWS_LMD_003")
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
