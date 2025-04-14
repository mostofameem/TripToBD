package mongodb

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/exp/slog"
)

type Author struct {
	UserID   int    `json:"userid" bson:"userid"`
	Username string `json:"username" bson:"username"`
}

type Restaurant struct {
	ID           string    `json:"id" bson:"id"`
	Title        string    `json:"title" bson:"title"`
	Descriptions string    `json:"content" bson:"content"`
	PictureUrl   string    `json:"picture_url" bson:"picture_url"`
	Rating       float32   `json:"rating" bson:"rating"`
	Author       Author    `json:"author" bson:"author"`
	Branches     []int     `json:"branches" bson:"branches"`
	Comments     []Comment `json:"comments" bson:"comments"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}

func (repo *RestaurantTypeRepo) AddRestaurent(restaurant *Restaurant) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	restaurant.Branches = []int{}
	restaurant.Comments = []Comment{}

	mongoDoc, err := BsonM(restaurant)
	if err != nil {
		slog.Error("Error converting struct to BSON", err)
		return err
	}

	key := GetKey(restaurant.ID)
	mongoDoc["_id"] = key

	_, err = repo.collection.InsertOne(ctx, mongoDoc)
	if err != nil {
		slog.Error("Error inserting document", err)
		return err
	}

	return nil
}

func GetKey(id string) string {
	return fmt.Sprintf("posts:%s", id)
}

func BsonM(data interface{}) (bson.M, error) {
	result := bson.M{}

	val := reflect.Indirect(reflect.ValueOf(data))
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Tag.Get("bson")

		if fieldName == "" {
			fieldName = typ.Field(i).Name
		}

		if val.Field(i).CanInterface() {
			result[fieldName] = field.Interface()
		}
	}

	return result, nil
}
