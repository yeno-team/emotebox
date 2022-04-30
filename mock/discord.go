package mock

import (
	"net/url"

	"github.com/yeno-team/emotebox"
)

type DiscordSession struct {
	GetUserFn            func(userId string) (*emotebox.User, error)
	GetUserFnInvoked     bool
	GetAvatarFn          func(userId string) (*url.URL, error)
	GetAvatarFnInvoked   bool
	CurrentUserFn        func() (*emotebox.User, error)
	CurrentUserFnInvoked bool
}

func (s *DiscordSession) GetUser(userId string) (*emotebox.User, error) {
	s.GetUserFnInvoked = true
	return s.GetUserFn(userId)
}

func (s *DiscordSession) GetAvatar(userId string, size string) (*url.URL, error) {
	s.GetAvatarFnInvoked = true
	return s.GetAvatarFn(userId)
}

func (s *DiscordSession) CurrentUser() (*emotebox.User, error) {
	s.CurrentUserFnInvoked = true
	return s.CurrentUserFn()
}
