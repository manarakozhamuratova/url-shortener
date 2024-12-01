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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/shortener": {
            "get": {
                "description": "Возвращает список всех созданных коротких ссылок",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shortUrl"
                ],
                "summary": "Получение списка коротких ссылок",
                "responses": {
                    "200": {
                        "description": "Список коротких ссылок",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Создание короткой ссылки",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shortUrl"
                ],
                "summary": "Создание короткой ссылки",
                "parameters": [
                    {
                        "description": "Входящие данные",
                        "name": "rq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_transport_handlers.CreateShortUrlRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_handlers.CreateShortUrlResponse"
                        }
                    },
                    "400": {
                        "description": "Некорректный URL",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/stats/{link}": {
            "get": {
                "description": "Возвращает статистику по количеству переходов по короткой ссылке",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shortUrl"
                ],
                "summary": "Получение статистики по ссылке",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Короткая ссылка",
                        "name": "link",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Статистика по ссылке",
                        "schema": {
                            "$ref": "#/definitions/urlshortener_internal_models.Stat"
                        }
                    },
                    "404": {
                        "description": "Статистика не найдена",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/{link}": {
            "get": {
                "description": "Перенаправляет на оригинальную ссылку, связанную с короткой ссылкой",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shortUrl"
                ],
                "summary": "Получение оригинальной ссылки",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Короткая ссылка",
                        "name": "link",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Перенаправление на оригинальную ссылку",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Короткая ссылка не найдена",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет короткую ссылку из базы данных",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shortUrl"
                ],
                "summary": "Удаление короткой ссылки",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Короткая ссылка",
                        "name": "link",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Короткая ссылка успешно удалена"
                    },
                    "404": {
                        "description": "Короткая ссылка не найдена",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "internal_transport_handlers.CreateShortUrlRequest": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "internal_transport_handlers.CreateShortUrlResponse": {
            "type": "object",
            "properties": {
                "short_url": {
                    "type": "string"
                }
            }
        },
        "urlshortener_internal_models.Stat": {
            "type": "object",
            "properties": {
                "url_id": {
                    "type": "integer"
                },
                "visited_at": {
                    "type": "string"
                },
                "visited_count": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9090",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Super API",
	Description:      "This is my first swagger documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
