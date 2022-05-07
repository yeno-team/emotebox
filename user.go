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

package emotebox

import (
	"net/url"
)

// Authenticated User Object
// @Description The user object
type User struct {
	Id        string `json:"id" example:"473740736526024714"`                                                                                           // User's Discord Id
	Username  string `json:"username" example:"Khai"`                                                                                                   // User's Name
	AvatarUrl string `json:"avatar_url" example:"https://cdn.discordapp.com/avatars/473740736526024714/a_721f35ba7b5ecca6317b93f893f66908.gif?size=80"` // User's Avatar Url
}

type UserService interface {
	GetCurrentUser() (*User, error)
	GetAvatar(userId string, size string) (*url.URL, error)
}
