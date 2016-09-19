package kprpb 

const (
swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "service.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/repos": {
      "get": {
        "operationId": "List",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/kprpbListResponse"
            }
          }
        },
        "tags": [
          "Repo"
        ]
      }
    },
    "/{name}/tags": {
      "get": {
        "operationId": "ListTags",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/kprpbListTagsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          }
        ],
        "tags": [
          "Repo"
        ]
      }
    }
  },
  "definitions": {
    "ListResponseRepos": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "ListTagsResponseTags": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "kprpbListRequest": {
      "type": "object"
    },
    "kprpbListResponse": {
      "type": "object",
      "properties": {
        "repos": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ListResponseRepos"
          }
        }
      }
    },
    "kprpbListTagsRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "kprpbListTagsResponse": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ListTagsResponseTags"
          }
        }
      }
    }
  }
}
`
)
