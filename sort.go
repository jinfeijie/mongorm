package mongorm

import (
	"go.mongodb.org/mongo-driver/bson"
)

func sortResolve(field string, sortType SortType) bson.E {
	return bson.E{Key: field, Value: sortType}
}
