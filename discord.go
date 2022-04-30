package emotebox

import "net/url"

type DiscordSession interface {
	GetUser(userId string) (*User, error)                   // Get User by Id
	GetAvatar(userId string, size string) (*url.URL, error) // Get Avtar from user id
	CurrentUser() (*User, error)                            // Gets the current authenticated user
}
