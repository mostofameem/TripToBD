package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	Id           primitive.ObjectID `json:"_id,omitempty"    bson:"_id,omitempty"`
	Title        string             `json:"title"            bson:"title"`
	Descriptions string             `json:"descriptions"     bson:"descriptions"`
	BestTime     string             `json:"best_time"        bson:"best_time"`
	PictureUrl   string             `json:"picture_url"      bson:"picture_url"`
	Rating       float32            `json:"rating"           bson:"rating"`
	VoteCount    int                `json:"vote_count"       bson:"vote_count"`
	CreatedBy    int                `json:"created_by"       bson:"created_by"`
	CreatedAt    time.Time          `json:"createdAt"        bson:"created_at"`
	UpdatedAt    *time.Time         `json:"updatedAt"        bson:"updated_at"`
	Restaurants  *[]int             `json:"restaurants"      bson:"restaurants"`
	Comments     *[]Comment         `json:"comments"         bson:"comments"`
}
