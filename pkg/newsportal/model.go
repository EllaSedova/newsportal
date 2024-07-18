package newsportal

import "time"

type NewsSummary struct {
	ID          int       `json:"newsId"`
	Title       string    `json:"title"`
	CategoryID  int       `json:"categoryId"`
	Foreword    string    `json:"foreword"`
	Content     *string   `json:"content"`
	TagIDs      []int     `json:"tagIds,array"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"publishedAt"`
	StatusID    int       `json:"statusId"`
}

type Category struct {
	ID          int    `json:"categoryId"`
	Title       string `json:"title"`
	OrderNumber *int   `json:"orderNumber"`
	Alias       string `json:"alias"`
	StatusID    int    `json:"statusId"`
}

type Tag struct {
	ID       int    `json:"tagId"`
	Title    string `json:"title"`
	StatusID int    `json:"statusId"`
}
