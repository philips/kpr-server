package echopb 

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
    }
  }
}
`
)
