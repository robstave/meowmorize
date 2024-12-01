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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cards": {
            "post": {
                "description": "Create a new card and associate it with a deck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cards"
                ],
                "summary": "Create a new card",
                "parameters": [
                    {
                        "description": "Create Card",
                        "name": "card",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateCardRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.Card"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/cards/stats": {
            "post": {
                "description": "Update the statistics of a card based on the specified action",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cards"
                ],
                "summary": "Update card statistics",
                "parameters": [
                    {
                        "description": "Card Stats Update",
                        "name": "stats",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CardStatsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Card"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/cards/{id}": {
            "get": {
                "description": "Retrieve a single card by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cards"
                ],
                "summary": "Get a card by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Card"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update the details of an existing card by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cards"
                ],
                "summary": "Update an existing card",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Card",
                        "name": "card",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateCardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Card"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a card by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cards"
                ],
                "summary": "Delete a card",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Card ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/decks": {
            "get": {
                "description": "Retrieve a list of all decks",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Decks"
                ],
                "summary": "Get all decks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Deck"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new deck with provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Decks"
                ],
                "summary": "Create a new deck",
                "parameters": [
                    {
                        "description": "Deck to create",
                        "name": "deck",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Deck"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.Deck"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/decks/default": {
            "post": {
                "description": "Create a new deck with provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Decks"
                ],
                "summary": "Create a default deck",
                "parameters": [
                    {
                        "description": "Deck to create",
                        "name": "deck",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Deck"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.Deck"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/decks/export/{id}": {
            "get": {
                "description": "Export a deck as a JSON file",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "decks"
                ],
                "summary": "Export a deck",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Deck ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Deck"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/decks/import": {
            "post": {
                "description": "Import a new deck by uploading a JSON file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Decks"
                ],
                "summary": "Import a deck from a JSON file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Deck JSON File",
                        "name": "deck_file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.Deck"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/decks/{id}": {
            "get": {
                "description": "Retrieve a single deck by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Decks"
                ],
                "summary": "Get a deck by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Deck ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Deck"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing deck by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Decks"
                ],
                "summary": "Update a deck",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Deck ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated Deck",
                        "name": "deck",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Deck"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Deck"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a deck by its ID",
                "tags": [
                    "Decks"
                ],
                "summary": "Delete a deck",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Deck ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.CardContentReq": {
            "type": "object",
            "required": [
                "text"
            ],
            "properties": {
                "text": {
                    "type": "string"
                }
            }
        },
        "controller.CardStatsRequest": {
            "type": "object",
            "required": [
                "action",
                "card_id"
            ],
            "properties": {
                "action": {
                    "type": "string",
                    "enum": [
                        "IncrementFail",
                        "IncrementPass",
                        "IncrementSkip",
                        "SetStars",
                        "Retire",
                        "Unretire",
                        "ResetStats"
                    ]
                },
                "card_id": {
                    "type": "string"
                },
                "value": {
                    "description": "Used only for SetStars",
                    "type": "integer"
                }
            }
        },
        "controller.CreateCardRequest": {
            "type": "object",
            "required": [
                "back",
                "deck_id",
                "front"
            ],
            "properties": {
                "back": {
                    "$ref": "#/definitions/controller.CardContentReq"
                },
                "deck_id": {
                    "type": "string"
                },
                "front": {
                    "$ref": "#/definitions/controller.CardContentReq"
                },
                "link": {
                    "type": "string"
                }
            }
        },
        "controller.UpdateCardRequest": {
            "type": "object",
            "properties": {
                "back": {
                    "$ref": "#/definitions/controller.CardContentReq"
                },
                "front": {
                    "$ref": "#/definitions/controller.CardContentReq"
                },
                "link": {
                    "type": "string"
                }
            }
        },
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "types.Card": {
            "type": "object",
            "properties": {
                "back": {
                    "$ref": "#/definitions/types.CardBack"
                },
                "created_at": {
                    "type": "string"
                },
                "deck_id": {
                    "type": "string"
                },
                "fail_count": {
                    "type": "integer"
                },
                "front": {
                    "$ref": "#/definitions/types.CardFront"
                },
                "id": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "pass_count": {
                    "type": "integer"
                },
                "retired": {
                    "type": "boolean"
                },
                "reviewed_at": {
                    "type": "string"
                },
                "skip_count": {
                    "type": "integer"
                },
                "star_rating": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "types.CardBack": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string"
                }
            }
        },
        "types.CardFront": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string"
                }
            }
        },
        "types.Deck": {
            "type": "object",
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Card"
                    }
                },
                "description": {
                    "description": "New Description field",
                    "type": "string"
                },
                "id": {
                    "description": "UUID string as the primary key",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "192.168.86.176:8789",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "MeowMorize Flashcard API",
	Description:      "API documentation for the MeowMorize Flashcard App.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
