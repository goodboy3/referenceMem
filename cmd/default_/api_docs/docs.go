// Package api_docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package api_docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://domain.com",
        "contact": {
            "name": "Support",
            "url": "https://domain.com",
            "email": "contact@domain.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/health": {
            "get": {
                "description": "health check",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "health check",
                "responses": {
                    "200": {
                        "description": "result",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/echoServer.RespBody"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "integer"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "echoServer.RespBody": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                },
                "result": {},
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:80",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "api example",
	Description:      "api example",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
