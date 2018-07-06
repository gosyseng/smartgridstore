package main
const SwaggerJSON = `{
  "swagger": "2.0",
  "info": {
    "title": "BTrDB v4 API",
    "version": "v4.11.1"
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
    "/v4/alignedwindows": {
      "post": {
        "operationId": "AlignedWindows",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceAlignedWindowsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceAlignedWindowsParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/changes": {
      "post": {
        "operationId": "Changes",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceChangesResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceChangesParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/create": {
      "post": {
        "operationId": "Create",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceCreateResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceCreateParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/delete": {
      "post": {
        "operationId": "Delete",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceDeleteResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceDeleteParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/faultinject": {
      "post": {
        "operationId": "FaultInject",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceFaultInjectResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceFaultInjectParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/flush": {
      "post": {
        "operationId": "Flush",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceFlushResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceFlushParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/generatecsv": {
      "post": {
        "operationId": "GenerateCSV",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceGenerateCSVResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceGenerateCSVParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/getmetadatausage": {
      "post": {
        "operationId": "GetMetadataUsage",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceMetadataUsageResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceMetadataUsageParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/info": {
      "post": {
        "operationId": "Info",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceInfoResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceInfoParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/insert": {
      "post": {
        "operationId": "Insert",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceInsertResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceInsertParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/listcollections": {
      "post": {
        "operationId": "ListCollections",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceListCollectionsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceListCollectionsParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/lookupstreams": {
      "post": {
        "operationId": "LookupStreams",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceLookupStreamsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceLookupStreamsParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/nearest": {
      "post": {
        "operationId": "Nearest",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceNearestResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceNearestParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/obliterate": {
      "post": {
        "operationId": "Obliterate",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceObliterateResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceObliterateParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/rawvalues": {
      "post": {
        "operationId": "RawValues",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceRawValuesResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceRawValuesParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/setstreamannotations": {
      "post": {
        "operationId": "SetStreamAnnotations",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceSetStreamAnnotationsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceSetStreamAnnotationsParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/streaminfo": {
      "post": {
        "operationId": "StreamInfo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceStreamInfoResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceStreamInfoParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/windows": {
      "post": {
        "operationId": "Windows",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceWindowsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceWindowsParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    }
  },
  "definitions": {
    "GenerateCSVParamsQueryType": {
      "type": "string",
      "enum": [
        "ALIGNED_WINDOWS_QUERY",
        "WINDOWS_QUERY",
        "RAW_QUERY"
      ],
      "default": "ALIGNED_WINDOWS_QUERY"
    },
    "grpcinterfaceAlignedWindowsParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "start": {
          "type": "string",
          "format": "int64"
        },
        "end": {
          "type": "string",
          "format": "int64"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "pointWidth": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "grpcinterfaceAlignedWindowsResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        },
        "values": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceStatPoint"
          }
        }
      }
    },
    "grpcinterfaceChangedRange": {
      "type": "object",
      "properties": {
        "start": {
          "type": "string",
          "format": "int64"
        },
        "end": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "grpcinterfaceChangesParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "fromMajor": {
          "type": "string",
          "format": "uint64"
        },
        "toMajor": {
          "type": "string",
          "format": "uint64"
        },
        "resolution": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "grpcinterfaceChangesResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        },
        "ranges": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceChangedRange"
          }
        }
      }
    },
    "grpcinterfaceCreateParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "collection": {
          "type": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyValue"
          }
        },
        "annotations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyValue"
          }
        }
      }
    },
    "grpcinterfaceCreateResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        }
      }
    },
    "grpcinterfaceDeleteParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "start": {
          "type": "string",
          "format": "int64"
        },
        "end": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "grpcinterfaceDeleteResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceFaultInjectParams": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "format": "uint64"
        },
        "params": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpcinterfaceFaultInjectResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "rv": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpcinterfaceFlushParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpcinterfaceFlushResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceGenerateCSVParams": {
      "type": "object",
      "properties": {
        "queryType": {
          "$ref": "#/definitions/GenerateCSVParamsQueryType"
        },
        "startTime": {
          "type": "string",
          "format": "int64"
        },
        "endTime": {
          "type": "string",
          "format": "int64"
        },
        "windowSize": {
          "type": "string",
          "format": "uint64"
        },
        "depth": {
          "type": "integer",
          "format": "int64"
        },
        "includeVersions": {
          "type": "boolean",
          "format": "boolean"
        },
        "streams": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceStreamCSVConfig"
          }
        }
      }
    },
    "grpcinterfaceGenerateCSVResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "isHeader": {
          "type": "boolean",
          "format": "boolean"
        },
        "row": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "grpcinterfaceInfoParams": {
      "type": "object"
    },
    "grpcinterfaceInfoResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "mash": {
          "$ref": "#/definitions/grpcinterfaceMash"
        },
        "majorVersion": {
          "type": "integer",
          "format": "int64"
        },
        "minorVersion": {
          "type": "integer",
          "format": "int64"
        },
        "build": {
          "type": "string"
        },
        "proxy": {
          "$ref": "#/definitions/grpcinterfaceProxyInfo"
        }
      }
    },
    "grpcinterfaceInsertParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "sync": {
          "type": "boolean",
          "format": "boolean"
        },
        "values": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceRawPoint"
          }
        }
      }
    },
    "grpcinterfaceInsertResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceKeyCount": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "count": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceKeyOptValue": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "val": {
          "$ref": "#/definitions/grpcinterfaceOptValue"
        }
      }
    },
    "grpcinterfaceKeyValue": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpcinterfaceListCollectionsParams": {
      "type": "object",
      "properties": {
        "prefix": {
          "type": "string"
        },
        "startWith": {
          "type": "string"
        },
        "limit": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceListCollectionsResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "collections": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "grpcinterfaceLookupStreamsParams": {
      "type": "object",
      "properties": {
        "collection": {
          "type": "string"
        },
        "isCollectionPrefix": {
          "type": "boolean",
          "format": "boolean"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyOptValue"
          }
        },
        "annotations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyOptValue"
          }
        }
      }
    },
    "grpcinterfaceLookupStreamsResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "results": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceStreamDescriptor"
          }
        }
      }
    },
    "grpcinterfaceMash": {
      "type": "object",
      "properties": {
        "revision": {
          "type": "string",
          "format": "int64"
        },
        "leader": {
          "type": "string"
        },
        "leaderRevision": {
          "type": "string",
          "format": "int64"
        },
        "totalWeight": {
          "type": "string",
          "format": "int64"
        },
        "healthy": {
          "type": "boolean",
          "format": "boolean"
        },
        "unmapped": {
          "type": "number",
          "format": "double"
        },
        "members": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceMember"
          }
        }
      }
    },
    "grpcinterfaceMember": {
      "type": "object",
      "properties": {
        "hash": {
          "type": "integer",
          "format": "int64"
        },
        "nodename": {
          "type": "string"
        },
        "up": {
          "type": "boolean",
          "format": "boolean"
        },
        "in": {
          "type": "boolean",
          "format": "boolean"
        },
        "enabled": {
          "type": "boolean",
          "format": "boolean"
        },
        "start": {
          "type": "string",
          "format": "int64"
        },
        "end": {
          "type": "string",
          "format": "int64"
        },
        "weight": {
          "type": "string",
          "format": "int64"
        },
        "readPreference": {
          "type": "number",
          "format": "double"
        },
        "httpEndpoints": {
          "type": "string"
        },
        "grpcEndpoints": {
          "type": "string"
        }
      }
    },
    "grpcinterfaceMetadataUsageParams": {
      "type": "object",
      "properties": {
        "prefix": {
          "type": "string"
        }
      }
    },
    "grpcinterfaceMetadataUsageResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyCount"
          }
        },
        "annotations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyCount"
          }
        }
      }
    },
    "grpcinterfaceNearestParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "time": {
          "type": "string",
          "format": "int64"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "backward": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "grpcinterfaceNearestResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        },
        "value": {
          "$ref": "#/definitions/grpcinterfaceRawPoint"
        }
      }
    },
    "grpcinterfaceObliterateParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpcinterfaceObliterateResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        }
      }
    },
    "grpcinterfaceOptValue": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpcinterfaceProxyInfo": {
      "type": "object",
      "properties": {
        "proxyEndpoints": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "grpcinterfaceRawPoint": {
      "type": "object",
      "properties": {
        "time": {
          "type": "string",
          "format": "int64"
        },
        "value": {
          "type": "number",
          "format": "double"
        }
      }
    },
    "grpcinterfaceRawValuesParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "start": {
          "type": "string",
          "format": "int64"
        },
        "end": {
          "type": "string",
          "format": "int64"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceRawValuesResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        },
        "values": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceRawPoint"
          }
        }
      }
    },
    "grpcinterfaceSetStreamAnnotationsParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "expectedAnnotationVersion": {
          "type": "string",
          "format": "uint64"
        },
        "annotations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyOptValue"
          }
        }
      }
    },
    "grpcinterfaceSetStreamAnnotationsResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        }
      }
    },
    "grpcinterfaceStatPoint": {
      "type": "object",
      "properties": {
        "time": {
          "type": "string",
          "format": "int64"
        },
        "min": {
          "type": "number",
          "format": "double"
        },
        "mean": {
          "type": "number",
          "format": "double"
        },
        "max": {
          "type": "number",
          "format": "double"
        },
        "count": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceStatus": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "msg": {
          "type": "string"
        },
        "mash": {
          "$ref": "#/definitions/grpcinterfaceMash"
        }
      }
    },
    "grpcinterfaceStreamCSVConfig": {
      "type": "object",
      "properties": {
        "version": {
          "type": "string",
          "format": "uint64"
        },
        "label": {
          "type": "string"
        },
        "uuid": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpcinterfaceStreamDescriptor": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "collection": {
          "type": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyValue"
          }
        },
        "annotations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyValue"
          }
        },
        "annotationVersion": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceStreamInfoParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "omitVersion": {
          "type": "boolean",
          "format": "boolean"
        },
        "omitDescriptor": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "grpcinterfaceStreamInfoResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        },
        "descriptor": {
          "$ref": "#/definitions/grpcinterfaceStreamDescriptor"
        }
      }
    },
    "grpcinterfaceWindowsParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "start": {
          "type": "string",
          "format": "int64"
        },
        "end": {
          "type": "string",
          "format": "int64"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "width": {
          "type": "string",
          "format": "uint64"
        },
        "depth": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "grpcinterfaceWindowsResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        },
        "values": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceStatPoint"
          }
        }
      }
    }
  }
}
`;