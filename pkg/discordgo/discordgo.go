/**
 * Copyright (C) 2022 Yeno
 *
 * This file is part of Emotebox.
 *
 * Emotebox is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Emotebox is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Emotebox.  If not, see <http://www.gnu.org/licenses/>.
 */

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
