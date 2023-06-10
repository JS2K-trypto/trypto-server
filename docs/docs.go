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
                "description": "Fetches user profile information. Gets the following information [nickname, number of travel plans, number of Dynamic NFTs, number of likes, number of comments].",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserProfileHandler(Get my profile information)"
                ],
                "summary": "Enter your account address to get your profile information.",
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
                "description": "This function allows you to register and edit user nicknames, and you can do so by entering your wallet account and entering the desired nickname.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserRegisterHandler( You can register or edit your nickname. )"
                ],
                "summary": "Enter your account address and nickname.",
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
                "description": "It takes the latitude, longitude, and references the country's resources (IPFS URI, NFT METADATA) to issue the badge.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CreateBadge (issuing a badge)"
                ],
                "summary": "Enter your mumbai wallet account, latitude, and longitude.",
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
                "description": "Get all my DYNAMIC NFT badges",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetMyBadgeAll(get all my dynamic nft badges)"
                ],
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
                "description": "Function to fetch all trip plans from MongoDB. No parameters. Retrieves all.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetAllTrip(Import all  trip plans.)"
                ],
                "summary": "Import all trip plans.",
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
                "description": "Enter the /delete path followed by the trip id and the corresponding trip will be deleted.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DeleteTrip(Delete my trip)"
                ],
                "summary": "Entering a tripid after delete will delete posts for that tripid.",
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
                "description": "If you enter a trip id after the detail path, it will display the detail page of my trip plan.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetDetailTrip(Prints the details page of my trip plan)"
                ],
                "summary": "Entering a tripid after detail will output the corresponding trip plan. ex) /detail/tripId",
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
                "description": "Import the itinerary you created from MongoDB.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetMyTrip(Importing my trip plans)"
                ],
                "summary": "Enter your wallet account to import your trip plans",
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
                "description": "DAYS has items in DAY1 and DAY2, and each DAY1 has time and imtes, and you can enter start time, end time, image, title, description, and notes. After inputting, a travel plan is created.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CreateTripPlan(Create my itinerary)"
                ],
                "summary": "Enter the wallet account, title, country, departure date, arrival date, etc. days is an array containing the items.",
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
                        "description": "dayItems",
                        "name": "dayItems",
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
                "description": "A search API that returns content for matching strings in the title of a travel plan, implemented on a word-by-word basis, e.g. q=\"South Korea\".",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SearchTrip(Search your travel plans by keyword)"
                ],
                "summary": "In Q, type the keyword you want to search for.",
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
                "description": "DAYS has items in DAY1 and DAY2, and each DAY1 has time and imtes, and you can enter start time, end time, image, title, description, and notes.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CreateSimpleTripPlan(Create my simple trip plan)"
                ],
                "summary": "Enter the wallet account, title, country, departure date, arrival date, etc. days is an array containing the items. Here it is entered as an empty array.",
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
                        "description": "dayItems",
                        "name": "dayItems",
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
                "description": "This API is used to make modifications such as patches based on existing trip IDs.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PatchSimpleTripPlan(Modify my simple trip plan with the patch)"
                ],
                "summary": "Enter the wallet account, title, country, departure date, arrival date, etc. Add the data after the function that takes dayItems as input.",
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
                        "description": "dayItems",
                        "name": "dayItems",
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
