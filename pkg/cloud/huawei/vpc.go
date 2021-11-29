package huawei

import (
	"fmt"

	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/model"
)

func (p *HuaweiCloud) CreateVPC(req cloud.CreateVpcRequest) (cloud.CreateVpcResponse, error) {

	return cloud.CreateVpcResponse{}, nil
}

func (p *HuaweiCloud) GetVPC(req cloud.GetVpcRequest) (cloud.GetVpcResponse, error) {

	return cloud.GetVpcResponse{}, nil
}

func (p *HuaweiCloud) DescribeVpcs(req cloud.DescribeVpcsRequest) (cloud.DescribeVpcsResponse, error) {
	var page int32 = 1
	vpcs := make([]cloud.VPC, 0, 128)
	for {

		if 1 > page*50 {
			page++
		} else {
			break
		}
	}

	return cloud.DescribeVpcsResponse{Vpcs: vpcs}, nil
}

func (p *HuaweiCloud) CreateSwitch(req cloud.CreateSwitchRequest) (cloud.CreateSwitchResponse, error) {

	return cloud.CreateSwitchResponse{}, nil
}

func (p *HuaweiCloud) GetSwitch(req cloud.GetSwitchRequest) (cloud.GetSwitchResponse, error) {

	return cloud.GetSwitchResponse{}, nil
}

func (p *HuaweiCloud) DescribeSwitches(req cloud.DescribeSwitchesRequest) (cloud.DescribeSwitchesResponse, error) {
	var page int32 = 1
	switches := make([]cloud.Switch, 0, 128)
	for {

		if 1 > page*50 {
			page++
		} else {
			break
		}
	}

	return cloud.DescribeSwitchesResponse{Switches: switches}, nil
}

// CreateSecurityGroup 将VpcId写入Description，方便查找
func (p *HuaweiCloud) CreateSecurityGroup(req cloud.CreateSecurityGroupRequest) (cloud.CreateSecurityGroupResponse, error) {
	request := &model.CreateSecurityGroupRequest{}
	securityGroupOpt := &model.CreateSecurityGroupOption{
		Name:        req.SecurityGroupName,
		Description: &req.VpcId,
	}
	request.Body = &model.CreateSecurityGroupRequestBody{
		SecurityGroup: securityGroupOpt,
	}
	response, err := p.vpcClient.CreateSecurityGroup(request)
	if err != nil {
		return cloud.CreateSecurityGroupResponse{}, err
	}
	if response.HttpStatusCode != 201 {
		return cloud.CreateSecurityGroupResponse{}, fmt.Errorf("httpcode %d, %v", response.HttpStatusCode, response)
	}

	return cloud.CreateSecurityGroupResponse{SecurityGroupId: response.SecurityGroup.Id,
		RequestId: *response.RequestId}, nil
}

// AddIngressSecurityGroupRule 入参各云得统一
func (p *HuaweiCloud) AddIngressSecurityGroupRule(req cloud.AddSecurityGroupRuleRequest) error {
	return p.addSecGrpRule(req, cloud.InSecGroupRule)
}

func (p *HuaweiCloud) AddEgressSecurityGroupRule(req cloud.AddSecurityGroupRuleRequest) error {
	return p.addSecGrpRule(req, cloud.OutSecGroupRule)
}

func (p *HuaweiCloud) DescribeSecurityGroups(req cloud.DescribeSecurityGroupsRequest) (cloud.DescribeSecurityGroupsResponse, error) {
	groups := make([]cloud.SecurityGroup, 0, _pageSize)
	request := &model.ListSecurityGroupsRequest{}
	var listDescription = []string{req.VpcId}
	request.Description = &listDescription
	limitRequest := int32(_pageSize)
	request.Limit = &limitRequest
	markerRequest := ""
	for {
		if markerRequest != "" {
			request.Marker = &markerRequest
		}
		response, err := p.vpcClient.ListSecurityGroups(request)
		if err != nil {
			return cloud.DescribeSecurityGroupsResponse{}, err
		}
		if response.HttpStatusCode != 200 {
			return cloud.DescribeSecurityGroupsResponse{}, fmt.Errorf("httpcode %d, %v", response.HttpStatusCode, response)
		}
		for _, group := range *response.SecurityGroups {
			markerRequest = group.Id
			groups = append(groups, cloud.SecurityGroup{
				SecurityGroupId:   group.Id,
				SecurityGroupType: "normal",
				SecurityGroupName: group.Name,
				CreateAt:          group.CreatedAt.String(),
				VpcId:             req.VpcId,
				RegionId:          req.RegionId,
			})
		}
		if len(*response.SecurityGroups) < _pageSize {
			break
		}
	}
	return cloud.DescribeSecurityGroupsResponse{Groups: groups}, nil
}

func (p *HuaweiCloud) DescribeGroupRules(req cloud.DescribeGroupRulesRequest) (cloud.DescribeGroupRulesResponse, error) {
	rules := make([]cloud.SecurityGroupRule, 0, _pageSize)
	request := &model.ShowSecurityGroupRequest{
		SecurityGroupId: req.SecurityGroupId,
	}
	response, err := p.vpcClient.ShowSecurityGroup(request)
	if err != nil {
		return cloud.DescribeGroupRulesResponse{}, err
	}
	if response.HttpStatusCode != 200 {
		return cloud.DescribeGroupRulesResponse{}, fmt.Errorf("httpcode %d, %v", response.HttpStatusCode, response)
	}

	for _, rule := range response.SecurityGroup.SecurityGroupRules {
		rules = append(rules, cloud.SecurityGroupRule{
			VpcId:           response.SecurityGroup.Description,
			SecurityGroupId: response.SecurityGroup.Id,
			PortRange:       rule.Multiport,
			Protocol:        rule.Protocol,
			Direction:       _secGrpRuleDirection[rule.Direction],
			GroupId:         rule.RemoteGroupId,
			CidrIp:          rule.RemoteIpPrefix,
			PrefixListId:    rule.RemoteAddressGroupId,
			CreateAt:        rule.CreatedAt.String(),
		})
	}

	return cloud.DescribeGroupRulesResponse{Rules: rules}, nil
}

func (p *HuaweiCloud) addSecGrpRule(req cloud.AddSecurityGroupRuleRequest, direction string) error {
	request := &model.CreateSecurityGroupRuleRequest{}
	secGrpRuleOpt := &model.CreateSecurityGroupRuleOption{
		SecurityGroupId: req.SecurityGroupId,
		Direction:       direction,
	}
	if req.IpProtocol != "" {
		secGrpRuleOpt.Protocol = &req.IpProtocol
	}
	if req.PortRange != "" {
		secGrpRuleOpt.Multiport = &req.PortRange
	}
	if req.CidrIp != "" {
		secGrpRuleOpt.RemoteIpPrefix = &req.CidrIp
	}
	if req.GroupId != "" {
		secGrpRuleOpt.RemoteGroupId = &req.GroupId
	}
	if req.PrefixListId != "" {
		secGrpRuleOpt.RemoteAddressGroupId = &req.PrefixListId
	}

	request.Body = &model.CreateSecurityGroupRuleRequestBody{
		SecurityGroupRule: secGrpRuleOpt,
	}
	response, err := p.vpcClient.CreateSecurityGroupRule(request)
	if err != nil {
		return err
	}
	if response.HttpStatusCode != 201 {
		return fmt.Errorf("httpcode %d, %v", response.HttpStatusCode, response)
	}
	return nil
}
