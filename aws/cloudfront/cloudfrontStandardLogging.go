package cloudfront

import (
	"github.com/stangirard/yatas/plugins/commons"
)

func CheckIfStandardLogginEnabled(checkConfig commons.CheckConfig, d []SummaryToConfig, testName string) {
	var check commons.Check
	check.InitCheck("Cloudfronts queries are logged", "Check if all cloudfront distributions have standard logging enabled", testName, []string{"Security", "Good Practice"})
	for _, cc := range d {

		if cc.config.Logging != nil && cc.config.Logging.Enabled != nil && *cc.config.Logging.Enabled {
			Message := "Standard logging is enabled on " + *cc.summary.Id
			result := commons.Result{Status: "OK", Message: Message, ResourceID: *cc.summary.ARN}
			check.AddResult(result)
		} else {
			Message := "Standard logging is not enabled on " + *cc.summary.Id
			result := commons.Result{Status: "FAIL", Message: Message, ResourceID: *cc.summary.ARN}
			check.AddResult(result)
		}
	}
	checkConfig.Queue <- check
}
