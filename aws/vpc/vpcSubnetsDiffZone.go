package vpc

import (
	"github.com/stangirard/yatas/plugins/commons"
)

func CheckIfSubnetInDifferentZone(checkConfig commons.CheckConfig, vpcToSubnets []VPCToSubnet, testName string) {
	var check commons.Check
	check.InitCheck("VPC's Subnets are in different zones", "Check if Subnet are in different zone", testName, []string{"Security", "Good Practice"})
	for _, vpcToSubnet := range vpcToSubnets {
		subnetsAZ := make(map[string]int)
		for _, subnet := range vpcToSubnet.Subnets {
			subnetsAZ[*subnet.AvailabilityZone]++
		}
		if len(subnetsAZ) > 1 {
			Message := "Subnets are in different zone on " + vpcToSubnet.VpcID
			result := commons.Result{Status: "OK", Message: Message, ResourceID: GetVPCArn(checkConfig.ConfigAWS, vpcToSubnet.VpcID)}
			check.AddResult(result)
		} else {
			Message := "Subnets are in same zone on " + vpcToSubnet.VpcID
			result := commons.Result{Status: "FAIL", Message: Message, ResourceID: GetVPCArn(checkConfig.ConfigAWS, vpcToSubnet.VpcID)}
			check.AddResult(result)
		}
	}
	checkConfig.Queue <- check
}
