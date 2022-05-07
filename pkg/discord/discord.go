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

package discord

import "golang.org/x/oauth2"

// Endpoint is Discord's OAuth 2.0 endpoint
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://discord.com/api/oauth2/authorize",
	TokenURL: "https://discord.com/api/oauth2/token",
}

// Scopes is Discord's available scopes for OAuth2
// For more information, https://discord.com/developers/docs/topics/oauth2#shared-resources-oauth2-scopes
var Scopes []string = []string{
	"activites.read",
	"activities.write",
	"applications.builds.read",
	"applications.builds.upload",
	"applications.commands",
	"applications.commands.update",
	"applications.commands.permissions.update",
	"applications.entitlements",
	"applications.store.update",
	"bot",
	"connections",
	"email",
	"gdm.join",
	"guilds",
	"guilds.join",
	"guilds.members.read",
	"identify",
	"messages.read",
	"relationships.read",
	"rpc",
	"rpc.activities.write",
	"rpc.notifications.read",
	"rpc.voice.read",
	"rpc.voice.write",
	"webhook.incoming",
}
