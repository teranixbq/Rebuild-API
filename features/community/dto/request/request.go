package request

type CommunityRequest struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Location    string `form:"location"`
	Max_Members int    `form:"max_members"`
	Image       string `form:"image"`
}

type EventRequest struct {
	Title       string `form:"title"`
	Image       string `form:"image"`
	Description string `form:"description"`
	Location    string `form:"location"`
	MapLink     string `form:"maplink"`
	FormLink    string `form:"formlink"`
	Quota       int    `form:"quota"`
	Date        string `form:"date"`
	Status      string `form:"status"`
}
