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
