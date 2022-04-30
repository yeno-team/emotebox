package discordgo_test

import (
	"errors"
	"net/url"
	"testing"

	"github.com/yeno-team/emotebox"
	"github.com/yeno-team/emotebox/mock"
	. "github.com/yeno-team/emotebox/pkg/discordgo"
)

func TestGetAvatar(t *testing.T) {
	expectedUrl, err := url.Parse("http://localhost.com")
	if err != nil {
		t.Error(err)
	}

	mockSession := &mock.DiscordSession{
		GetAvatarFn: func(userId string) (*url.URL, error) {
			return expectedUrl, nil
		},
	}
	s := UserService{Session: mockSession}

	avatar, err := s.GetAvatar("1", "")
	if err != nil {
		t.Error(err)
	}

	if avatar != expectedUrl {
		t.Errorf("Expected %s to equal %s", avatar.String(), expectedUrl.String())
	}
}

func TestCurrentUser(t *testing.T) {
	expectedUser := &emotebox.User{Id: "1", AvatarUrl: "http://localhost", Username: "testUser"}

	mockSession := &mock.DiscordSession{
		GetUserFn: func(userId string) (*emotebox.User, error) {
			if userId != "@me" {
				return nil, errors.New("Expected GetCurrentUser to call GetUser with parmeter '@me'")
			}
			return expectedUser, nil
		},
	}

	s := UserService{Session: mockSession}

	user, err := s.GetCurrentUser()
	if err != nil {
		t.Error(err)
	}

	if user.Id != expectedUser.Id || user.AvatarUrl != expectedUser.AvatarUrl || user.Username != expectedUser.Username {
		t.Errorf("Expected %v to equal %v", user, expectedUser)
	}
}
