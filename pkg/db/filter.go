package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
)

type Filter struct {
	Field string
	Value interface{}
}

type Sort struct {
	Field     string
	Ascending bool
}

type QueryBuilder struct {
	Filters    []Filter
	NewFilters []Filter
	Sorts      []Sort
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (qb *QueryBuilder) AddFilter(categoryID, tagID *int) {
	if categoryID != nil {
		qb.AddFilterEqual(Columns.News.CategoryID, *categoryID)
	}
	if tagID != nil {
		qb.AddFilterAny(Columns.News.TagIDs, *tagID)
	}
}
func (qb *QueryBuilder) AddFilterEqual(field string, value interface{}) *QueryBuilder {
	qb.Filters = append(qb.Filters, Filter{Field: field, Value: value})
	return qb
}

func (qb *QueryBuilder) AddFilterAny(field string, value interface{}) *QueryBuilder {
	qb.NewFilters = append(qb.NewFilters, Filter{Field: field, Value: value})
	return qb
}

func (qb *QueryBuilder) AddSort(field string, ascending bool) *QueryBuilder {
	qb.Sorts = append(qb.Sorts, Sort{Field: field, Ascending: ascending})
	return qb
}

func (qb *QueryBuilder) Apply(query *pg.Query) *pg.Query {
	for _, filter := range qb.Filters {
		query.Where(fmt.Sprintf(`t."%s" = ?`, filter.Field), filter.Value)
	}
	for _, newFilter := range qb.NewFilters {
		query.Where(fmt.Sprintf(`? = ANY (t."%s")`, newFilter.Field), newFilter.Value)
	}
	for _, sort := range qb.Sorts {
		order := "ASC"
		if !sort.Ascending {
			order = "DESC"
		}
		query.Order(fmt.Sprintf(`t.%s %s`, sort.Field, order))
	}
	return query
}
