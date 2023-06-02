package model

// 유저 계정
type Account struct {
	AccountID     int64  `json:"accountId"		, bson:"accountid"`        //계정 ID
	WalletAccount string `json:"walletAccount"	, bson:"walletaccount"` //지갑 계정 주소
	NickName      string `json:"nickName" 		, bson:"nickname"`         //닉네임
	MyTripCount   int64  `json:"myTripCount"	, bson:"mytripcount"`     //여행계획 카운트
	MyDNFTCount   int64  `json:"myDnftCount" 	, bson:"mydnftcount"`    //DNFT뱃지 카운트
	LikeCount     int64  `json:"likeCount"    	, bson:"likecount"`     //좋아요 카운트
	CommentCount  int64  `json:"commentCount" 	, bson:"commentcount"`  //댓글 카운트
}

// 여행계획
type TripPlan struct {
	TripId        int64     `json:"tripId"				,bson:"tripid"`             //여행계획 아이디
	WalletAccount string    `json:"walletAccount"		,bson:"walletaccount"` //지갑주소 계정
	NickName      string    `json:"nickName"			,bson:"nickname"`          //닉네임
	TripTitle     string    `json:"tripTitle"			,bson:"triptitle"`        //여행계획 제목
	TripCountry   string    `json:"tripCountry"			,bson:"tripcountry"`    //여행국가
	TripDeparture string    `json:"tripDeparture"		,bson:"tripdeparture"` //여행도착
	TripArrival   string    `json:"tripArrival"			,bson:"triparrival"`    //여행출발
	DayItems      []DayItem `json:"dayItems"				,bson:"dayitems"`         //days 아이템
}

type DayItem struct {
	StartDate string `json:"startDate"		,bson:"startdate"` //세부 여행 시작시간 	/예시  : 09:15:00
	EndDate   string `json:"endDate"		,bson:"enddate"`     //세부 여행 종료날짜 	/예시  : 15:15:00
	Title     string `json:"title"			,bson:"title"`        //세부 여행 타이틀
	Note      string `json:"note"			,bson:"note"`          //세부 여행 메모
}

// 다이나믹 NFT 구조체
type EncyclopediaDNFT struct {
	DnftId        int64   `json:"dnftId"  		 ,bson:"dnftid"`            //전체 dnftID
	WalletAccount string  `json:"walletAccount"	 ,bson:"walletaccount"` //지갑계정
	Latitude      float64 `json:"latitude"       ,bson:"latitude"`      //위도
	Longitude     float64 `json:"longitude"      ,bson:"longitude"`     //경도
	DnftCountry   string  `json:"dnftCountry"    ,bson:"dnftcountry"`   //국가
	DnftImgUrl    string  `json:"dnftImgUrl"     ,bson:"dnftimgUrl"`    //이미지URL로 쓸 변수
	DnftBronzeUrl string  `json:"dnftBronzeUrl"	 ,bson:"dnftbronzeUrl"` //브론즈 URL
	DnftSilverUrl string  `json:"dnftSilverUrl"  ,bson:"dnftsilverUrl"` //실버 URL
	DnftGoldUrl   string  `json:"dnftGoldUrl"    ,bson:"dnftgoldUrl"`   //골드 URL
	DnftTime      string  `json:"dnftTime"		 ,bson:"dnfttime"`          //발급 시간
	BadgeTier     string  `json:"dnftTier"       ,bson:"dnfttier"`      //티어
	IssueCount    int64   `json:"issueCount"	 ,bson:"issuecount"`       //계정별, 국가별 발급횟수를 체크하는 변수
}

// 뱃지 리소스
type BadgeResource struct {
	BadgeId        int    `json:"badgeId"			,bson:"badgeId"`              //뱃지 아이디
	BadgeCountry   string `json:"badgeCountry" 	,bson:"badgeCountry"`     //뱃지 국가
	BadgeUrlBronze string `json:"badgeUrlBronze	,bson:"badgeUrlBronze""`  //브론즈 뱃지
	BadgeUrlSilver string `json:"badgeUrlSilver"  ,bson:"badgeUrlSilver"` //실버 뱃지
	BadgeUrlGold   string `json:"badgeUrlGold"	,bson:"badgeUrlGold"`      //골드 뱃지
}

// 로케이션 리소스 구조체

type Location struct {
	FormattedAddress string `json:"formattedAddress"	,bson:"formattedAddress"`
	Street           string `json:"street"				,bson:"street"`
	HouseNumber      string `json:"houseNumber"			,bson:"houseNumber"`
	Suburb           string `json:"suburb"				,bson:"suburb"`
	Postcode         string `json:"postcode"			,bson:"postcode"`
	State            string `json:"state"				,bson:"state"`
	StateCode        string `json:"stateCode"			,bson:"stateCode"`
	StateDistrict    string `json:"stateDistrict"		,bson:"stateDistrict"`
	County           string `json:"county"				,bson:"county"`
	Country          string `json:"country"				,bson:"country"`
	CountryCode      string `json:"countryCode"			,bson:"countryCode"`
	City             string `json:"city"				,bson:"city"`
}

var ABI = `[
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_interval",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "approved",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			}
		],
		"name": "Approval",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "operator",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "bool",
				"name": "approved",
				"type": "bool"
			}
		],
		"name": "ApprovalForAll",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			}
		],
		"name": "approve",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_fromTokenId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_toTokenId",
				"type": "uint256"
			}
		],
		"name": "BatchMetadataUpdate",
		"type": "event"
	},
	{
		"inputs": [],
		"name": "eventDrop",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_tokenId",
				"type": "uint256"
			}
		],
		"name": "increasebadgeLevel",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "_tokenId",
				"type": "uint256"
			}
		],
		"name": "MetadataUpdate",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "previousOwner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "OwnershipTransferred",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "bytes",
				"name": "",
				"type": "bytes"
			}
		],
		"name": "performUpkeep",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "renounceOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "string",
				"name": "uri",
				"type": "string"
			}
		],
		"name": "safeMint",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			}
		],
		"name": "safeTransferFrom",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			},
			{
				"internalType": "bytes",
				"name": "data",
				"type": "bytes"
			}
		],
		"name": "safeTransferFrom",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "operator",
				"type": "address"
			},
			{
				"internalType": "bool",
				"name": "approved",
				"type": "bool"
			}
		],
		"name": "setApprovalForAll",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			}
		],
		"name": "Transfer",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			}
		],
		"name": "transferFrom",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "transferOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "upgrade",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_tokenId",
				"type": "uint256"
			},
			{
				"internalType": "string",
				"name": "_uri",
				"type": "string"
			}
		],
		"name": "upgradeBadge",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "badgeLevel",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			}
		],
		"name": "balanceOf",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bytes",
				"name": "",
				"type": "bytes"
			}
		],
		"name": "checkUpkeep",
		"outputs": [
			{
				"internalType": "bool",
				"name": "upkeepNeeded",
				"type": "bool"
			},
			{
				"internalType": "bytes",
				"name": "",
				"type": "bytes"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			}
		],
		"name": "getApproved",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "getLatestPrice",
		"outputs": [
			{
				"internalType": "int256",
				"name": "",
				"type": "int256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_address",
				"type": "address"
			}
		],
		"name": "getNftsOf",
		"outputs": [
			{
				"internalType": "string[]",
				"name": "",
				"type": "string[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "operator",
				"type": "address"
			}
		],
		"name": "isApprovedForAll",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "name",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			}
		],
		"name": "ownerOf",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bytes4",
				"name": "interfaceId",
				"type": "bytes4"
			}
		],
		"name": "supportsInterface",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "symbol",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			}
		],
		"name": "tokenURI",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`
