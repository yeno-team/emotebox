package mock

import "github.com/yeno-team/emotebox"

// Mock implmentation of emotebox.Status
type StatusService struct {
	StatusFn        func() (*emotebox.Status, error)
	StatusFnInvoked bool
}

// Marks mock status function as invoked
func (s *StatusService) Status() (*emotebox.Status, error) {
	s.StatusFnInvoked = true
	return s.StatusFn()
}
