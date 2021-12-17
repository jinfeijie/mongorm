package mongorm

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestToMap(t *testing.T) {
	data := bson.M{
		"kkk":   "!212",
		"sds":   12,
		"dsads": "dsasdssd",
	}

	t.Log(ToMap(data))
}
