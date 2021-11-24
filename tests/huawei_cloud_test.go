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

	ids := []string{""}
	res, err := client.GetInstances(ids)
	if err != nil {
		t.Log(err)
		return
	}
	resStr, _ := json.Marshal(res)
	t.Log(string(resStr))

	tags := []cloud.Tag{{Key: cloud.TaskId, Value: "12345"}}
	res1, err1 := client.GetInstancesByTags("", tags)
	if err1 != nil {
		t.Log(err1)
		return
	}
	resStr1, _ := json.Marshal(res1)
	t.Log(string(resStr1))
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
