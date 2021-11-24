package huawei

import (
	"fmt"

	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	ecsRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	iam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
	iamModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	iamRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
	ims "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2"
	imsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2/model"
	imsRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2/region"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3"
	vpcRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/region"
)

type HuaweiCloud struct {
	ecsClient *ecs.EcsClient
	imsClient *ims.ImsClient
	vpcClient *vpc.VpcClient
	iamClient *iam.IamClient
}

func New(ak, sk, regionId string) (*HuaweiCloud, error) {
	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	ecsClt := ecs.NewEcsClient(
		ecs.EcsClientBuilder().
			WithRegion(ecsRegion.ValueOf(regionId)).
			WithCredential(auth).
			Build())
	imsClt := ims.NewImsClient(
		ims.ImsClientBuilder().
			WithRegion(imsRegion.ValueOf(regionId)).
			WithCredential(auth).
			Build())
	vpcClt := vpc.NewVpcClient(
		vpc.VpcClientBuilder().
			WithRegion(vpcRegion.ValueOf(regionId)).
			WithCredential(auth).
			Build())

	gAuth := global.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	iamClt := iam.NewIamClient(
		iam.IamClientBuilder().
			WithRegion(iamRegion.ValueOf(regionId)).
			WithCredential(gAuth).
			Build())
	return &HuaweiCloud{ecsClient: ecsClt, imsClient: imsClt, vpcClient: vpcClt, iamClient: iamClt}, nil
}

func (HuaweiCloud) ProviderType() string {
	return cloud.HuaweiCloud
}

func (p *HuaweiCloud) CreateVPC(req cloud.CreateVpcRequest) (cloud.CreateVpcResponse, error) {

	return cloud.CreateVpcResponse{}, nil
}

func (p *HuaweiCloud) GetVPC(req cloud.GetVpcRequest) (cloud.GetVpcResponse, error) {

	return cloud.GetVpcResponse{}, nil
}

func (p HuaweiCloud) DescribeVpcs(req cloud.DescribeVpcsRequest) (cloud.DescribeVpcsResponse, error) {
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

func (p HuaweiCloud) DescribeSwitches(req cloud.DescribeSwitchesRequest) (cloud.DescribeSwitchesResponse, error) {
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

func (p *HuaweiCloud) CreateSecurityGroup(req cloud.CreateSecurityGroupRequest) (cloud.CreateSecurityGroupResponse, error) {

	return cloud.CreateSecurityGroupResponse{}, nil
}

func (p HuaweiCloud) AddIngressSecurityGroupRule(req cloud.AddSecurityGroupRuleRequest) error {

	return nil
}

func (p HuaweiCloud) AddEgressSecurityGroupRule(req cloud.AddSecurityGroupRuleRequest) error {

	return nil
}

func (p HuaweiCloud) DescribeSecurityGroups(req cloud.DescribeSecurityGroupsRequest) (cloud.DescribeSecurityGroupsResponse, error) {
	return cloud.DescribeSecurityGroupsResponse{}, nil
}

func (p *HuaweiCloud) DescribeGroupRules(req cloud.DescribeGroupRulesRequest) (cloud.DescribeGroupRulesResponse, error) {
	rules := make([]cloud.SecurityGroupRule, 0, 128)

	return cloud.DescribeGroupRulesResponse{Rules: rules}, nil
}

// GetRegions 暂时返回中文名字
func (p *HuaweiCloud) GetRegions() (cloud.GetRegionsResponse, error) {
	request := &iamModel.KeystoneListRegionsRequest{}
	response, err := p.iamClient.KeystoneListRegions(request)
	if err != nil {
		return cloud.GetRegionsResponse{}, err
	}
	if response.HttpStatusCode != 200 {
		return cloud.GetRegionsResponse{}, fmt.Errorf("httpcode %d, %v", response.HttpStatusCode, response)
	}

	regions := make([]cloud.Region, 0, len(*response.Regions))
	for _, region := range *response.Regions {
		regions = append(regions, cloud.Region{
			RegionId:  region.Id,
			LocalName: region.Locales.ZhCn,
		})
	}
	return cloud.GetRegionsResponse{Regions: regions}, nil
}

// DescribeImages osType转成字符串
func (p *HuaweiCloud) DescribeImages(req cloud.DescribeImagesRequest) (cloud.DescribeImagesResponse, error) {
	request := &imsModel.ListImagesRequest{}
	limitRequest := int32(_pageSize)
	request.Limit = &limitRequest
	markerRequest := ""
	images := make([]cloud.Image, 0, _pageSize)
	for {
		request.Marker = &markerRequest
		response, err := p.imsClient.ListImages(request)
		if err != nil {
			return cloud.DescribeImagesResponse{}, err
		}
		if response.HttpStatusCode != 200 {
			return cloud.DescribeImagesResponse{}, fmt.Errorf("httpcode %d, %v", response.HttpStatusCode, response)
		}

		for _, img := range *response.Images {
			markerRequest = img.Id
			osType, _ := img.OsType.MarshalJSON()
			images = append(images, cloud.Image{
				ImageId: img.Id,
				OsType:  string(osType),
				OsName:  *img.OsVersion,
			})
		}
		if len(*response.Images) < _pageSize {
			break
		}
	}
	return cloud.DescribeImagesResponse{Images: images}, nil
}

func (p *HuaweiCloud) GetOrders(req cloud.GetOrdersRequest) (cloud.GetOrdersResponse, error) {

	orders := make([]cloud.Order, 0, 1)

	return cloud.GetOrdersResponse{Orders: orders}, nil
}
