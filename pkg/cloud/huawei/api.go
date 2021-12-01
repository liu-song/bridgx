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
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
	vpcRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/region"
	secGrp "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3"
	secGrpRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/region"
)

type HuaweiCloud struct {
	ecsClient    *ecs.EcsClient
	imsClient    *ims.ImsClient
	secGrpClient *secGrp.VpcClient
	vpcClient    *vpc.VpcClient
	iamClient    *iam.IamClient
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
	secGrpClt := secGrp.NewVpcClient(
		secGrp.VpcClientBuilder().
			WithRegion(secGrpRegion.ValueOf(regionId)).
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
	return &HuaweiCloud{ecsClient: ecsClt, imsClient: imsClt, secGrpClient: secGrpClt, vpcClient: vpcClt, iamClient: iamClt}, nil
}

func (HuaweiCloud) ProviderType() string {
	return cloud.HuaweiCloud
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

// DescribeImages osType转成字符串;返回太多了
func (p *HuaweiCloud) DescribeImages(req cloud.DescribeImagesRequest) (cloud.DescribeImagesResponse, error) {
	pageSize := 500
	images := make([]cloud.Image, 0, pageSize)
	request := &imsModel.ListImagesRequest{}
	sortDirRequest := imsModel.GetListImagesRequestSortDirEnum().DESC
	request.SortDir = &sortDirRequest
	sortKeyRequest := imsModel.GetListImagesRequestSortKeyEnum().NAME
	request.SortKey = &sortKeyRequest
	statusRequest := imsModel.GetListImagesRequestStatusEnum().ACTIVE
	request.Status = &statusRequest
	if req.FlavorId != "" {
		request.FlavorId = &req.FlavorId
	}
	limitRequest := int32(pageSize)
	request.Limit = &limitRequest
	markerRequest := ""
	for {
		if markerRequest != "" {
			request.Marker = &markerRequest
		}
		response, err := p.imsClient.ListImages(request)
		if err != nil {
			return cloud.DescribeImagesResponse{}, err
		}
		if response.HttpStatusCode != 200 {
			return cloud.DescribeImagesResponse{}, fmt.Errorf("httpcode %d, %v", response.HttpStatusCode, response)
		}

		for _, img := range *response.Images {
			osType, _ := img.OsType.MarshalJSON()
			images = append(images, cloud.Image{
				ImageId: img.Id,
				OsType:  _osType[string(osType)],
				OsName:  *img.OsVersion,
			})
		}
		imgNum := len(*response.Images)
		if imgNum < pageSize {
			break
		}
		markerRequest = (*response.Images)[imgNum-1].Id
	}
	return cloud.DescribeImagesResponse{Images: images}, nil
}

func (p *HuaweiCloud) GetOrders(req cloud.GetOrdersRequest) (cloud.GetOrdersResponse, error) {
	orders := make([]cloud.Order, 0, 1)

	return cloud.GetOrdersResponse{Orders: orders}, nil
}
