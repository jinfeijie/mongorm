package mongorm

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 这个文件用于组装查询的字段和排序条件

type Search struct {
	Projection string // 返回字段select
	Skip       int64
	Limit      int64
	Sort       bson.D
	Hint       string
}

type SortType int
type SearchSort struct {
	Field string
	SortType
}

const DESC = SortType(-1)
const ASC = SortType(1)

func NewSearch() *Search {
	return &Search{}
}

func SetOffset(offset int64) *Search {
	search := NewSearch()
	return search.SetOffset(offset)
}
func (s *Search) SetOffset(offset int64) *Search {
	s.Skip = offset
	return s
}

func SetLimit(limit int64) *Search {
	search := NewSearch()
	return search.SetLimit(limit)
}
func (s *Search) SetLimit(limit int64) *Search {
	s.Limit = limit
	return s
}

func SetSort(field string, sort SortType) *Search {
	search := NewSearch()
	return search.SetSort(field, sort)
}
func (s *Search) SetSort(field string, sort SortType) *Search {
	s.Sort = append(s.Sort, sortResolve(field, sort))
	return s
}

func SetHint(hint string) *Search {
	search := NewSearch()
	return search.SetHint(hint)
}
func (s *Search) SetHint(hint string) *Search {
	s.Hint = hint
	return s
}

func SetSelect(fields string) *Search {
	search := NewSearch()
	return search.SetSelect(fields)
}
func (s *Search) SetSelect(fields string) *Search {
	s.Projection = fields
	return s
}

func GetOption() *options.FindOptions {
	search := NewSearch()
	return search.GetOption()
}
func (s *Search) GetOption() *options.FindOptions {
	find := options.Find()

	if len(s.Projection) > 0 {
		find.Projection = projectionResolve(s.Projection)
	}

	if s.Skip > 0 {
		find.Skip = &s.Skip
	}

	if s.Limit > 0 {
		find.Limit = &s.Limit
	}

	if len(s.Sort) > 0 {
		find.Sort = s.Sort
	}

	if len(s.Hint) > 0 {
		find.Hint = s.Hint
	}
	return find
}

func GetOneOption() *options.FindOneOptions {
	search := NewSearch()
	return search.GetOneOption()
}
func (s *Search) GetOneOption() *options.FindOneOptions {
	find := options.FindOne()

	if len(s.Projection) > 0 {
		if projectionResolve(s.Projection) != nil {
			find.Projection = projectionResolve(s.Projection)
		}
	}

	if s.Skip > 0 {
		find.Skip = &s.Skip
	}

	if len(s.Sort) > 0 {
		find.Sort = s.Sort
	}

	if len(s.Hint) > 0 {
		find.Hint = s.Hint
	}
	return find
}
