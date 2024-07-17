package newsportal

import "time"

type News struct {
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
