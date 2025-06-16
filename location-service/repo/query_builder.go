package repo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoQueryBuilder struct {
	filter     bson.M
	projection bson.M
	opts       *options.FindOptions
}

func NewMongoQueryBuilder() *MongoQueryBuilder {
	return &MongoQueryBuilder{
		filter:     bson.M{},
		projection: bson.M{},
		opts:       options.Find(),
	}
}

func (b *MongoQueryBuilder) AddObjectID(field, value string) *MongoQueryBuilder {
	if value != "" {
		id, err := primitive.ObjectIDFromHex(value)
		if err == nil {
			b.filter[field] = id
		}
	}
	return b
}

func (b *MongoQueryBuilder) AddRegex(field, value string) *MongoQueryBuilder {
	if value != "" {
		b.filter[field] = bson.M{"$regex": value, "$options": "i"}
	}
	return b
}

func (b *MongoQueryBuilder) AddEqual(field string, value interface{}) *MongoQueryBuilder {
	if value != nil && value != "" {
		b.filter[field] = value
	}
	return b
}

func (b *MongoQueryBuilder) SetPagination(page, limit int) *MongoQueryBuilder {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	skip := int64((page - 1) * limit)
	limit64 := int64(limit)
	b.opts.SetSkip(skip)
	b.opts.SetLimit(limit64)
	return b
}

func (b *MongoQueryBuilder) SetSorting(sortBy, sortOrder string) *MongoQueryBuilder {
	order := 1
	if sortOrder == "desc" {
		order = -1
	}
	if sortBy != "" {
		b.opts.SetSort(bson.D{{Key: sortBy, Value: order}})
	}
	return b
}

func (b *MongoQueryBuilder) SetProjection(fields ...string) *MongoQueryBuilder {
	for _, field := range fields {
		b.projection[field] = 1
	}
	return b
}

func (b *MongoQueryBuilder) Build() (bson.M, *options.FindOptions) {
	if len(b.projection) > 0 {
		b.opts.SetProjection(b.projection)
	}
	return b.filter, b.opts
}
