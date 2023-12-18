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
        "/api/image_names": {
            "get": {
                "description": "查看图片简单信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "查看图片简单信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-images_api_ImageResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/images": {
            "get": {
                "description": "图片列表查询",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "图片列表查询",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "列表页数",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "查询关键词",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每一页的条数",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "显示顺序规则",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-models_BannerModel"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "修改图片名称",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "修改图片名称",
                "parameters": [
                    {
                        "description": "修改图片名称",
                        "name": "images",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/images_api.ImageUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "上传图片",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "上传图片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "上传的一系列图片(实际参数:multipart.Form)",
                        "name": "images",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-image_ser_FileUploadResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "删除图片",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "删除图片",
                "parameters": [
                    {
                        "description": "删除图片的列表",
                        "name": "images",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RemoveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/menu_names": {
            "get": {
                "description": "菜单名称列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "菜单名称列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-models_MenuModel"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/menus": {
            "get": {
                "description": "菜单列表查询",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "菜单列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-models_MenuModel"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "添加菜单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单管理"
                ],
                "summary": "添加菜单",
                "parameters": [
                    {
                        "description": "添加菜单",
                        "name": "file",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/menu_api.MenuRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/settings/:name": {
            "get": {
                "description": "显示某一项配置信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "配置管理"
                ],
                "summary": "显示配置信息",
                "parameters": [
                    {
                        "description": "表示单个参数",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/settings_api.SettingsUri"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "修改某一项配置信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "配置管理"
                ],
                "summary": "修改配置信息",
                "parameters": [
                    {
                        "description": "表示单个参数",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/settings_api.SettingsUri"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ctype.ImageType": {
            "type": "integer",
            "enum": [
                1,
                2
            ],
            "x-enum-comments": {
                "Local": "本地",
                "Telegraph": "telegraph"
            },
            "x-enum-varnames": [
                "Local",
                "Telegraph"
            ]
        },
        "image_ser.FileUploadResponse": {
            "type": "object",
            "properties": {
                "file_name": {
                    "description": "文件名",
                    "type": "string"
                },
                "is_success": {
                    "description": "是否上传成功",
                    "type": "boolean"
                },
                "msg": {
                    "description": "消息",
                    "type": "string"
                }
            }
        },
        "images_api.ImageResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "图片id",
                    "type": "integer"
                },
                "name": {
                    "description": "图片名称",
                    "type": "string"
                },
                "path": {
                    "description": "图片路径",
                    "type": "string"
                }
            }
        },
        "images_api.ImageUpdate": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "menu_api.ImageSort": {
            "type": "object",
            "properties": {
                "image_id": {
                    "type": "integer"
                },
                "sort": {
                    "type": "integer"
                }
            }
        },
        "menu_api.MenuRequest": {
            "type": "object",
            "required": [
                "path",
                "sort",
                "title"
            ],
            "properties": {
                "abstract": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "abstract_time": {
                    "description": "切换的时间，单位秒",
                    "type": "integer"
                },
                "banner_time": {
                    "description": "切换的时间，单位秒",
                    "type": "integer"
                },
                "image_sort_list": {
                    "description": "具体图片的顺序",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/menu_api.ImageSort"
                    }
                },
                "path": {
                    "type": "string"
                },
                "slogan": {
                    "type": "string"
                },
                "sort": {
                    "description": "菜单的序号",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.BannerModel": {
            "type": "object",
            "properties": {
                "Hash": {
                    "description": "图片的Hash值，用于判断重复图片",
                    "type": "string"
                },
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "image_type": {
                    "description": "图片的位置(本地还是telegraph)",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ctype.ImageType"
                        }
                    ]
                },
                "name": {
                    "description": "图片名称",
                    "type": "string"
                },
                "path": {
                    "description": "图片路径",
                    "type": "string"
                }
            }
        },
        "models.MenuModel": {
            "type": "object",
            "properties": {
                "abstract": {
                    "description": "简介",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "abstract_time": {
                    "description": "简介的切换时间",
                    "type": "integer"
                },
                "banner_time": {
                    "description": "菜单的切换时间 为0表示不切换",
                    "type": "integer"
                },
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "menu_images": {
                    "description": "菜单的图片列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.BannerModel"
                    }
                },
                "path": {
                    "type": "string"
                },
                "slogan": {
                    "description": "slogan",
                    "type": "string"
                },
                "sort": {
                    "description": "菜单的顺序",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.RemoveRequest": {
            "type": "object",
            "properties": {
                "id_list": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "res.CodeType": {
            "type": "integer",
            "enum": [
                0,
                1,
                1001,
                1002,
                1003
            ],
            "x-enum-comments": {
                "ArgumentError": "参数错误",
                "SettingsError": "系统错误",
                "UploadError": "上传错误"
            },
            "x-enum-varnames": [
                "SUCCESS",
                "Error",
                "SettingsError",
                "ArgumentError",
                "UploadError"
            ]
        },
        "res.ListResponse-image_ser_FileUploadResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "$ref": "#/definitions/image_ser.FileUploadResponse"
                }
            }
        },
        "res.ListResponse-images_api_ImageResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "$ref": "#/definitions/images_api.ImageResponse"
                }
            }
        },
        "res.ListResponse-models_BannerModel": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "$ref": "#/definitions/models.BannerModel"
                }
            }
        },
        "res.ListResponse-models_MenuModel": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "$ref": "#/definitions/models.MenuModel"
                }
            }
        },
        "res.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/res.CodeType"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "settings_api.SettingsUri": {
            "type": "object",
            "properties": {
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
	Host:             "127.0.0.1:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "gvb_service API文档",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
