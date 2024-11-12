package domain

import "github.com/ei-sugimoto/datapuppy/internal"

type Span struct {
	ID           int
	HostID       int
	DurationTime int
}

func NewSpan(hostID, durationTime int) *Span {
	return &Span{
		HostID:       hostID,
		DurationTime: durationTime,
	}
}

func ValidateSpan(hostID, durationTime int) error {
	if hostID <= 0 {
		return internal.ErrHostIDMustBeGreaterThanZero
	}
	if durationTime < 0 {
		return internal.ErrDurationTimeMustBeGreaterThanZero
	}
	return nil
}
