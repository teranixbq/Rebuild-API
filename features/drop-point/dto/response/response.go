package response

type DropPointResponse struct {
	Id        string             `json:"id"`
	Name      string             `json:"name"`
	Address   string             `json:"address"`
	Latitude  float64                  `json:"latitude"`
	Longitude float64                  `json:"longitude"`
	Schedule  []ScheduleResponse `json:"schedule"`
}

type ScheduleResponse struct {
	Day        string `json:"day"`
	Open_Time  string `json:"open_time"`
	Close_Time string `json:"close_time"`
	Closed     bool   `json:"closed"`
}

type DropPointDetailResponse struct {
	Id        string                   `json:"id"`
	Name      string                   `json:"name"`
	Address   string                   `json:"address"`
	Latitude  float64                  `json:"latitude"`
	Longitude float64                  `json:"longitude"`
	Schedule  []ScheduleDetailResponse `json:"schedule"`
}

type ScheduleDetailResponse struct {
	Day        string `json:"day"`
	Open_Time  string `json:"open_time"`
	Close_Time string `json:"close_time"`
	Closed     bool   `json:"closed"`
}
