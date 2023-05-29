package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 유저 계정
type Account struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	WalletAccount string             `json:"walletAccount"`
	NickName      string             `json:"nickName"`
	MyTravelCount int                `json:"myTravelCount"`
	MyDNFTCount   int                `json:"myDnftCount"`
	LikeCount     int                `json:"likeCount"`
	CommentCount  int                `json:"commentCount"`
}

type TripPlans struct {
	Arr []TripPlan
}

// 여행계획
type TripPlan struct {
	TripId        int                `json:"travelId", 		bson:"travelId"`
	WalletAccount string             `json:"walletAccount", 	bson:"walletAccount"`
	TripTitle     string             `json:"tripTitle", 		bson:"tripTitle"`
	TripCountry   string             `json:"tripCountry",    bson:"tripCountry"`
	TripDeparture string             `json:"tripDeparture", 	bson:"tripDeparture"`
	TripArrival   string             `json:"tripArrival", 	bson:"tripArrival"`
	Days          map[string]DayItem `json:"days", 			bson:"days"`
}

type DayItem struct {
	DayTime string          `json:"dayTime", 	bson:"dayTime"`
	Items   map[string]Item `json:"items", 		bson:"items"`
}

type Item struct {
	StartDate       string `json:"startDate", 	bson:"endDate"`
	EndDate         string `json:"endDate",		bson:"endDate"`
	ImgSrc          string `json:"imgSrc",		bson:"imgSrc"`
	ItemTitle       string `json:"title", 		bson:"title"`
	ItemDescription string `json:"description", bson:"description"`
	ItemMemo        string `json:"memo", 		bson:"memo"`
}

// 다이나믹 NFT 배열
type EncyclopediaDNFTs struct {
	Arr []EncyclopediaDNFT
}

// 다이나믹 NFT 구조체
type EncyclopediaDNFT struct {
	DnftId        int64   `bson:"dnftId", json:"dnftId"`               //전체 dnftID
	WalletAccount string  `bson:"walletAccount", json:"walletAccount"` //지갑계정
	Latitude      float64 `bson:"latitude",json:"latitude"`            //위도
	Longitude     float64 `bson:"longitude",json:"longitude"`          //경도
	DnftCountry   string  `bson:"dnftCountry",json:"dnftCountry"`      //국가
	DnftImgUrl    string  `bson:"dnftImgUrl",json:"dnftImgUrl"`        //이미지URL로 쓸 변수
	DnftBronzeUrl string  `bson:"dnftBronzeUrl",json:"dnftBronzeUrl"`  //브론즈 URL
	DnftSilverUrl string  `bson:"dnftSilverUrl",json:"dnftSilverUrl"`  //실버 URL
	DnftGoldUrl   string  `bson:"dnftGoldUrl",json:"dnftGoldUrl"`      //골드 URL
	DnftTime      string  `bson:"dnfttime",json:"dnfttime"`            //발급 시간
	BadgeTier     string  `bson:"dnftTier",json:"dnftTier"`            //티어
	IssueCount    int64   `bson:"issueCount",bson:"issueCount"`        //계정별, 국가별 발급횟수를 체크하는 변수
}

// 뱃지 리소스
type BadgeResource struct {
	BadgeId         int    `json:"badgeId"`
	BadgeCountry    string `json:"badgeCountry"`
	BadgeUrl_bronze string `json:"badgeUrl_bronze"`
	BadgeUrl_silver string `json:"badgeUrl_silver"`
	BadgeUrl_gold   string `json:"badgeUrl_gold"`
}

// 로케이션 리소스 구조체

type Location struct {
	FormattedAddress string `json:"formattedAddress"`
	Street           string `json:"street"`
	HouseNumber      string `json:"houseNumber"`
	Suburb           string `json:"suburb"`
	Postcode         string `json:"postcode"`
	State            string `json:"state"`
	StateCode        string `json:"statecode"`
	StateDistrict    string `json:"stateDistrict"`
	County           string `json:"county"`
	Country          string `json:"country"`
	CountryCode      string `json:"countryCode"`
	City             string `json:"city"`
}

// Error implements error
func (*Account) Error() string {
	panic("unimplemented")
}
