package mongorm

import (
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type Condition struct {
	condition bson.D
}

func NewCondition() *Condition {
	return &Condition{}
}

// Where 条件
func Where(field string, opt string, val interface{}) *Condition {
	return NewCondition().Where(field, opt, val)
}
func (con *Condition) Where(field string, opt string, val interface{}) *Condition {
	result := bson.E{}
	switch strings.ToLower(opt) {
	case "==":
		result = bson.E{Key: field, Value: val}
	case ">":
		result = bson.E{Key: field, Value: bson.D{{Key: "$gt", Value: val}}}
	case ">=":
		result = bson.E{Key: field, Value: bson.D{{"$gte", val}}}
	case "<":
		result = bson.E{Key: field, Value: bson.D{{"$lt", val}}}
	case "<=":
		result = bson.E{Key: field, Value: bson.D{{"$lte", val}}}
	case "!=":
		result = bson.E{Key: field, Value: bson.D{{"$ne", val}}}
	case "in":
		result = bson.E{Key: field, Value: bson.D{{"$in", val}}}
	case "nin", "not in":
		result = bson.E{Key: field, Value: bson.D{{"$nin", val}}}
	case "regex":
		result = bson.E{Key: field, Value: bson.D{{"$regex", val}}}
	default:
		result = bson.E{Key: field, Value: bson.D{{opt, val}}}
	}
	con.condition = append(con.condition, result)
	return con
}

// WhereOrResult 或条件
func WhereOrResult(condition *Condition, secondCondition ...*Condition) bson.M {
	return NewCondition().WhereOrResult(condition, secondCondition...)
}
func (con *Condition) WhereOrResult(condition *Condition, secondCondition ...*Condition) bson.M {
	var ret []interface{}
	ret = append(ret, condition.condition)
	for _, con := range secondCondition {
		ret = append(ret, con.condition)
	}
	return bson.M{"$or": ret}
}

// Result alias Cond
func (con *Condition) Result() bson.D {
	return con.Cond()
}

// Cond 输出查询条件
func (con *Condition) Cond() bson.D {
	return con.condition
}
