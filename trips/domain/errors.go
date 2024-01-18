package domain

import "fmt"

var ErrorOrderIsAlreadyAssigned = fmt.Errorf("order is still on it's way")
var TripNotFound = fmt.Errorf("no trip found with given id")

var ErrorTripStatusCanOnlyChangeToAtVendor = fmt.Errorf("trip can only change to AT_VENDOR")
var ErrorTripStatusCanOnlyChangeToPicked = fmt.Errorf("trip can only change to PICKED")
var ErrorTripStatusCanOnlyChangeToDelivered = fmt.Errorf("trip can only change to Delivered")
