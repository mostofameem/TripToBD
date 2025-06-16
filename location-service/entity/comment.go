package entity

import "time"

type Comment struct {
	Userid     int
	UserName   string
	Content    string
	Created_at time.Time
	Updated_at time.Time
	Vote       int
}
