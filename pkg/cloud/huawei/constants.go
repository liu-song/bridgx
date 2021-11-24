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
	"BUILD":         cloud.Building,
	"REBUILD":       cloud.Building,
	"REBOOT":        cloud.Starting,
	"HARD_REBOOT":   cloud.Starting,
	"RESIZE":        cloud.Starting,
	"REVERT_RESIZE": cloud.Starting,
	"VERIFY_RESIZE": cloud.Starting,
	"MIGRATING":     cloud.Running,
	"ACTIVE":        cloud.Running,
	"SHUTOFF":       cloud.Stopped,
	"ERROR":         cloud.Abnormal,
	"DELETED":       cloud.Deleted,
}
