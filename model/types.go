package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	WalletAccount string             `json:"walletAccount"`
	NickName      string             `json:"nickName"`
	Password      string             `json:"password"`
	MyTravelCount int                `json:"myTravelCount"`
	MyDNFTCount   int                `json:"myDnftCount"`
	LikeCount     int                `json:"likeCount"`
	CommentCount  int                `json:"commentCount"`
}

type TripPlan struct {
	TripId          int    `json:"travelId"`
	WalletAccount   string `bson:"walletAccount"`
	TripTitle       string `bson:"tripTitle"`
	TripDescription string `bson:"tripDescription"`
	TripMemo        string `bson:"tripMemo"`
	TripImgSrc      string `bson:"tripImgSrc"`
	TripTime        string `json:"tripTime"`
}

type EncyclopediaDNFTs struct {
	Arr []EncyclopediaDNFT
}

type EncyclopediaDNFT struct {
	DnftId        int64   `bson:"dnftId"`        //전체 dnftID
	WalletAccount string  `bson:"walletAccount"` //지갑계정
	Latitude      float64 `json:"latitude"`      //위도
	Longitude     float64 `json:"longitude"`     //경도
	DnftCountry   string  `bson:"dnftCountry"`   //국가
	DnftImgUrl    string  `json:"dnftImgUrl"`    //이미지URL로 쓸 변수
	DnftBronzeUrl string  `json:"dnftBronzeUrl"` //브론즈 URL
	DnftSilverUrl string  `json:"dnftSilverUrl"` //실버 URL
	DnftGoldUrl   string  `json:"dnftGoldUrl"`   //골드 URL
	DnftTime      string  `json:"dnfttime"`      //발급 시간
	BadgeTier     string  `json:"dnftTier"`      //티어
	IssueCount    int64   `bson:"issueCount"`    //계정별, 국가별 발급횟수를 체크하는 변수
}

type BadgeResource struct {
	BadgeId         int    `json:"badgeId"`
	BadgeCountry    string `json:"badgeCountry"`
	BadgeUrl_bronze string `json:"badgeUrl_bronze"`
	BadgeUrl_silver string `json:"badgeUrl_silver"`
	BadgeUrl_gold   string `json:"badgeUrl_gold"`
}

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
