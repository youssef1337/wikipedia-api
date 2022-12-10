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
        "contact": {
            "name": "Youssef Sobhy",
            "email": "youssefsobhy22@gmail.com"
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
        "/api": {
            "get": {
                "description": "Check if the API is operational.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Check if the API is operational.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.SuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/search": {
            "get": {
                "description": "Search for a short description of a person, place, or thing.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Search for a short description of a person, place, or thing.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The name of the person, place, or thing you want to search for.",
                        "name": "query",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal.InternalServerErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "internal.Data": {
            "type": "object",
            "properties": {
                "short_description": {
                    "type": "string",
                    "example": "A short description of the person, place, or thing you searched for."
                }
            }
        },
        "internal.ErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/internal.HTTPError"
                    }
                },
                "status": {
                    "type": "string",
                    "example": "error"
                }
            }
        },
        "internal.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "detail": {
                    "type": "string",
                    "example": "Query parameter is required"
                },
                "request_id": {
                    "type": "string",
                    "example": "f7a4c0c0-5b5e-4b4c-9c1f-1b5c1b5c1b5c"
                }
            }
        },
        "internal.InternalServerError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "detail": {
                    "type": "string",
                    "example": "An internal server error occurred. Please contact the developer at youssefsobhy22@gmail.com and provide the request ID."
                },
                "request_id": {
                    "type": "string",
                    "example": "f7a4c0c0-5b5e-4b4c-9c1f-1b5c1b5c1b5c"
                }
            }
        },
        "internal.InternalServerErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/internal.InternalServerError"
                    }
                },
                "status": {
                    "type": "string",
                    "example": "error"
                }
            }
        },
        "internal.SuccessResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/internal.Data"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "wikipedia-api.youssefsobhy.com",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "Wikipedia API",
	Description:      "A simple API to get the short description of a wikipedia article.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
