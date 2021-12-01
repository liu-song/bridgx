package huawei

import (
	"github.com/galaxy-future/BridgX/pkg/cloud"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
)

const (
	_maxNumEcsPerOperation = 1000
	_pageSize              = 1000
)

var _rootDiskCategory = map[string]model.PrePaidServerRootVolumeVolumetype{
	"SATA":  model.GetPrePaidServerRootVolumeVolumetypeEnum().SATA,
	"SAS":   model.GetPrePaidServerRootVolumeVolumetypeEnum().SAS,
	"SSD":   model.GetPrePaidServerRootVolumeVolumetypeEnum().SSD,
	"GPSSD": model.GetPrePaidServerRootVolumeVolumetypeEnum().GPSSD,
	"CO_P1": model.GetPrePaidServerRootVolumeVolumetypeEnum().CO_P1,
	"UH_L1": model.GetPrePaidServerRootVolumeVolumetypeEnum().UH_L1,
}

var _dataDiskCategory = map[string]model.PrePaidServerDataVolumeVolumetype{
	"SATA":  model.GetPrePaidServerDataVolumeVolumetypeEnum().SATA,
	"SAS":   model.GetPrePaidServerDataVolumeVolumetypeEnum().SAS,
	"SSD":   model.GetPrePaidServerDataVolumeVolumetypeEnum().SSD,
	"GPSSD": model.GetPrePaidServerDataVolumeVolumetypeEnum().GPSSD,
	"CO_P1": model.GetPrePaidServerDataVolumeVolumetypeEnum().CO_P1,
	"UH_L1": model.GetPrePaidServerDataVolumeVolumetypeEnum().UH_L1,
}

var _ecsChargeType = map[string]string{
	"0": cloud.PostPaid,
	"1": cloud.PrePaid,
}

var _ecsStatus = map[string]string{
	"BUILD":         cloud.EcsBuilding,
	"REBUILD":       cloud.EcsBuilding,
	"REBOOT":        cloud.EcsStarting,
	"HARD_REBOOT":   cloud.EcsStarting,
	"RESIZE":        cloud.EcsStarting,
	"REVERT_RESIZE": cloud.EcsStarting,
	"VERIFY_RESIZE": cloud.EcsStarting,
	"MIGRATING":     cloud.EcsRunning,
	"ACTIVE":        cloud.EcsRunning,
	"SHUTOFF":       cloud.EcsStopped,
	"ERROR":         cloud.EcsAbnormal,
	"DELETED":       cloud.EcsDeleted,
}

var _secGrpRuleDirection = map[string]string{
	"ingress": cloud.SecGroupRuleIn,
	"egress":  cloud.SecGroupRuleOut,
}

var _osType = map[string]string{
	"\"Linux\"\n":   cloud.OsLinux,
	"\"Windows\"\n": cloud.OsWindows,
	"\"Other\"\n":   cloud.OsOther,
}

var _vpcStatus = map[string]string{
	"\"CREATING\"\n": cloud.VPCStatusPending,
	"\"OK\"\n":       cloud.VPCStatusAvailable,
	"\"ERROR\"\n":    cloud.VPCStatusAbnormal,
}

var _subnetStatus = map[string]string{
	"\"UNKNOWN\"\n": cloud.SubnetPending,
	"\"ACTIVE\"\n":  cloud.SubnetAvailable,
	"\"ERROR\"\n":   cloud.SubnetAbnormal,
}
