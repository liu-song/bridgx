package cloud

const (
	AlibabaCloud = "AlibabaCloud"
	HuaweiCloud  = "HuaweiCloud"
)

const (
	VPCStatusPending   = "Pending"
	VPCStatusAvailable = "Available"
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
	Building = "Pending"
	Running  = "Running"
	Starting = "Starting"
	Stopping = "Stopping"
	Stopped  = "Stopped"
	Abnormal = "Abnormal"
	Deleted  = "Deleted"
)
