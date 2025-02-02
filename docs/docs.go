// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/currency/add/{coin}": {
            "get": {
                "description": "adds currency for observing",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "currency"
                ],
                "summary": "adds currency for observing",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Coin Name",
                        "name": "coin",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AddCurrency"
                        }
                    }
                }
            }
        },
        "/api/v1/currency/price/{coin}/{timestamp}": {
            "get": {
                "description": "get currency price by timestamp",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "currency"
                ],
                "summary": "get currency price",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Coin Name",
                        "name": "coin",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Timestamp",
                        "name": "timestamp",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.GetCurrencyPrice"
                        }
                    }
                }
            }
        },
        "/api/v1/currency/remove/{coin}": {
            "get": {
                "description": "get recommendations by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "currency"
                ],
                "summary": "delete currency from observing",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Coin Name",
                        "name": "coin",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.DeleteCurrency"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.AddCurrency": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "response.DeleteCurrency": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "response.GetCurrencyPrice": {
            "type": "object",
            "properties": {
                "coin": {
                    "type": "string"
                },
                "status": {
                    "type": "object",
                    "additionalProperties": true
                },
                "timestamp": {
                    "type": "integer"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Тестовое задание",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
