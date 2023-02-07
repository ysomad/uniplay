// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Uniplay private API",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/v1",
  "paths": {
    "/compendiums/weapon-classes": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "compendium"
        ],
        "summary": "Получение списка классов оружий",
        "operationId": "getWeaponClasses",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/WeaponClassList"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/compendiums/weapons": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "compendium"
        ],
        "summary": "Получение списка оружий",
        "operationId": "getWeapons",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/WeaponList"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/institutions": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "institution"
        ],
        "summary": "Получение списка учебных заведений",
        "operationId": "getInstitutions",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/WeaponList"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/matches": {
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "match"
        ],
        "summary": "Создание матча из записи",
        "operationId": "createMatch",
        "parameters": [
          {
            "type": "file",
            "description": "файл записи матча с расширением .dem",
            "name": "replay",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/CreateMatchResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "409": {
            "description": "Conflict",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/matches/{match_id}": {
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "match"
        ],
        "summary": "Удаление матча",
        "operationId": "deleteMatch",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "ID матча",
            "name": "match_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "No Content"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/players/{steam_id}/stats": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "player"
        ],
        "summary": "Получение статистики игрока",
        "operationId": "getPlayerStats",
        "parameters": [
          {
            "type": "string",
            "description": "Steam ID игрока",
            "name": "steam_id",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "format": "uuid",
            "description": "Фильтр по матчу",
            "name": "match_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/PlayerStats"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/players/{steam_id}/weapons": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "player"
        ],
        "summary": "Получение статистики игрока по оружию",
        "operationId": "getWeaponStats",
        "parameters": [
          {
            "type": "string",
            "description": "Steam ID игрока",
            "name": "steam_id",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "format": "uuid",
            "description": "Фильтр по матчу",
            "name": "match_id",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "Фильтр по оружию",
            "name": "weapon_id",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "Фильтр по классу оружия",
            "name": "class_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/PlayerWeaponStats"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/teams": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "team"
        ],
        "summary": "Получение списка команд",
        "operationId": "getTeamList",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/WeaponList"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CreateMatchResponse": {
      "type": "object",
      "required": [
        "match_id",
        "match_number"
      ],
      "properties": {
        "match_id": {
          "type": "string",
          "format": "uuid",
          "x-isnullable": false
        },
        "match_number": {
          "type": "number",
          "format": "int32",
          "x-isnullable": false
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "код ошибки, равен статусу или \u003e= 600",
          "type": "integer",
          "format": "int32",
          "x-nullable": false
        },
        "message": {
          "description": "сообщение ошибки",
          "type": "string",
          "x-nullable": false
        }
      }
    },
    "PlayerStats": {
      "type": "object",
      "required": [
        "calculated_stats",
        "round_stats",
        "total_stats"
      ],
      "properties": {
        "calculated_stats": {
          "$ref": "#/definitions/PlayerStats_calculated_stats"
        },
        "round_stats": {
          "$ref": "#/definitions/PlayerStats_round_stats"
        },
        "total_stats": {
          "$ref": "#/definitions/PlayerStats_total_stats"
        }
      }
    },
    "PlayerStats_calculated_stats": {
      "description": "высчитанная статистика на основе статистики по матчам",
      "type": "object",
      "properties": {
        "headshot_percentage": {
          "type": "number",
          "format": "double"
        },
        "kill_death_ratio": {
          "type": "number",
          "format": "double"
        },
        "win_rate": {
          "type": "number",
          "format": "double"
        }
      },
      "x-nullable": false
    },
    "PlayerStats_round_stats": {
      "description": "набор средних показателей за раунд",
      "type": "object",
      "properties": {
        "assists": {
          "description": "среднее кол-во ассистов за раунд",
          "type": "number",
          "format": "double"
        },
        "blinded_players": {
          "description": "средне кол-во ослепленных игроков за раунд",
          "type": "number",
          "format": "double"
        },
        "blinded_times": {
          "description": "среднее кол-во раз ослеплен за раунд",
          "type": "number",
          "format": "double"
        },
        "damage_dealt": {
          "description": "средний урон за раунд",
          "type": "number",
          "format": "double"
        },
        "deaths": {
          "description": "среднее кол-во смертей за раунд",
          "type": "number",
          "format": "double"
        },
        "grenade_damage_dealt": {
          "description": "средний урон гранатами за раунд",
          "type": "number",
          "format": "double"
        },
        "kills": {
          "description": "среднее кол-во убийств за раунд",
          "type": "number",
          "format": "double"
        }
      },
      "x-nullable": false
    },
    "PlayerStats_total_stats": {
      "description": "статистика игрока по всем сыгранным матчам",
      "type": "object",
      "properties": {
        "assists": {
          "type": "integer",
          "format": "int32"
        },
        "blind_kills": {
          "type": "integer",
          "format": "int32"
        },
        "blinded_players": {
          "type": "integer",
          "format": "int32"
        },
        "blinded_times": {
          "type": "integer",
          "format": "int32"
        },
        "bombs_defused": {
          "type": "integer",
          "format": "int32"
        },
        "bombs_planted": {
          "type": "integer",
          "format": "int32"
        },
        "damage_dealt": {
          "type": "integer",
          "format": "int32"
        },
        "damage_taken": {
          "type": "integer",
          "format": "int32"
        },
        "deaths": {
          "type": "integer",
          "format": "int32"
        },
        "draws": {
          "type": "integer",
          "format": "int32"
        },
        "flashbang_assists": {
          "type": "integer",
          "format": "int32"
        },
        "grenade_damage_dealt": {
          "type": "integer",
          "format": "int32"
        },
        "headshot_kills": {
          "type": "integer",
          "format": "int32"
        },
        "kills": {
          "type": "integer",
          "format": "int32"
        },
        "loses": {
          "type": "integer",
          "format": "int32"
        },
        "matches_played": {
          "type": "integer",
          "format": "int32"
        },
        "mvp_count": {
          "type": "integer",
          "format": "int32"
        },
        "noscope_kills": {
          "type": "integer",
          "format": "int32"
        },
        "rounds_played": {
          "type": "integer",
          "format": "int32"
        },
        "through_smoke_kills": {
          "type": "integer",
          "format": "int32"
        },
        "time_played": {
          "type": "integer",
          "format": "int64"
        },
        "wallbang_kills": {
          "type": "integer",
          "format": "int32"
        },
        "wins": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "PlayerWeaponStats": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/PlayerWeaponStats_inner"
      }
    },
    "PlayerWeaponStats_inner": {
      "type": "object",
      "required": [
        "accuracy_stats",
        "total_stats"
      ],
      "properties": {
        "accuracy_stats": {
          "$ref": "#/definitions/PlayerWeaponStats_inner_accuracy_stats"
        },
        "total_stats": {
          "$ref": "#/definitions/PlayerWeaponStats_inner_total_stats"
        }
      },
      "x-nullable": false
    },
    "PlayerWeaponStats_inner_accuracy_stats": {
      "type": "object",
      "properties": {
        "arms": {
          "type": "number",
          "format": "double"
        },
        "chest": {
          "type": "number",
          "format": "double"
        },
        "head": {
          "type": "number",
          "format": "double"
        },
        "legs": {
          "type": "number",
          "format": "double"
        },
        "neck": {
          "type": "number",
          "format": "double"
        },
        "stomach": {
          "type": "number",
          "format": "double"
        },
        "total": {
          "type": "number",
          "format": "double"
        }
      },
      "x-nullable": false
    },
    "PlayerWeaponStats_inner_total_stats": {
      "type": "object",
      "properties": {
        "assists": {
          "type": "integer",
          "format": "int32"
        },
        "blind_kills": {
          "type": "integer",
          "format": "int32"
        },
        "chest_hits": {
          "type": "integer",
          "format": "int32"
        },
        "damage_dealt": {
          "type": "integer",
          "format": "int32"
        },
        "damage_taken": {
          "type": "integer",
          "format": "int32"
        },
        "deaths": {
          "type": "integer",
          "format": "int32"
        },
        "head_hits": {
          "type": "integer",
          "format": "int32"
        },
        "headshot_kills": {
          "type": "integer",
          "format": "int32"
        },
        "kills": {
          "type": "integer",
          "format": "int32"
        },
        "left_arm_hits": {
          "type": "integer",
          "format": "int32"
        },
        "left_leg_hits": {
          "type": "integer",
          "format": "int32"
        },
        "neck_hits": {
          "type": "integer",
          "format": "int32"
        },
        "noscope_kills": {
          "type": "integer",
          "format": "int32"
        },
        "right_arm_hits": {
          "type": "integer",
          "format": "int32"
        },
        "right_leg_hits": {
          "type": "integer",
          "format": "int32"
        },
        "shots": {
          "type": "integer",
          "format": "int32"
        },
        "stomach_hits": {
          "type": "integer",
          "format": "int32"
        },
        "through_smoke_kills": {
          "type": "integer",
          "format": "int32"
        },
        "wallbang_kills": {
          "type": "integer",
          "format": "int32"
        },
        "weapon": {
          "type": "string"
        },
        "weapon_id": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "WeaponClassList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/WeaponClassList_inner"
      }
    },
    "WeaponClassList_inner": {
      "type": "object",
      "required": [
        "class",
        "id"
      ],
      "properties": {
        "class": {
          "type": "string",
          "x-nullable": false
        },
        "id": {
          "type": "integer",
          "format": "int32",
          "x-nullable": false
        }
      },
      "x-nullable": false
    },
    "WeaponList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/WeaponList_inner"
      }
    },
    "WeaponList_inner": {
      "type": "object",
      "required": [
        "class",
        "class_id",
        "weapon",
        "weapon_id"
      ],
      "properties": {
        "class": {
          "description": "имя класса оружия",
          "type": "string",
          "x-nullable": false
        },
        "class_id": {
          "type": "integer",
          "format": "int32",
          "x-nullable": false
        },
        "weapon": {
          "description": "название оружия",
          "type": "string",
          "x-nullable": false
        },
        "weapon_id": {
          "type": "integer",
          "format": "int32",
          "x-nullable": false
        }
      },
      "x-nullable": false
    }
  },
  "tags": [
    {
      "description": "Профиль игрока",
      "name": "player"
    },
    {
      "description": "Матч",
      "name": "match"
    },
    {
      "description": "Команда",
      "name": "team"
    },
    {
      "description": "Учебное заведение",
      "name": "institution"
    },
    {
      "description": "Справочник",
      "name": "compendium"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Uniplay private API",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/v1",
  "paths": {
    "/compendiums/weapon-classes": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "compendium"
        ],
        "summary": "Получение списка классов оружий",
        "operationId": "getWeaponClasses",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/WeaponClassList"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/compendiums/weapons": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "compendium"
        ],
        "summary": "Получение списка оружий",
        "operationId": "getWeapons",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/WeaponList"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/institutions": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "institution"
        ],
        "summary": "Получение списка учебных заведений",
        "operationId": "getInstitutions",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/WeaponList"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/matches": {
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "match"
        ],
        "summary": "Создание матча из записи",
        "operationId": "createMatch",
        "parameters": [
          {
            "type": "file",
            "description": "файл записи матча с расширением .dem",
            "name": "replay",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/CreateMatchResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "409": {
            "description": "Conflict",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/matches/{match_id}": {
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "match"
        ],
        "summary": "Удаление матча",
        "operationId": "deleteMatch",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "ID матча",
            "name": "match_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "No Content"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/players/{steam_id}/stats": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "player"
        ],
        "summary": "Получение статистики игрока",
        "operationId": "getPlayerStats",
        "parameters": [
          {
            "type": "string",
            "description": "Steam ID игрока",
            "name": "steam_id",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "format": "uuid",
            "description": "Фильтр по матчу",
            "name": "match_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/PlayerStats"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/players/{steam_id}/weapons": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "player"
        ],
        "summary": "Получение статистики игрока по оружию",
        "operationId": "getWeaponStats",
        "parameters": [
          {
            "type": "string",
            "description": "Steam ID игрока",
            "name": "steam_id",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "format": "uuid",
            "description": "Фильтр по матчу",
            "name": "match_id",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "Фильтр по оружию",
            "name": "weapon_id",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "Фильтр по классу оружия",
            "name": "class_id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/PlayerWeaponStats"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/teams": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "team"
        ],
        "summary": "Получение списка команд",
        "operationId": "getTeamList",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/WeaponList"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CreateMatchResponse": {
      "type": "object",
      "required": [
        "match_id",
        "match_number"
      ],
      "properties": {
        "match_id": {
          "type": "string",
          "format": "uuid",
          "x-isnullable": false
        },
        "match_number": {
          "type": "number",
          "format": "int32",
          "x-isnullable": false
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "код ошибки, равен статусу или \u003e= 600",
          "type": "integer",
          "format": "int32",
          "x-nullable": false
        },
        "message": {
          "description": "сообщение ошибки",
          "type": "string",
          "x-nullable": false
        }
      }
    },
    "PlayerStats": {
      "type": "object",
      "required": [
        "calculated_stats",
        "round_stats",
        "total_stats"
      ],
      "properties": {
        "calculated_stats": {
          "$ref": "#/definitions/PlayerStats_calculated_stats"
        },
        "round_stats": {
          "$ref": "#/definitions/PlayerStats_round_stats"
        },
        "total_stats": {
          "$ref": "#/definitions/PlayerStats_total_stats"
        }
      }
    },
    "PlayerStats_calculated_stats": {
      "description": "высчитанная статистика на основе статистики по матчам",
      "type": "object",
      "properties": {
        "headshot_percentage": {
          "type": "number",
          "format": "double"
        },
        "kill_death_ratio": {
          "type": "number",
          "format": "double"
        },
        "win_rate": {
          "type": "number",
          "format": "double"
        }
      },
      "x-nullable": false
    },
    "PlayerStats_round_stats": {
      "description": "набор средних показателей за раунд",
      "type": "object",
      "properties": {
        "assists": {
          "description": "среднее кол-во ассистов за раунд",
          "type": "number",
          "format": "double"
        },
        "blinded_players": {
          "description": "средне кол-во ослепленных игроков за раунд",
          "type": "number",
          "format": "double"
        },
        "blinded_times": {
          "description": "среднее кол-во раз ослеплен за раунд",
          "type": "number",
          "format": "double"
        },
        "damage_dealt": {
          "description": "средний урон за раунд",
          "type": "number",
          "format": "double"
        },
        "deaths": {
          "description": "среднее кол-во смертей за раунд",
          "type": "number",
          "format": "double"
        },
        "grenade_damage_dealt": {
          "description": "средний урон гранатами за раунд",
          "type": "number",
          "format": "double"
        },
        "kills": {
          "description": "среднее кол-во убийств за раунд",
          "type": "number",
          "format": "double"
        }
      },
      "x-nullable": false
    },
    "PlayerStats_total_stats": {
      "description": "статистика игрока по всем сыгранным матчам",
      "type": "object",
      "properties": {
        "assists": {
          "type": "integer",
          "format": "int32"
        },
        "blind_kills": {
          "type": "integer",
          "format": "int32"
        },
        "blinded_players": {
          "type": "integer",
          "format": "int32"
        },
        "blinded_times": {
          "type": "integer",
          "format": "int32"
        },
        "bombs_defused": {
          "type": "integer",
          "format": "int32"
        },
        "bombs_planted": {
          "type": "integer",
          "format": "int32"
        },
        "damage_dealt": {
          "type": "integer",
          "format": "int32"
        },
        "damage_taken": {
          "type": "integer",
          "format": "int32"
        },
        "deaths": {
          "type": "integer",
          "format": "int32"
        },
        "draws": {
          "type": "integer",
          "format": "int32"
        },
        "flashbang_assists": {
          "type": "integer",
          "format": "int32"
        },
        "grenade_damage_dealt": {
          "type": "integer",
          "format": "int32"
        },
        "headshot_kills": {
          "type": "integer",
          "format": "int32"
        },
        "kills": {
          "type": "integer",
          "format": "int32"
        },
        "loses": {
          "type": "integer",
          "format": "int32"
        },
        "matches_played": {
          "type": "integer",
          "format": "int32"
        },
        "mvp_count": {
          "type": "integer",
          "format": "int32"
        },
        "noscope_kills": {
          "type": "integer",
          "format": "int32"
        },
        "rounds_played": {
          "type": "integer",
          "format": "int32"
        },
        "through_smoke_kills": {
          "type": "integer",
          "format": "int32"
        },
        "time_played": {
          "type": "integer",
          "format": "int64"
        },
        "wallbang_kills": {
          "type": "integer",
          "format": "int32"
        },
        "wins": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "PlayerWeaponStats": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/PlayerWeaponStats_inner"
      }
    },
    "PlayerWeaponStats_inner": {
      "type": "object",
      "required": [
        "accuracy_stats",
        "total_stats"
      ],
      "properties": {
        "accuracy_stats": {
          "$ref": "#/definitions/PlayerWeaponStats_inner_accuracy_stats"
        },
        "total_stats": {
          "$ref": "#/definitions/PlayerWeaponStats_inner_total_stats"
        }
      },
      "x-nullable": false
    },
    "PlayerWeaponStats_inner_accuracy_stats": {
      "type": "object",
      "properties": {
        "arms": {
          "type": "number",
          "format": "double"
        },
        "chest": {
          "type": "number",
          "format": "double"
        },
        "head": {
          "type": "number",
          "format": "double"
        },
        "legs": {
          "type": "number",
          "format": "double"
        },
        "neck": {
          "type": "number",
          "format": "double"
        },
        "stomach": {
          "type": "number",
          "format": "double"
        },
        "total": {
          "type": "number",
          "format": "double"
        }
      },
      "x-nullable": false
    },
    "PlayerWeaponStats_inner_total_stats": {
      "type": "object",
      "properties": {
        "assists": {
          "type": "integer",
          "format": "int32"
        },
        "blind_kills": {
          "type": "integer",
          "format": "int32"
        },
        "chest_hits": {
          "type": "integer",
          "format": "int32"
        },
        "damage_dealt": {
          "type": "integer",
          "format": "int32"
        },
        "damage_taken": {
          "type": "integer",
          "format": "int32"
        },
        "deaths": {
          "type": "integer",
          "format": "int32"
        },
        "head_hits": {
          "type": "integer",
          "format": "int32"
        },
        "headshot_kills": {
          "type": "integer",
          "format": "int32"
        },
        "kills": {
          "type": "integer",
          "format": "int32"
        },
        "left_arm_hits": {
          "type": "integer",
          "format": "int32"
        },
        "left_leg_hits": {
          "type": "integer",
          "format": "int32"
        },
        "neck_hits": {
          "type": "integer",
          "format": "int32"
        },
        "noscope_kills": {
          "type": "integer",
          "format": "int32"
        },
        "right_arm_hits": {
          "type": "integer",
          "format": "int32"
        },
        "right_leg_hits": {
          "type": "integer",
          "format": "int32"
        },
        "shots": {
          "type": "integer",
          "format": "int32"
        },
        "stomach_hits": {
          "type": "integer",
          "format": "int32"
        },
        "through_smoke_kills": {
          "type": "integer",
          "format": "int32"
        },
        "wallbang_kills": {
          "type": "integer",
          "format": "int32"
        },
        "weapon": {
          "type": "string"
        },
        "weapon_id": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "WeaponClassList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/WeaponClassList_inner"
      }
    },
    "WeaponClassList_inner": {
      "type": "object",
      "required": [
        "class",
        "id"
      ],
      "properties": {
        "class": {
          "type": "string",
          "x-nullable": false
        },
        "id": {
          "type": "integer",
          "format": "int32",
          "x-nullable": false
        }
      },
      "x-nullable": false
    },
    "WeaponList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/WeaponList_inner"
      }
    },
    "WeaponList_inner": {
      "type": "object",
      "required": [
        "class",
        "class_id",
        "weapon",
        "weapon_id"
      ],
      "properties": {
        "class": {
          "description": "имя класса оружия",
          "type": "string",
          "x-nullable": false
        },
        "class_id": {
          "type": "integer",
          "format": "int32",
          "x-nullable": false
        },
        "weapon": {
          "description": "название оружия",
          "type": "string",
          "x-nullable": false
        },
        "weapon_id": {
          "type": "integer",
          "format": "int32",
          "x-nullable": false
        }
      },
      "x-nullable": false
    }
  },
  "tags": [
    {
      "description": "Профиль игрока",
      "name": "player"
    },
    {
      "description": "Матч",
      "name": "match"
    },
    {
      "description": "Команда",
      "name": "team"
    },
    {
      "description": "Учебное заведение",
      "name": "institution"
    },
    {
      "description": "Справочник",
      "name": "compendium"
    }
  ]
}`))
}
