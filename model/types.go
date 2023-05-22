package model

import "time"

type Account struct {
	WalletAccount string `json:"walletAccount"`
	NickName      string `json:"nickName"`
	Password      string `json:"password"`
	MyTravelCount int    `json:"myTravelCount"`
	MyDNFTCount   int    `json:"myDnftCount"`
	LikeCount     int    `json:"likeCount"`
	CommentCount  int    `json:"commentCount"`
}

type TravelPlan struct {
	TravelId          int       `json:"travelId"`
	TravelTitle       string    `json:"travelTitle"`
	TravelDescription string    `json:"travelDescription"`
	TravelMemo        string    `json:"travelMemo"`
	TravelImgSrc      string    `json:"travelImgSrc"`
	TravelTime        time.Time `json:"travelTime"`
}

type EncyclopediaDNFT struct {
	DnftId        int     `json:"dnftId"`
	WalletAccount string  `json:"walletAccount"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	DnftCountry   string  `json:"dnftCountry"`
	DnftImgUrl    string  `json:"dnftImgUrl"`
	IssueCount    int     `json:"issueCount"`
	DnftTime      string  `json:"dnfttime"`
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
