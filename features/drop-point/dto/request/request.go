package request

type DropPointRequest struct {
	Name      string            `json:"name"`
	Address   string            `json:"address"`
	Latitude  float64           `json:"latitude"`
	Longitude float64           `json:"longitude"`
	Schedule  []ScheduleRequest `json:"schedule"`
}

type ScheduleRequest struct {
	Day        string `json:"day"`
	Open_Time  string `json:"open_time"`
	Close_Time string `json:"close_time"`
	Closed     bool   `json:"closed"`
}
