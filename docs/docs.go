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
        "/task": {
            "get": {
                "description": "Get total time spent by a user on tasks within a period",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Get total time spent by a user on tasks within a period",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"2023-30-12 00:00:00\"",
                        "description": "Begin period",
                        "name": "begin_period",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"2023-31-12 23:59:59\"",
                        "description": "End period",
                        "name": "end_period",
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
                                "type": "object",
                                "properties": {
                                    "task_id": {
                                        "type": "integer"
                                    },
                                    "time_sum": {
                                        "type": "number"
                                    }
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/start": {
            "post": {
                "description": "Start a task for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Start a task for a user",
                "parameters": [
                    {
                        "description": "User and Task IDs",
                        "name": "userReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/actions.userRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TaskBind"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User or Task not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/stop": {
            "post": {
                "description": "Stop a task for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Stop a task for a user",
                "parameters": [
                    {
                        "description": "User and Task IDs",
                        "name": "userReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/actions.userRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TaskBind"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User or Task not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Get a list of all users with optional filters for pagination and search",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get all users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "User name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "User surname",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "User patronymic",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "User passport serie",
                        "name": "passportSerie",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "User passport number",
                        "name": "passportNumber",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "User address",
                        "name": "address",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Number of users per page",
                        "name": "perPage",
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
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Users not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing user with the provided information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User request body",
                        "name": "userReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/actions.updateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user with the provided information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User request body",
                        "name": "userReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/actions.createUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "delete": {
                "description": "Delete an existing user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "actions.createUserRequest": {
            "type": "object",
            "properties": {
                "passportNumber": {
                    "type": "string"
                }
            }
        },
        "actions.updateUserRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "passport_number": {
                    "type": "integer"
                },
                "passport_serie": {
                    "type": "integer"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "actions.userRequest": {
            "type": "object",
            "properties": {
                "task_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.TaskBind": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "finish_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "start_at": {
                    "type": "string"
                },
                "task_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "passport_number": {
                    "type": "integer"
                },
                "passport_serie": {
                    "type": "integer"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "updated_at": {
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
