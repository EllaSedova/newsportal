package newsportal

import (
	"newsportal/pkg/db"
)

type News struct {
	*db.News
	Category *Category
	Tags     []Tag
}

type Category struct {
	ID          int
	Title       string
	OrderNumber *int
	Alias       string
	StatusID    int
}

type Tag struct {
	ID       int
	Title    string
	StatusID int
}
