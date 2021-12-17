package mongorm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestToObjectID(t *testing.T) {
	t.Log(ToObjectID(""))
}


func TestObjToString(t *testing.T) {
	objId := primitive.ObjectID{1,2,3,4,5,6,7,8,9,1,3,23}
	t.Log(ToObjectID(objId.String()))
}
