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
        "/expert/ads/{adID}": {
            "get": {
                "description": "retrieve expert check request for expert or user",
                "tags": [
                    "expert"
                ],
                "summary": "retrieve expert check request for expert or user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ad ID",
                        "name": "adID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetExpertRequestResponse"
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete expert check request for expert or user",
                "tags": [
                    "expert"
                ],
                "summary": "delete expert check request for expert or user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ad ID",
                        "name": "adID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/expert/ads/{adID}/check-request": {
            "post": {
                "description": "Request to expert check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expert"
                ],
                "summary": "Request to expert check",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Ad ID",
                        "name": "adID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/expert/check-request/{expertRequestID}": {
            "put": {
                "description": "Update expert check request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expert"
                ],
                "summary": "Update expert check request",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "expert request ID",
                        "name": "expertRequestID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Expert check object",
                        "name": "expertCheckRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateExpertCheckRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ExpertRequestResponse"
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/expert/check-requests": {
            "get": {
                "description": "ListExpertRequest retrieves all expert requests for an expert",
                "tags": [
                    "expert"
                ],
                "summary": "ListExpertRequest retrieves all expert requests for an expert",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Ad ID",
                        "name": "ads_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "From date",
                        "name": "from_date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ExpertRequestResponse"
                            }
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/repair/ads/{adID}": {
            "get": {
                "description": "retrieve repair check request for repair or user",
                "tags": [
                    "repair"
                ],
                "summary": "retrieve repair check request for repair or user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ad ID",
                        "name": "adID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetRepairRequestResponse"
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/repair/ads/{adID}/check-request": {
            "post": {
                "description": "Request to repair check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "repair"
                ],
                "summary": "Request to repair check",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Ad ID",
                        "name": "adID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessResponse"
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/repair/check-request/{repairRequestID}": {
            "put": {
                "description": "Update repair request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "repair"
                ],
                "summary": "Update repair request",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "repair request ID",
                        "name": "repairRequestID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "repair object",
                        "name": "repairCheckRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateRepairRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RepairRequestResponse"
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/repair/check-requests": {
            "get": {
                "description": "ListRepairRequest retrieves all repair requests for an repair",
                "tags": [
                    "repair"
                ],
                "summary": "ListRepairRequest retrieves all repair requests for an repair",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Ad ID",
                        "name": "ads_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "From date",
                        "name": "from_date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.RepairRequestResponse"
                            }
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.ExpertRequestResponse": {
            "type": "object",
            "properties": {
                "adID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "expertID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.GetExpertRequestResponse": {
            "type": "object",
            "properties": {
                "adSubject": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "expertID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "userName": {
                    "type": "integer"
                }
            }
        },
        "models.GetRepairRequestResponse": {
            "type": "object",
            "properties": {
                "adSubject": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "userName": {
                    "type": "integer"
                }
            }
        },
        "models.RepairRequestResponse": {
            "type": "object",
            "properties": {
                "adID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.SuccessResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.UpdateExpertCheckRequest": {
            "type": "object",
            "properties": {
                "report": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/utils.Status"
                }
            }
        },
        "models.UpdateRepairRequest": {
            "type": "object",
            "properties": {
                "status": {
                    "$ref": "#/definitions/utils.Status"
                }
            }
        },
        "utils.Status": {
            "type": "string",
            "enum": [
                "Wait for payment status",
                "Pending for expert",
                "Pending for matin",
                "In progress",
                "Done"
            ],
            "x-enum-varnames": [
                "WAIT_FOR_PAYMENT_STATUS",
                "EXPERT_PENDING_STATUS",
                "MATIN_PENDING_STATUS",
                "IN_PROGRESS_STATUS",
                "DONE_STATUS"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Airplane-Divar",
	Description:      "Quera Airplane-Divar server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
