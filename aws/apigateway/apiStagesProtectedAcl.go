package apigateway

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/stangirard/yatas/plugins/commons"
)

func CheckIfStagesProtectedByAcl(checkConfig commons.CheckConfig, stages map[string][]types.Stage, testName string) {
	var check commons.Check
	check.InitCheck("ApiGateways are protected by an ACL", "Check if all stages are protected by ACL", testName, []string{"Security", "Good Practice"})
	for apigateway, id := range stages {
		for _, stage := range id {
			if stage.WebAclArn != nil && *stage.WebAclArn != "" {
				Message := "Stage " + *stage.StageName + " is protected by ACL" + " of ApiGateway " + apigateway
				result := commons.Result{Status: "OK", Message: Message, ResourceID: GetStageArn(checkConfig.ConfigAWS, apigateway, stage)}
				check.AddResult(result)
			} else {
				Message := "Stage " + *stage.StageName + " is not protected by ACL" + " of ApiGateway " + apigateway
				result := commons.Result{Status: "FAIL", Message: Message, ResourceID: GetStageArn(checkConfig.ConfigAWS, apigateway, stage)}
				check.AddResult(result)
			}
		}
	}
	checkConfig.Queue <- check
}
