package discordgo

import (
	"net/url"

	_discordgo "github.com/bwmarrin/discordgo"
	"github.com/yeno-team/emotebox"
)

func New(token string) (*DiscordSession, error) {
	discordgoSession, err := _discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	return &DiscordSession{Session: discordgoSession}, nil
}

type DiscordSession struct {
	Session *_discordgo.Session
}

func (s *DiscordSession) GetUser(userId string) (*emotebox.User, error) {
	user, err := s.Session.User(userId)
	if err != nil {
		return nil, err
	}

	return &emotebox.User{Id: user.ID, Username: user.Username, AvatarUrl: user.AvatarURL("")}, nil
}

func (s *DiscordSession) GetAvatar(userId string, size string) (*url.URL, error) {
	user, err := s.Session.User(userId)
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(user.AvatarURL(size))
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (s *DiscordSession) CurrentUser() (*emotebox.User, error) {
	return s.GetUser("@me")
}
