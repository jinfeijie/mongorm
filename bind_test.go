package mongorm

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func TestBindData(t *testing.T) {
	data := bson.M{
		"name": "张三",
		"age":  10,
	}
	var u User
	if err := BindData(data, &u); err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(u.Name, u.Age)
}
