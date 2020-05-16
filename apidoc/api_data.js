define({ "api": [
  {
    "type": "get",
    "url": "/article/detail",
    "title": "文章详情",
    "group": "Article",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "integer",
            "optional": false,
            "field": "id",
            "description": "<p>文章id</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"data\": {\n        \"id\": 1,\n        \"created_at\": \"2020-03-24T09:20:05+08:00\",\n        \"updated_at\": \"2020-03-24T09:20:05+08:00\",\n        \"deleted_at\": null,\n        \"uid\": 2,\n        \"title\": \"标题2\",\n        \"content\": \"内容2\",\n        \"favorite\": 0\n    },\n    \"message\": \"\",\n    \"statusCode\": 200\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "route/api/article.go",
    "groupTitle": "Article",
    "name": "GetArticleDetail"
  },
  {
    "type": "post",
    "url": "/article/add",
    "title": "添加文章",
    "group": "Article",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "title",
            "description": "<p>文章标题</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "content",
            "description": "<p>文章内容</p>"
          },
          {
            "group": "Parameter",
            "type": "integer",
            "optional": false,
            "field": "uid",
            "description": "<p>用户</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"data\": null,\n    \"message\": \"\",\n    \"statusCode\": 200\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "route/api/article.go",
    "groupTitle": "Article",
    "name": "PostArticleAdd"
  },
  {
    "type": "get",
    "url": "/user/info",
    "title": "用户详情",
    "group": "User",
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"data\": {\n        \"age\": 19,\n        \"last_login_at\": null,\n        \"user_name\": \"张三2\"\n    },\n    \"message\": \"\",\n    \"statusCode\": 200\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "route/api/user.go",
    "groupTitle": "User",
    "name": "GetUserInfo"
  },
  {
    "type": "get",
    "url": "/user/list",
    "title": "用户列表",
    "group": "User",
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"data\": [\n        {\n            \"id\": 1,\n            \"created_at\": \"2020-05-16T17:44:51+08:00\",\n            \"updated_at\": \"2020-05-16T17:44:51+08:00\",\n            \"deleted_at\": null,\n            \"user_name\": \"张三2\",\n            \"age\": 19,\n            \"password\": \"$2a$04$O13ojGm.yz0NPrRSuRpkhuEVXvR13S501SchwQUvHMAjbOEVC1U9e\",\n            \"last_login_at\": null,\n            \"articles\": null\n        }\n    ],\n    \"message\": \"\",\n    \"statusCode\": 200\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "route/api/user.go",
    "groupTitle": "User",
    "name": "GetUserList"
  },
  {
    "type": "get",
    "url": "/user/list_with_article",
    "title": "用户列表以及发表的文章",
    "group": "User",
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"data\": [\n        {\n            \"age\": 19,\n            \"articles\": [\n                {\n                    \"content\": \"内容2\",\n                    \"id\": 1,\n                    \"title\": \"标题2\"\n                }\n            ],\n            \"id\": 1,\n            \"last_login_at\": null,\n            \"user_name\": \"张三2\"\n        }\n    ],\n    \"message\": \"\",\n    \"statusCode\": 200\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "route/api/user.go",
    "groupTitle": "User",
    "name": "GetUserList_with_article"
  },
  {
    "type": "get",
    "url": "/user/token",
    "title": "获取token",
    "group": "User",
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"data\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIn0.NgowRjbanOQa3-B5-q6JjxCVjT9dmn1AeeercH1zGSU\",\n    \"message\": \"\",\n    \"statusCode\": 200\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "route/api/user.go",
    "groupTitle": "User",
    "name": "GetUserToken"
  },
  {
    "type": "post",
    "url": "/user/add",
    "title": "添加用户",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "age",
            "description": "<p>年龄</p>"
          },
          {
            "group": "Parameter",
            "type": "integer",
            "optional": false,
            "field": "password",
            "description": "<p>登录密码</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"data\": null,\n    \"message\": \"\",\n    \"statusCode\": 200\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "route/api/user.go",
    "groupTitle": "User",
    "name": "PostUserAdd"
  }
] });
