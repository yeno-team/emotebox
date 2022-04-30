package discordgo

import (
	"net/url"

	"github.com/yeno-team/emotebox"
)

type UserService struct {
	Session emotebox.DiscordSession
}

func (s *UserService) GetCurrentUser() (*emotebox.User, error) {
	user, err := s.Session.GetUser("@me")
	if err != nil {
		return nil, err
	}

	return &emotebox.User{Id: user.Id, Username: user.Username, AvatarUrl: user.AvatarUrl}, nil
}

func (s *UserService) GetAvatar(userId string, size string) (*url.URL, error) {
	return s.Session.GetAvatar(userId, size)
}
