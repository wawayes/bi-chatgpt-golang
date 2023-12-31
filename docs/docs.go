// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/chart/all_list": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ChartApi"
                ],
                "summary": "ListAllChart",
                "parameters": [
                    {
                        "description": "查询请求参数",
                        "name": "ChartQueryRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.ChartQueryRequest"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/response.BiResp"
                        }
                    },
                    "40002": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40003": {
                        "description": "系统错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    }
                }
            }
        },
        "/chart/gen": {
            "post": {
                "description": "通过上传文件和发送请求生成图表",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ChartApi"
                ],
                "summary": "生成图表",
                "parameters": [
                    {
                        "type": "file",
                        "description": "要上传的文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "生成图表的目标",
                        "name": "goal",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图表类型",
                        "name": "chartType",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/response.BiResp"
                        }
                    },
                    "40002": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40003": {
                        "description": "系统错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    }
                }
            }
        },
        "/chart/list": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ChartApi"
                ],
                "summary": "Chart List",
                "parameters": [
                    {
                        "description": "查询请求参数",
                        "name": "ChartQueryRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.ChartQueryRequest"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/response.BiResp"
                        }
                    },
                    "40002": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40003": {
                        "description": "系统错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    }
                }
            }
        },
        "/table/list": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TableApi"
                ],
                "summary": "ListUserTable",
                "parameters": [
                    {
                        "description": "登录请求参数",
                        "name": "pageRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.Page"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "成功",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserChart"
                            }
                        }
                    },
                    "40002": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40003": {
                        "description": "系统错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    }
                }
            }
        },
        "/user/current": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "Current",
                "responses": {
                    "0": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/serializers.CurrentUser"
                        }
                    },
                    "40005": {
                        "description": "获取当前用户信息失败",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "List",
                "responses": {
                    "0": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40002": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40003": {
                        "description": "系统错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "登录请求参数",
                        "name": "loginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40002": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40003": {
                        "description": "系统错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "Logout",
                "responses": {
                    "0": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40002": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40003": {
                        "description": "系统错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    }
                }
            }
        },
        "/user/refresh_token": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "RefreshToken",
                "responses": {
                    "0": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40005": {
                        "description": "认证失败",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "User Register",
                "parameters": [
                    {
                        "description": "注册请求参数",
                        "name": "registerRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40002": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40003": {
                        "description": "系统错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.UserChart": {
            "type": "object",
            "properties": {
                "freeCount": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "token": {
                    "type": "integer"
                },
                "userAccount": {
                    "type": "string"
                },
                "userAvatar": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "r.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "requests.ChartQueryRequest": {
            "type": "object",
            "properties": {
                "chartType": {
                    "type": "string"
                },
                "goal": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                },
                "pageNum": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "requests.LoginRequest": {
            "type": "object",
            "properties": {
                "userAccount": {
                    "type": "string"
                },
                "userPassword": {
                    "type": "string"
                }
            }
        },
        "requests.Page": {
            "type": "object",
            "properties": {
                "pageNum": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        },
        "requests.RegisterRequest": {
            "type": "object",
            "properties": {
                "checkPassword": {
                    "type": "string"
                },
                "userAccount": {
                    "type": "string"
                },
                "userPassword": {
                    "type": "string"
                }
            }
        },
        "response.BiResp": {
            "type": "object",
            "properties": {
                "genChart": {
                    "type": "string"
                },
                "genResult": {
                    "type": "string"
                }
            }
        },
        "serializers.CurrentUser": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "userAccount": {
                    "type": "string"
                },
                "userAvatar": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                },
                "userRole": {
                    "type": "string"
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "BI Pro API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
