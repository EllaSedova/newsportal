// Code generated by mfd-generator v0.4.4; DO NOT EDIT.

//nolint:all
//lint:file-ignore U1000 ignore unused code, it's generated
package db

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

const condition = "?.? = ?"

// base filters
type applier func(query *orm.Query) (*orm.Query, error)

type search struct {
	appliers []applier
}

func (s *search) apply(query *orm.Query) {
	for _, applier := range s.appliers {
		query.Apply(applier)
	}
}

func (s *search) where(query *orm.Query, table, field string, value interface{}) {
	query.Where(condition, pg.Ident(table), pg.Ident(field), value)
}

func (s *search) WithApply(a applier) {
	if s.appliers == nil {
		s.appliers = []applier{}
	}
	s.appliers = append(s.appliers, a)
}

func (s *search) With(condition string, params ...interface{}) {
	s.WithApply(func(query *orm.Query) (*orm.Query, error) {
		return query.Where(condition, params...), nil
	})
}

// Searcher is interface for every generated filter
type Searcher interface {
	Apply(query *orm.Query) *orm.Query
	Q() applier

	With(condition string, params ...interface{})
	WithApply(a applier)
}

type CategorySearch struct {
	search

	ID          *int
	Title       *string
	OrderNumber *int
	Alias       *string
	StatusID    *int
	IDs         []int
	NotID       *int
	TitleILike  *string
}

func (cs *CategorySearch) Apply(query *orm.Query) *orm.Query {
	if cs == nil {
		return query
	}
	if cs.ID != nil {
		cs.where(query, Tables.Category.Alias, Columns.Category.ID, cs.ID)
	}
	if cs.Title != nil {
		cs.where(query, Tables.Category.Alias, Columns.Category.Title, cs.Title)
	}
	if cs.OrderNumber != nil {
		cs.where(query, Tables.Category.Alias, Columns.Category.OrderNumber, cs.OrderNumber)
	}
	if cs.Alias != nil {
		cs.where(query, Tables.Category.Alias, Columns.Category.Alias, cs.Alias)
	}
	if cs.StatusID != nil {
		cs.where(query, Tables.Category.Alias, Columns.Category.StatusID, cs.StatusID)
	}
	if len(cs.IDs) > 0 {
		Filter{Columns.Category.ID, cs.IDs, SearchTypeArray, false}.Apply(query)
	}
	if cs.NotID != nil {
		Filter{Columns.Category.ID, *cs.NotID, SearchTypeEquals, true}.Apply(query)
	}
	if cs.TitleILike != nil {
		Filter{Columns.Category.Title, *cs.TitleILike, SearchTypeILike, false}.Apply(query)
	}

	cs.apply(query)

	return query
}

func (cs *CategorySearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		if cs == nil {
			return query, nil
		}
		return cs.Apply(query), nil
	}
}

type NewsSearch struct {
	search

	ID            *int
	Title         *string
	CategoryID    *int
	Foreword      *string
	Content       *string
	Author        *string
	PublishedAt   *time.Time
	StatusID      *int
	IDs           []int
	TitleILike    *string
	ForewordILike *string
	ContentILike  *string
	AuthorILike   *string
}

func (ns *NewsSearch) Apply(query *orm.Query) *orm.Query {
	if ns == nil {
		return query
	}
	if ns.ID != nil {
		ns.where(query, Tables.News.Alias, Columns.News.ID, ns.ID)
	}
	if ns.Title != nil {
		ns.where(query, Tables.News.Alias, Columns.News.Title, ns.Title)
	}
	if ns.CategoryID != nil {
		ns.where(query, Tables.News.Alias, Columns.News.CategoryID, ns.CategoryID)
	}
	if ns.Foreword != nil {
		ns.where(query, Tables.News.Alias, Columns.News.Foreword, ns.Foreword)
	}
	if ns.Content != nil {
		ns.where(query, Tables.News.Alias, Columns.News.Content, ns.Content)
	}
	if ns.Author != nil {
		ns.where(query, Tables.News.Alias, Columns.News.Author, ns.Author)
	}
	if ns.PublishedAt != nil {
		ns.where(query, Tables.News.Alias, Columns.News.PublishedAt, ns.PublishedAt)
	}
	if ns.StatusID != nil {
		ns.where(query, Tables.News.Alias, Columns.News.StatusID, ns.StatusID)
	}
	if len(ns.IDs) > 0 {
		Filter{Columns.News.ID, ns.IDs, SearchTypeArray, false}.Apply(query)
	}
	if ns.TitleILike != nil {
		Filter{Columns.News.Title, *ns.TitleILike, SearchTypeILike, false}.Apply(query)
	}
	if ns.ForewordILike != nil {
		Filter{Columns.News.Foreword, *ns.ForewordILike, SearchTypeILike, false}.Apply(query)
	}
	if ns.ContentILike != nil {
		Filter{Columns.News.Content, *ns.ContentILike, SearchTypeILike, false}.Apply(query)
	}
	if ns.AuthorILike != nil {
		Filter{Columns.News.Author, *ns.AuthorILike, SearchTypeILike, false}.Apply(query)
	}

	ns.apply(query)

	return query
}

func (ns *NewsSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		if ns == nil {
			return query, nil
		}
		return ns.Apply(query), nil
	}
}

type TagSearch struct {
	search

	ID         *int
	Title      *string
	StatusID   *int
	IDs        []int
	TitleILike *string
}

func (ts *TagSearch) Apply(query *orm.Query) *orm.Query {
	if ts == nil {
		return query
	}
	if ts.ID != nil {
		ts.where(query, Tables.Tag.Alias, Columns.Tag.ID, ts.ID)
	}
	if ts.Title != nil {
		ts.where(query, Tables.Tag.Alias, Columns.Tag.Title, ts.Title)
	}
	if ts.StatusID != nil {
		ts.where(query, Tables.Tag.Alias, Columns.Tag.StatusID, ts.StatusID)
	}
	if len(ts.IDs) > 0 {
		Filter{Columns.Tag.ID, ts.IDs, SearchTypeArray, false}.Apply(query)
	}
	if ts.TitleILike != nil {
		Filter{Columns.Tag.Title, *ts.TitleILike, SearchTypeILike, false}.Apply(query)
	}

	ts.apply(query)

	return query
}

func (ts *TagSearch) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		if ts == nil {
			return query, nil
		}
		return ts.Apply(query), nil
	}
}
