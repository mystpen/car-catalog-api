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
        "/cars": {
            "get": {
                "description": "listing cars data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name search by regNum",
                        "name": "regNum",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "name search by mark",
                        "name": "mark",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "name search by model",
                        "name": "model",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search by year",
                        "name": "year",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Cars"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.ErrRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrRes"
                        }
                    }
                }
            },
            "post": {
                "description": "add car info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "add car info",
                "parameters": [
                    {
                        "description": "regNum collection",
                        "name": "regNums",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegNumsInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Cars"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrRes"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.ErrRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrRes"
                        }
                    }
                }
            }
        },
        "/cars/{id}": {
            "delete": {
                "description": "delete car data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "car info ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CarInfo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrRes"
                        }
                    }
                }
            },
            "patch": {
                "description": "update car data by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "car info ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "car info struct",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CarInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CarInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrRes"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrRes"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.ErrRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CarInfo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/model.Person"
                },
                "regNum": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "model.CarInput": {
            "type": "object",
            "properties": {
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "owner": {
                    "type": "object",
                    "properties": {
                        "name": {
                            "type": "string"
                        },
                        "patronymic": {
                            "type": "string"
                        },
                        "surname": {
                            "type": "string"
                        }
                    }
                },
                "regNum": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "model.Cars": {
            "type": "object",
            "properties": {
                "cars": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CarInfo"
                    }
                }
            }
        },
        "model.ErrRes": {
            "type": "object",
            "properties": {
                "error": {}
            }
        },
        "model.Person": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "model.RegNumsInput": {
            "type": "object",
            "properties": {
                "regNums": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Car Catalog API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
