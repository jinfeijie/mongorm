package mongorm

import (
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"time"
)

type Update struct {
	update map[string]map[string]interface{}
}

func NewUpdate() *Update {
	up := make(map[string]map[string]interface{})
	return &Update{
		update: up,
	}
}

// Set 设置更新的字段和内容
func Set(field string, val interface{}) *Update {
	return NewUpdate().Set(field, val)
}
func (up *Update) Set(field string, val interface{}) *Update {
	up.data(field, Eq, val)
	return up
}

// SetAll 设置从map里面批量传入k=>v的参数
func (up *Update) SetAll(dt map[string]interface{}) *Update {
	for field, value := range dt {
		up.data(field, Eq, value)
	}
	return up
}

// Sset 特殊的操作
func (up *Update) Sset(opt string, field string, val interface{}) *Update {
	up.data(field, opt, val)
	return up
}

// Incr 字段加val
func Incr(field string, val int) *Update {
	up := NewUpdate()
	return up.Incr(field, val)
}
func (up *Update) Incr(field string, val int) *Update {
	up.data(field, Add, val)
	return up
}

func (up *Update) Result(autoFields ...string) bson.D {
	for _, autoField := range autoFields {
		up.Set(autoField, time.Now().Unix())
	}
	return ToBsonD(up.update)
}
func (up *Update) data(field string, opt string, val interface{}) *Update {
	if _, ok := up.update[resolveOpt(opt)]; !ok {
		up.update[resolveOpt(opt)] = make(map[string]interface{})
	}
	up.update[resolveOpt(opt)][field] = val
	return up
}

func resolveOpt(opt string) string {
	switch strings.ToLower(opt) {
	case Eq:
		return "$set"
	case Add:
		fallthrough
	case "add":
		fallthrough
	case "incr":
		fallthrough
	case "inc":
		fallthrough
	case "$inc":
		fallthrough
	case "$incr":
		return "$inc"
	default:
		return opt
	}
}
