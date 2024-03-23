package model



type DailyPoint struct {
	Id          int `sql:"AUTO_INCREMENT" gorm:"primary key"`
	Point       int
	Description string
}


