// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Adam Drew",
            "url": "https://github.com/RedHatInsights/consoledot-go-starter-app",
            "email": "addrew@redhat.com"
        },
        "license": {
            "name": "MIT License",
            "url": "https://opensource.org/license/mit/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/starter-app/v1/hello": {
            "get": {
                "description": "Recieve a greeting from the API",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "Recieve a greeting",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/healthz": {
            "get": {
                "description": "Determines readiness of the application",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "probes"
                ],
                "summary": "Determines readiness of the application",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/livez": {
            "get": {
                "description": "Determines if application is still alive",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "probes"
                ],
                "summary": "Determines if application is still alive",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "ConsoleDot Go Starter App API",
	Description:      "This is a sample API for the ConsoleDot Go Starter App.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
