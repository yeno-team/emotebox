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

// Package oauth2 acts as a fascade to the external package oauth2
package oauth2

import (
	"net/url"

	"github.com/yeno-team/emotebox/pkg/random"
	"golang.org/x/oauth2"
)

// Endpoint is an object that holds a third-party application's OAuth2 Endpoints
type Endpoint struct {
	TokenURL string `json:"tokenURL"`
	AuthURL  string `json:"authURL"`
}

// OAuth2Client is the external "oauth2" implementation of the interface internal.OAuth2Client
type OAuth2Client struct {
	ClientID     string   `json:"clientId"`
	ClientSecret string   `json:"clientSecret"`
	Scopes       []string `json:"scopes"`
	Endpoint     Endpoint
}

// AuthURL is an implementation of the AuthURL from internal.OAuth2Client
func (c *OAuth2Client) AuthURL(redirectUrl url.URL) (*url.URL, error) {
	endpoint := oauth2.Endpoint{
		TokenURL: c.Endpoint.TokenURL,
		AuthURL:  c.Endpoint.AuthURL,
	}

	conf := &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		Scopes:       c.Scopes,
		Endpoint:     endpoint,
		RedirectURL:  redirectUrl.String(),
	}

	s, err := random.RandString(32)
	if err != nil {
		return nil, err
	}

	return url.Parse(conf.AuthCodeURL(s))
}
