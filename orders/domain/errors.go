package domain

import "fmt"

var ErrorWrongOrderID = fmt.Errorf("wrong order id")
var ErrorOrderIsOnTheWay = fmt.Errorf("order is still on it's way")
var ErrorOrderDelayReportAlreadyExists = fmt.Errorf("order delay report is being processed")
