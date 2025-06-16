package location

import "time"

type AddLocationReq struct {
	Title        string
	Descriptions string
	BestTime     string
	PictureUrl   string
	Rating       float32
	Voted        int
	CreatedBy    int
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

type GetPageWithFilter struct {
	ID        *string
	Title     *string
	BestTime  *string
	Page      int
	Limit     int
	SortBy    string
	SortOrder string
}
type LocationPageResponse struct {
	Id         string  `json:"id"               bson:"_id,omitempty"`
	Title      string  `json:"title"            bson:"title"`
	PictureUrl string  `json:"picture_url"      bson:"picture_url"`
	Rating     float32 `json:"rating"           bson:"rating"`
}
type Comment struct {
	Userid     int
	UserName   string
	Content    string
	Created_at time.Time
	Updated_at time.Time
	Vote       int
}
