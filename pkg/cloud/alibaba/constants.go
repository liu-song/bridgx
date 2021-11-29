package alibaba

import "github.com/galaxy-future/BridgX/pkg/cloud"

const (
	_subOrderNumPerMain    = 3
	_maxNumEcsPerOperation = 100
)

var _chargeType = map[string]string{
	"Subscription": cloud.PrePaid,
	"PayAsYouGo":   cloud.PostPaid,
}

var _payStatus = map[string]int8{
	"Paid":      cloud.Paid,
	"Unpaid":    cloud.Unpaid,
	"Cancelled": cloud.Cancelled,
}

var _ecsStatus = map[string]string{
	"Pending":  cloud.Building,
	"Running":  cloud.Running,
	"Starting": cloud.Starting,
	"Stopping": cloud.Stopping,
	"Stopped":  cloud.Stopped,
}

var _secGrpRuleDirection = map[string]string{
	"ingress": cloud.InSecGroupRule,
	"egress":  cloud.OutSecGroupRule,
}
