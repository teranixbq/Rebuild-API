package response

import "time"

type CommunityResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
}

type CommunityResponseForDetails struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	MaxMembers  int       `json:"max_members"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
}

type EventResponse struct {
	Id          string `json:"id"`
	CommunityId string `json:"communityId"`
	Title       string `json:"title"`
	Quota       int    `json:"quota"`
	Date        string `json:"date"`
	Status      string `json:"status"`
	Image       string `json:"image"`
}

type EventResponseDetail struct {
	Id          string `json:"id"`
	CommunityId string `json:"communityId"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Location    string `json:"location"`
	MapLink     string `json:"maplink"`
	FormLink    string `json:"formlink"`
	Quota       int    `json:"quota"`
	Date        string `json:"date"`
	Status      string `json:"status"`
}
