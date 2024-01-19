package domain

import "fmt"

var ErrorAgentIsAlreadyAssigned = fmt.Errorf("agent already has an active delay report")
