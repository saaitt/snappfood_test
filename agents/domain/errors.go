package domain

import "fmt"

var ErrorAgentIsAlreadyAssigned = fmt.Errorf("agent already has an active delay report")
var ErrorNoDelayReportIsOnQueue = fmt.Errorf("no delay report is on queue ")
