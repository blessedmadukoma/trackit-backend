// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://trakkit.vercel.app",
        "contact": {
            "name": "Madukoma Blessed",
            "url": "https://mblessed.vercel.app",
            "email": "blessedmadukoma@gmail.com"
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
        "/expense": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Responds with a list of expense records as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expense"
                ],
                "summary": "Get Expenses Transactions",
                "parameters": [
                    {
                        "description": "Expense JSON",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.listExpensesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.expenseResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "responses": {}
            }
        }
    },
    "definitions": {
        "api.expenseResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "created_at": {
                    "description": "budgetid int64 ` + "`" + `json:\"budgetid\" binding:\"required\"` + "`" + `",
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "api.listExpensesRequest": {
            "type": "object",
            "required": [
                "page_id",
                "page_size"
            ],
            "properties": {
                "page_id": {
                    "type": "integer",
                    "minimum": 1
                },
                "page_size": {
                    "type": "integer",
                    "maximum": 10,
                    "minimum": 5
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "https://trackit-blessedmadukoma.koyeb.app",
	BasePath:         "/api/",
	Schemes:          []string{"https"},
	Title:            "Trakkit Backend",
	Description:      "Backend for TrakkIT, a financial management tracking tool",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
