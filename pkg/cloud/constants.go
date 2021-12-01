package cloud

const (
	AlibabaCloud = "AlibabaCloud"
	HuaweiCloud  = "HuaweiCloud"
)

const (
	TaskId      = "TaskId"
	ClusterName = "ClusterName"
)

const (
	PrePaid  = "PrePaid"
	PostPaid = "PostPaid"
)

const (
	Paid = iota + 1
	Unpaid
	Cancelled
)

const (
	EcsBuilding = "Pending"
	EcsRunning  = "Running"
	EcsStarting = "Starting"
	EcsStopping = "Stopping"
	EcsStopped  = "Stopped"
	EcsAbnormal = "Abnormal"
	EcsDeleted  = "Deleted"
)

const (
	SecGroupRuleIn  = "ingress"
	SecGroupRuleOut = "egress"
)

const (
	OsLinux   = "linux"
	OsWindows = "windows"
	OsOther   = "other"
)

const (
	VPCStatusPending   = "Pending"
	VPCStatusAvailable = "Available"
	VPCStatusAbnormal  = "abnormal"
)

const (
	SubnetPending   = "Pending"
	SubnetAvailable = "Available"
	SubnetAbnormal  = "abnormal"
)
