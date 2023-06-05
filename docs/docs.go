// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/v01/acc/profile": {
            "get": {
                "description": "유저 프로필 정보를 가져는 함수다. 다음과 같은 정보를 가져온다. [닉네임, 나의 여행계획 카운트, 나의 Dynamic NFT 카운트, 좋아요 카운트 , 댓글 카운트]",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserProfileHandler(나의 프로필 정보 가져오기)"
                ],
                "summary": "계정주소를 입력합니다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "walletAccount",
                        "name": "walletAccount",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "nickName",
                        "name": "nickName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Account"
                            }
                        }
                    }
                }
            }
        },
        "/v01/acc/register": {
            "post": {
                "description": "유저 닉네임을 등록 및 수정 해주는 함수로 지갑계정으로 연결 후 사용자가 닉네임을 입력할 수 있다. 이후 닉네임 수정은 자유롭게 가능하다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserRegisterHandler(닉네임 등록/수정)"
                ],
                "summary": "계정주소, 닉네임을 입력합니다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "walletAccount",
                        "name": "walletAccount",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "nickName",
                        "name": "nickName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Account"
                            }
                        }
                    }
                }
            }
        },
        "/v01/badge/issue": {
            "post": {
                "description": "위도, 경도를 입력받고 해당하는 나라의 리소스(ipfs uri, nft metadata)를 참고해서 뱃지를 발급해줍니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CreateBadge(뱃지 발급)"
                ],
                "summary": "지갑계정, 위도, 경도를 입력해줍니다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "walletAccount",
                        "name": "walletAccount",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "latitude",
                        "name": "latitude",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "longitude",
                        "name": "longitude",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.EncyclopediaDNFT"
                            }
                        }
                    }
                }
            }
        },
        "/v01/badge/user": {
            "get": {
                "description": "사용자 위치를 참고해서 뱃지를 발급하는 함수",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetMyBadge(나의 뱃지 가져오기)"
                ],
                "summary": "지갑계정을 입력해줍니다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "walletAccount",
                        "name": "walletAccount",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.EncyclopediaDNFT"
                            }
                        }
                    }
                }
            }
        },
        "/v01/trip/allplan": {
            "get": {
                "description": "모든 여행계획을  MongoDB에서 가져오는 함수. 아무 파라미터가 없다 전체를 조회한다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetAllTrip(전체 여행계획 가져오기)"
                ],
                "summary": "모든 여행계획을 가져온다.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TripPlan"
                            }
                        }
                    }
                }
            }
        },
        "/v01/trip/delete/:num": {
            "delete": {
                "description": "트립아이디를 입력하면 해당하는 게시물을 삭제합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DeleteTrip(나의 여행계획 중 디테일 페이지를 가져와줌)"
                ],
                "summary": "트립아이디를 delete 뒤에 입력하면 해당하는 트립 아이디의 게시물이 삭제됩니다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tripId",
                        "name": "tripId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TripPlan"
                            }
                        }
                    }
                }
            }
        },
        "/v01/trip/detail/:num": {
            "get": {
                "description": "트립아이디를 입력하면 해당하는 상세페이지를 보여줍니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetDetailTrip(나의 여행계획 중 디테일 페이지를 가져와줌)"
                ],
                "summary": "트립아이디를 detail 뒤에 입력하면 해당하는 트립 아이디가 출력됩니다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tripId",
                        "name": "tripId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TripPlan"
                            }
                        }
                    }
                }
            }
        },
        "/v01/trip/myplan": {
            "get": {
                "description": "나의 여행계획을 MongoDB에서 가져오는 함수, 계정주소로 파악한 후 가져온다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetMyTrip(나의 여행계획 가져오기)"
                ],
                "summary": "지갑계정을 입력한다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "walletAccount",
                        "name": "walletAccount",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TripPlan"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "days에는 day1, day2단위로 아이템이 있고 각 day1별로 시간과 imtes가 있으며 각각 여행시작시간, 종료시간, 이미지, 타이틀, 설명, 메모등을 입력할 수 있습니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CreateTripPlan(나의 여행계획표 생성하기)"
                ],
                "summary": "지갑계정, 제목, 나라, 출발날짜, 도착날짜 등을 입력합니다. days는 아이템을 담은 배열입니다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "walletAccount",
                        "name": "walletAccount",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tripTitle",
                        "name": "tripTitle",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tripCountry",
                        "name": "tripCountry",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tripDeparture",
                        "name": "tripDeparture",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tripArrival",
                        "name": "tripArrival",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "days",
                        "name": "days",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TripPlan"
                            }
                        }
                    }
                }
            }
        },
        "/v01/trip/search": {
            "get": {
                "description": "여행계획의 제목 중 일치하는 문자열에 대해 콘텐츠를 리스폰스해주는 검색 API, 단어 단위로 구현, 예를 들어 Paris로 무작정이라고 하면 \"Paris로\" 까지 입력해야된다. q=\"Paris로\" 이런식으로 입력하면 된다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SearchTrip(여행계획 단어단위 검색하기)"
                ],
                "summary": "q에 검색하고자 하는 키워드를 입력한다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "q",
                        "name": "q",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TripPlan"
                            }
                        }
                    }
                }
            }
        },
        "/v01/trip/simpleplan": {
            "post": {
                "description": "days에는 day1, day2단위로 아이템이 있고 각 day1별로 시간과 imtes가 있으며 각각 여행시작시간, 종료시간, 이미지, 타이틀, 설명, 메모등을 입력할 수 있습니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CreateTripPlan(나의 여행계획표 생성하기)"
                ],
                "summary": "지갑계정, 제목, 나라, 출발날짜, 도착날짜 등을 입력합니다. days는 아이템을 담은 배열입니다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "walletAccount",
                        "name": "walletAccount",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "triplTitle",
                        "name": "triplTitle",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tripCountry",
                        "name": "tripCountry",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tripDeparture",
                        "name": "tripDeparture",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tripArrival",
                        "name": "tripArrival",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "days",
                        "name": "days",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TripPlan"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "days에는 day1, day2단위로 아이템이 있고 각 day1별로 시간과 imtes가 있으며 각각 여행시작시간, 종료시간, 이미지, 타이틀, 설명, 메모등을 입력할 수 있습니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CreateTripPlan(나의 여행계획표 생성하기)"
                ],
                "summary": "지갑계정, 제목, 나라, 출발날짜, 도착날짜 등을 입력합니다. days를 입력으로 받는 함수 이후 데이터를 추가한다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tripId",
                        "name": "tripId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "days",
                        "name": "days",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TripPlan"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Account": {
            "type": "object",
            "properties": {
                "accountId": {
                    "description": "계정 ID",
                    "type": "integer"
                },
                "commentCount": {
                    "description": "댓글 카운트",
                    "type": "integer"
                },
                "likeCount": {
                    "description": "좋아요 카운트",
                    "type": "integer"
                },
                "myDnftCount": {
                    "description": "DNFT뱃지 카운트",
                    "type": "integer"
                },
                "myTripCount": {
                    "description": "여행계획 카운트",
                    "type": "integer"
                },
                "nickName": {
                    "description": "닉네임",
                    "type": "string"
                },
                "walletAccount": {
                    "description": "지갑 계정 주소",
                    "type": "string"
                }
            }
        },
        "model.DayItem": {
            "type": "object",
            "properties": {
                "endDate": {
                    "description": "세부 여행 종료날짜 \t/예시  : 15:15:00",
                    "type": "string"
                },
                "note": {
                    "description": "세부 여행 메모",
                    "type": "string"
                },
                "startDate": {
                    "description": "세부 여행 시작시간 \t/예시  : 09:15:00",
                    "type": "string"
                },
                "title": {
                    "description": "세부 여행 타이틀",
                    "type": "string"
                }
            }
        },
        "model.EncyclopediaDNFT": {
            "type": "object",
            "properties": {
                "dnftCountry": {
                    "description": "국가",
                    "type": "string"
                },
                "dnftId": {
                    "description": "전체 dnftID",
                    "type": "integer"
                },
                "dnftImgUrl": {
                    "description": "이미지URL로 쓸 변수",
                    "type": "string"
                },
                "dnftTier": {
                    "description": "티어",
                    "type": "string"
                },
                "dnftTime": {
                    "description": "발급 시간",
                    "type": "string"
                },
                "issueCount": {
                    "description": "계정별, 국가별 발급횟수를 체크하는 변수",
                    "type": "integer"
                },
                "latitude": {
                    "description": "위도",
                    "type": "number"
                },
                "longitude": {
                    "description": "경도",
                    "type": "number"
                },
                "walletAccount": {
                    "description": "지갑계정",
                    "type": "string"
                }
            }
        },
        "model.TripPlan": {
            "type": "object",
            "properties": {
                "dayItems": {
                    "description": "days 아이템",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.DayItem"
                    }
                },
                "nickName": {
                    "description": "닉네임",
                    "type": "string"
                },
                "tripArrival": {
                    "description": "여행출발",
                    "type": "string"
                },
                "tripCountry": {
                    "description": "여행국가",
                    "type": "string"
                },
                "tripDeparture": {
                    "description": "여행도착",
                    "type": "string"
                },
                "tripId": {
                    "description": "여행계획 아이디",
                    "type": "integer"
                },
                "tripTitle": {
                    "description": "여행계획 제목",
                    "type": "string"
                },
                "walletAccount": {
                    "description": "지갑주소 계정",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
