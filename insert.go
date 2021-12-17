package mongorm

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Insert struct {
	insert map[string]interface{}
}

func NewInsertOne() *Insert {
	insert := make(map[string]interface{})
	return &Insert{
		insert: insert,
	}
}

func (in *Insert) Value(field string, val interface{}) *Insert {
	in.insert[field] = val
	return in
}

func (in *Insert) Result(autoFields ...string) bson.D {
	for _, autoField := range autoFields {
		in.Value(autoField, time.Now().Unix())
	}
	return ToBsonD(in.insert)
}

// ResultRAW 输出未bson化前的内容
func (in *Insert) ResultRAW(autoFields ...string) map[string]interface{} {
	for _, autoField := range autoFields {
		in.Value(autoField, time.Now().Unix())
	}
	return in.insert
}
