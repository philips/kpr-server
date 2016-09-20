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
    "/repos/index": {
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
    "/repos/{name}/blobs/sha256/{digest}": {
      "get": {
        "summary": "TODO: custom marshaller",
        "operationId": "GetBlob",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/kprpbBlobResponse"
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
          },
          {
            "name": "digest",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          }
        ],
        "tags": [
          "Blob"
        ]
      },
      "put": {
        "operationId": "PutBlob",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/kprpbDescriptor"
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
          },
          {
            "name": "digest",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/kprpbBlobRequest"
            }
          }
        ],
        "tags": [
          "Blob"
        ]
      }
    },
    "/repos/{name}/tags": {
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
    },
    "/repos/{name}/tags/{tag}": {
      "get": {
        "operationId": "GetTag",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/kprpbPackageManifest"
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
          },
          {
            "name": "tag",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          }
        ],
        "tags": [
          "Repo"
        ]
      },
      "put": {
        "operationId": "DeleteTag",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/kprpbDescriptor"
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
          },
          {
            "name": "tag",
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
    "kprpbBlobRequest": {
      "type": "object",
      "properties": {
        "blob": {
          "type": "string",
          "format": "byte"
        },
        "digest": {
          "type": "string",
          "format": "string"
        },
        "name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "kprpbBlobResponse": {
      "type": "object",
      "properties": {
        "blob": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "kprpbDescriptor": {
      "type": "object",
      "properties": {
        "digest": {
          "type": "string",
          "format": "string"
        },
        "mediaType": {
          "type": "string",
          "format": "string"
        },
        "size": {
          "type": "string",
          "format": "int64"
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
    },
    "kprpbPackageManifest": {
      "type": "object",
      "properties": {
        "dependencies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/kprpbDescriptor"
          }
        },
        "mediaType": {
          "type": "string",
          "format": "string"
        },
        "package": {
          "$ref": "#/definitions/kprpbDescriptor"
        },
        "packages": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/kprpbDescriptor"
          }
        },
        "schemaVersion": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "kprpbTagRequest": {
      "type": "object",
      "properties": {
        "desc": {
          "$ref": "#/definitions/kprpbDescriptor"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "tag": {
          "type": "string",
          "format": "string"
        }
      }
    }
  }
}
`
)
