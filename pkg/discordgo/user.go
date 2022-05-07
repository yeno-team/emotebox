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
