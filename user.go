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
