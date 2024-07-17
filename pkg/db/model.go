// nolint
//
//lint:file-ignore U1000 ignore unused code, it's generated
package db

import (
	"time"
)

var Columns = struct {
	Category struct {
		ID, Title, OrderNumber, Alias, StatusID string

		Status string
	}
	News struct {
		ID, Title, CategoryID, Foreword, Content, TagIDs, Author, PublishedAt, StatusID string

		Category, Status string
	}
	Status struct {
		ID, Title string
	}
	Tag struct {
		ID, Title, StatusID string

		Status string
	}
}{
	Category: struct {
		ID, Title, OrderNumber, Alias, StatusID string

		Status string
	}{
		ID:          "categoryId",
		Title:       "title",
		OrderNumber: "orderNumber",
		Alias:       "alias",
		StatusID:    "statusId",

		Status: "Status",
	},
	News: struct {
		ID, Title, CategoryID, Foreword, Content, TagIDs, Author, PublishedAt, StatusID string

		Category, Status string
	}{
		ID:          "newsId",
		Title:       "title",
		CategoryID:  "categoryId",
		Foreword:    "foreword",
		Content:     "content",
		TagIDs:      "tagIds",
		Author:      "author",
		PublishedAt: "publishedAt",
		StatusID:    "statusId",

		Category: "Category",
		Status:   "Status",
	},
	Status: struct {
		ID, Title string
	}{
		ID:    "statusId",
		Title: "title",
	},
	Tag: struct {
		ID, Title, StatusID string

		Status string
	}{
		ID:       "tagId",
		Title:    "title",
		StatusID: "statusId",

		Status: "Status",
	},
}

var Tables = struct {
	Category struct {
		Name, Alias string
	}
	News struct {
		Name, Alias string
	}
	Status struct {
		Name, Alias string
	}
	Tag struct {
		Name, Alias string
	}
}{
	Category: struct {
		Name, Alias string
	}{
		Name:  "categories",
		Alias: "t",
	},
	News: struct {
		Name, Alias string
	}{
		Name:  "news",
		Alias: "t",
	},
	Status: struct {
		Name, Alias string
	}{
		Name:  "statuses",
		Alias: "t",
	},
	Tag: struct {
		Name, Alias string
	}{
		Name:  "tags",
		Alias: "t",
	},
}

type Category struct {
	tableName struct{} `pg:"categories,alias:t,discard_unknown_columns"`

	ID          int    `pg:"categoryId,pk"`
	Title       string `pg:"title,use_zero"`
	OrderNumber *int   `pg:"orderNumber"`
	Alias       string `pg:"alias,use_zero"`
	StatusID    int    `pg:"statusId,use_zero"`

	Status *Status `pg:"fk:statusId,rel:has-one"`
}

type News struct {
	tableName struct{} `pg:"news,alias:t,discard_unknown_columns"`

	ID          int       `pg:"newsId,pk"`
	Title       string    `pg:"title,use_zero"`
	CategoryID  int       `pg:"categoryId,use_zero"`
	Foreword    string    `pg:"foreword,use_zero"`
	Content     *string   `pg:"content"`
	TagIDs      []int     `pg:"tagIds,array,use_zero"`
	Author      string    `pg:"author,use_zero"`
	PublishedAt time.Time `pg:"publishedAt,use_zero"`
	StatusID    int       `pg:"statusId,use_zero"`

	Category *Category `pg:"fk:categoryId,rel:has-one"`
	Status   *Status   `pg:"fk:statusId,rel:has-one"`
}

type Status struct {
	tableName struct{} `pg:"statuses,alias:t,discard_unknown_columns"`

	ID    int    `pg:"statusId,pk"`
	Title string `pg:"title,use_zero"`
}

type Tag struct {
	tableName struct{} `pg:"tags,alias:t,discard_unknown_columns"`

	ID       int    `pg:"tagId,pk"`
	Title    string `pg:"title,use_zero"`
	StatusID int    `pg:"statusId,use_zero"`

	Status *Status `pg:"fk:statusId,rel:has-one"`
}
