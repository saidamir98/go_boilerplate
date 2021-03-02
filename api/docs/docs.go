// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Saidamir Botirov",
            "url": "https://www.linkedin.com/in/saidamir-botirov-a08559192",
            "email": "saidamir.botirov@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/config": {
            "get": {
                "description": "shows config of the project only on the development phase",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "gets project config",
                "operationId": "get-config",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.SuccessModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/config.Config"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.ErrorModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "this returns \"pong\" messsage to show service is working",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "returns \"pong\" message",
                "operationId": "ping",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.SuccessModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.ErrorModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/application": {
            "get": {
                "description": "gets application list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "application"
                ],
                "summary": "gets application list",
                "operationId": "get-application-list",
                "parameters": [
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "name": "arrangement",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "id",
                            "body",
                            "created_at",
                            " updated_at"
                        ],
                        "type": "string",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.SuccessModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/application_service.ApplicationListModel"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorModel"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorModel"
                        }
                    }
                }
            },
            "post": {
                "description": "creates an application",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "application"
                ],
                "summary": "creates an application",
                "operationId": "create-application",
                "parameters": [
                    {
                        "description": "application body",
                        "name": "application",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/application_service.CreateApplicationModel"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.SuccessModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/application_service.ApplicationCreatedModel"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorModel"
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.ErrorModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorModel"
                        }
                    }
                }
            }
        },
        "/v1/application/{id}": {
            "get": {
                "description": "gets an application by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "application"
                ],
                "summary": "gets an application by its id",
                "operationId": "get-application-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "application id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.SuccessModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/application_service.ApplicationModel"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorModel"
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.ErrorModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorModel"
                        }
                    }
                }
            },
            "put": {
                "description": "gets an application by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "application"
                ],
                "summary": "gets an application by its id",
                "operationId": "update-application",
                "parameters": [
                    {
                        "type": "string",
                        "description": "application id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "application body",
                        "name": "application",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/application_service.UpdateApplicationModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.SuccessModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/application_service.ApplicationUpdatedModel"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorModel"
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.ErrorModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorModel"
                        }
                    }
                }
            },
            "delete": {
                "description": "deletes an application by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "application"
                ],
                "summary": "deletes an application by its id",
                "operationId": "delete-application",
                "parameters": [
                    {
                        "type": "string",
                        "description": "application id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.SuccessModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/application_service.DeleteApplicationModel"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorModel"
                        }
                    },
                    "422": {
                        "description": "Validation Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.ErrorModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "application_service.ApplicationCreatedModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "application_service.ApplicationListModel": {
            "type": "object",
            "properties": {
                "applications": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/application_service.ApplicationModel"
                    }
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "application_service.ApplicationModel": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "application_service.ApplicationUpdatedModel": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "application_service.CreateApplicationModel": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                }
            }
        },
        "application_service.DeleteApplicationModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "application_service.UpdateApplicationModel": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                }
            }
        },
        "config.Config": {
            "type": "object",
            "properties": {
                "app": {
                    "type": "string"
                },
                "basePath": {
                    "type": "string"
                },
                "defaultLimit": {
                    "type": "string"
                },
                "defaultOffset": {
                    "type": "string"
                },
                "environment": {
                    "description": "development, staging, production",
                    "type": "string"
                },
                "httpport": {
                    "type": "string"
                },
                "logLevel": {
                    "description": "debug, info, warn, error, dpanic, panic, fatal",
                    "type": "string"
                },
                "postgresDatabase": {
                    "type": "string"
                },
                "postgresHost": {
                    "type": "string"
                },
                "postgresPassword": {
                    "type": "string"
                },
                "postgresPort": {
                    "type": "integer"
                },
                "postgresUser": {
                    "type": "string"
                },
                "rabbitURI": {
                    "type": "string"
                },
                "serviceHost": {
                    "type": "string"
                }
            }
        },
        "response.ErrorModel": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "response.SuccessModel": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Go Boilerplate API",
	Description: "This is a Go Boilerplate for medium sized projects",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
