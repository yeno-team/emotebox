// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/yeno-team/emotebox/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/status": {
            "get": {
                "description": "get status of app",
                "tags": [
                    "Status"
                ],
                "summary": "Gives the default status of the app",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/emotebox.Status"
                        }
                    }
                }
            }
        },
        "/api/v1/user": {
            "get": {
                "description": "Gets the current authenticated user",
                "tags": [
                    "Authentication"
                ],
                "summary": "Gets the current authenticated user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/emotebox.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "emotebox.Status": {
            "description": "Default Status Object for presentation",
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "message": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 4,
                    "example": "OK"
                }
            }
        },
        "emotebox.User": {
            "description": "The user object",
            "type": "object",
            "properties": {
                "avatar_url": {
                    "description": "User's Avatar Url",
                    "type": "string",
                    "example": "https://cdn.discordapp.com/avatars/473740736526024714/a_721f35ba7b5ecca6317b93f893f66908.gif?size=80"
                },
                "id": {
                    "description": "User's Discord Id",
                    "type": "string",
                    "example": "473740736526024714"
                },
                "username": {
                    "description": "User's Name",
                    "type": "string",
                    "example": "Khai"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Emotebox API",
	Description:      "Emotebox's Server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
