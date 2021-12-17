package mongorm

import (
	"go.mongodb.org/mongo-driver/bson"
)

// BindData 将data数据绑定到ret的结构体上，注意ret为指针类型
// data 数据看中的数据
// ret 指针类型的结构体
func BindData(data interface{}, ret interface{}) (err error) {
	bsonBytes, err := bson.Marshal(data)
	if err != nil {
		return
	}

	err = bson.Unmarshal(bsonBytes, ret)

	return
}
