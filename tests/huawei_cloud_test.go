package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/galaxy-future/BridgX/pkg/cloud/huawei"
)

func getClient() (*huawei.HuaweiCloud, error) {
	client, err := huawei.New("ak", "sk", "region")
	if err != nil {
		return nil, err
	}
	return client, nil
}

func TestCreateServers(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Log(err)
		return
	}

	param := cloud.Params{
		InstanceType: "c6s.large.2",
		ImageId:      "75e56468-4f96-4e87-927e-35081d13fd79",
		Network: &cloud.Network{
			VpcId:                   "dd57a464-b590-466e-b572-8fe19fe7d67f",
			SubnetId:                "3719e837-9897-43bc-819a-d1c9e60ab72f",
			SecurityGroup:           "",
			InternetChargeType:      "traffic",
			InternetMaxBandwidthOut: 0,
			InternetIpType:          "5_bgp",
		},
		Disks: &cloud.Disks{
			SystemDisk: cloud.DiskConf{Size: 40, Category: "SSD"},
			DataDisk:   []cloud.DiskConf{},
		},
		Password: "ASDqwe123",
		Tags: []cloud.Tag{
			{
				Key:   cloud.TaskId,
				Value: "12345",
			},
			{
				Key:   cloud.ClusterName,
				Value: "cluster2",
			},
		},
	}
	res, err := client.BatchCreate(param, 1)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(res)
}

func TestShowIns(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Log(err)
		return
	}

	var res interface{}
	var resStr []byte
	ids := []string{""}
	res, err = client.GetInstances(ids)
	if err != nil {
		t.Log(err)
		return
	}
	resStr, _ = json.Marshal(res)
	t.Log(string(resStr))

	tags := []cloud.Tag{{Key: cloud.TaskId, Value: "12345"}}
	res, err = client.GetInstancesByTags("", tags)
	if err != nil {
		t.Log(err)
		return
	}
	resStr, _ = json.Marshal(res)
	t.Log(string(resStr))
}

func TestCtlIns(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Log(err)
		return
	}

	ids := []string{""}

	err = client.StopInstances(ids)
	if err != nil {
		t.Log(err.Error())
	}
	time.Sleep(time.Duration(60) * time.Second)

	err = client.StartInstances(ids)
	if err != nil {
		t.Log(err.Error())
	}
	time.Sleep(time.Duration(60) * time.Second)

	err = client.BatchDelete(ids, "")
	if err != nil {
		t.Log(err.Error())
	}
}

func TestGetResource(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Log(err)
		return
	}

	var res interface{}
	var resStr []byte
	res, err = client.GetRegions()
	if err != nil {
		t.Log(err.Error())
		return
	}
	resStr, _ = json.Marshal(res)
	t.Log(string(resStr))

	res, err = client.GetZones(cloud.GetZonesRequest{})
	if err != nil {
		t.Log(err.Error())
		return
	}
	resStr, _ = json.Marshal(res)
	t.Log(string(resStr))

	res, err = client.DescribeAvailableResource(cloud.DescribeAvailableResourceRequest{})
	if err != nil {
		t.Log(err.Error())
		return
	}
	resStr, _ = json.Marshal(res)
	t.Log(string(resStr))

	res, err = client.DescribeInstanceTypes(cloud.DescribeInstanceTypesRequest{TypeName: []string{"1"}})
	if err != nil {
		t.Log(err.Error())
		return
	}
	resStr, _ = json.Marshal(res)
	t.Log(string(resStr))

	res, err = client.DescribeImages(cloud.DescribeImagesRequest{FlavorId: "c6s.large.2"})
	if err != nil {
		t.Log(err.Error())
		return
	}
	resStr, _ = json.Marshal(res)
	t.Log(string(resStr))
}

func TestCreateSecGrp(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Log(err)
		return
	}

	req := cloud.CreateSecurityGroupRequest{
		SecurityGroupName: "test2",
		VpcId:             "dd57a464-b590-466e-b572-8fe19fe7d67f",
	}
	res, err := client.CreateSecurityGroup(req)
	if err != nil {
		t.Log(err.Error())
		return
	}
	resStr, _ := json.Marshal(res)
	t.Log(string(resStr))
}

func TestAddSecGrpRule(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Log(err)
		return
	}

	req := cloud.AddSecurityGroupRuleRequest{
		SecurityGroupId: "347040a4-1ace-454f-b6b3-320a76b334a4",
		IpProtocol:      "udp",
		PortRange:       "8894",
		CidrIp:          "192.168.1.1/24",
	}
	err = client.AddIngressSecurityGroupRule(req)
	if err != nil {
		t.Log(err.Error())
		return
	}

	req = cloud.AddSecurityGroupRuleRequest{
		SecurityGroupId: "347040a4-1ace-454f-b6b3-320a76b334a4",
		IpProtocol:      "tcp",
		PortRange:       "1000",
		CidrIp:          "192.168.1.1/24",
	}
	err = client.AddEgressSecurityGroupRule(req)
	if err != nil {
		t.Log(err.Error())
		return
	}
}

func TestShowSecGrp(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Log(err)
		return
	}

	var res interface{}
	var resStr []byte
	res, err = client.DescribeSecurityGroups(cloud.DescribeSecurityGroupsRequest{
		VpcId: "dd57a464-b590-466e-b572-8fe19fe7d67f",
	})
	if err != nil {
		t.Log(err.Error())
		return
	}
	resStr, _ = json.Marshal(res)
	t.Log(string(resStr))

	res, err = client.DescribeGroupRules(cloud.DescribeGroupRulesRequest{
		SecurityGroupId: "347040a4-1ace-454f-b6b3-320a76b334a4",
	})
	if err != nil {
		t.Log(err.Error())
		return
	}
	resStr, _ = json.Marshal(res)
	t.Log(string(resStr))
}
