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
        "/address": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all addresses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "address"
                ],
                "summary": "Get all addresses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.AddressDTO"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create an address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "address"
                ],
                "summary": "Create an address",
                "parameters": [
                    {
                        "description": "Address to be created",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.AddressDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.AddressDTO"
                        }
                    }
                }
            }
        },
        "/address/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get an address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "address"
                ],
                "summary": "Get an address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Address ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.AddressDTO"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete an address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "address"
                ],
                "summary": "Delete an address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Address ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Address deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/address/{id}/main": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Mark an address as default",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "address"
                ],
                "summary": "Mark an address as default",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Address ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Check if the server is healthy",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/product": {
            "post": {
                "description": "Create a product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Create a product",
                "parameters": [
                    {
                        "description": "Product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.ProductDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.ProductDTO"
                        }
                    }
                }
            }
        },
        "/product/{sku}": {
            "get": {
                "description": "Get a product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Get a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product SKU",
                        "name": "sku",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ProductDTO"
                        }
                    }
                }
            }
        },
        "/product/{terms}": {
            "get": {
                "description": "Search a product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Search a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search terms",
                        "name": "terms",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.ProductDTO"
                            }
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.UserRegisterDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.UserLoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserLoginResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/refresh-token": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Refresh token",
                "tags": [
                    "user"
                ],
                "summary": "Refresh token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserLoginResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.AddressDTO": {
            "type": "object",
            "required": [
                "city",
                "country",
                "neighborhood",
                "number",
                "state",
                "street",
                "zip_code"
            ],
            "properties": {
                "city": {
                    "type": "string"
                },
                "complement": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "main_address": {
                    "type": "boolean"
                },
                "neighborhood": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "zip_code": {
                    "type": "string"
                }
            }
        },
        "dtos.ProductDTO": {
            "type": "object",
            "required": [
                "description",
                "name",
                "name_url",
                "principal_image",
                "quantity",
                "sku"
            ],
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "brands": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "categorys": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "name_url": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "principal_image": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "sku": {
                    "type": "string"
                }
            }
        },
        "dtos.UserDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dtos.UserLoginDTO": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dtos.UserLoginResponseDTO": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dtos.UserRegisterDTO": {
            "type": "object",
            "required": [
                "confirm_password",
                "document",
                "email",
                "last_name",
                "name",
                "password"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "document": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
