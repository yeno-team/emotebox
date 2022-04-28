package example

import "github.com/yeno-team/emotebox"

type StatusService struct{}

func (s *StatusService) Status() (*emotebox.Status, error) {
	return &emotebox.Status{Message: "OK", Code: 200}, nil
}
