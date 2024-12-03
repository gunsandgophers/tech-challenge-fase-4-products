// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/product/": {
            "post": {
                "description": "register the product information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Create a product",
                "parameters": [
                    {
                        "description": "Insert Product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ProductDTO"
                        }
                    },
                    "400": {
                        "description": "when bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "when create product process error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/product/category/{category}/": {
            "get": {
                "description": "List a set of products information over followed categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "List products",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product category",
                        "name": "category",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Page defaults to 1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Size defaults to 50",
                        "name": "size",
                        "in": "query"
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
                    },
                    "400": {
                        "description": "when bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "when list products process error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/product/{id}/": {
            "get": {
                "description": "Get product information by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product identifier",
                        "name": "id",
                        "in": "path",
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
                    },
                    "400": {
                        "description": "when bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "when list products process error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "change the product information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Update a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product identifier",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Product object",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ProductDTO"
                        }
                    },
                    "400": {
                        "description": "when bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "when update product process error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "remove the product information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Delete a product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product identifier",
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
                    },
                    "400": {
                        "description": "when when bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "when delete product process error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.ProductRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "dtos.ProductDTO": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
