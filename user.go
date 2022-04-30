package emotebox

import (
	"net/url"
)

// Authenticated User Object
// @Description The user object
type User struct {
	Id        string `json:"id"`         // User's Discord Id
	Username  string `json:"username"`   // User's Name
	AvatarUrl string `json:"avatar_url"` // User's Avatar Url
}

type UserService interface {
	GetCurrentUser() (*User, error)
	GetAvatar(userId string, size string) (*url.URL, error)
}
