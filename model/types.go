package model

import "time"

type Account struct {
	AccountId     int    `json:"accountId"`
	WalletAccount string `json:"walletAccount"`
	NickName      string `json:"nickName"`
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
	DnftId          int     `json:"dnftId"`
	WalletAccount   string  `json:"walletAccount"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	DnftCountry     string  `json:"dnftCountry"`
	DnftImgUrl      string  `json:"dnftImgUrl"`
	DnftDescription string  `json:"dnftDescription"`
	DnftCount       int     `json:"dnftCount"`
}

type BadgeResource struct {
	BadgeId            int    `json:"badgeId"`
	BadgeCountry       string `json:"badgeCountry"`
	BadgeImgUrl_bronze string `json:"badgeImgUrl_bronze"`
	BadgeImgUrl_silver string `json:"badgeImgUrl_silver"`
	BadgeImgUrl_gold   string `json:"badgeImgUrl_gold"`
	BadgeDescription   string `json:"badgeDescription"`
}
