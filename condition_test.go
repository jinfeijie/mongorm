package mongorm

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestWhere(t *testing.T) {
	t.Log(Where("id", "in", []int{1, 2, 3}).Cond())
}

func TestWhereOrResult(t *testing.T) {
	condition := &Condition{
		condition: bson.D{
			{"foo", "bar"},
		},
	}
	secondCondition := &Condition{
		condition: bson.D{
			{"fooOr", "barOr"},
		},
	}

	t.Log(WhereOrResult(condition, secondCondition))
	t.Log(ToMap(WhereOrResult(condition, secondCondition)))
}
