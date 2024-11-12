package internal

import "fmt"

var ErrNameIsRequired = fmt.Errorf("Name is required")
var ErrPortMustBeGreaterThanZero = fmt.Errorf("Port must be greater than 0")
var ErrHostIDMustBeGreaterThanZero = fmt.Errorf("HostID must be greater than 0")
var ErrDurationTimeMustBeGreaterThanZero = fmt.Errorf("DurationTime must be greater than 0")
