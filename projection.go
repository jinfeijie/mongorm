package mongorm

import (
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func projectionResolve(field string) bson.D {
	if field == "*" {
		return nil
	}

	fields := strings.Split(field, ",")
	projection := bson.D{}
	for _, f := range fields {
		data := bson.E{Key: f, Value: 1}
		projection = append(projection, data)
	}
	return projection
}
