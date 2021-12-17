package mongorm

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
)

// ToMap bson.M 转 map
func ToMap(m bson.M) map[string]interface{} {
	bs, err := json.Marshal(m)
	if err != nil {
		return nil
	}

	td := make(map[string]interface{})
	if err := json.Unmarshal(bs, &td); err != nil {
		return nil
	}
	return td
}

// ToBsonD 数据转BsonD格式
func ToBsonD(data interface{}) bson.D {
	var (
		bs  interface{}
		err error
	)

	bs, err = bson.Marshal(data)
	if err != nil {
		return bson.D{}
	}

	var d interface{}
	err = bson.Unmarshal(bs.([]byte), &d)
	if err != nil {
		return bson.D{}
	}
	return d.(bson.D)
}
