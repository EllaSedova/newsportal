package newsportal

import "time"

type NewsSummary struct {
	ID          int
	Title       string
	CategoryID  int
	Foreword    string
	Content     *string
	TagIDs      []int
	Author      string
	PublishedAt time.Time
	StatusID    int
	Category    *Category
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
