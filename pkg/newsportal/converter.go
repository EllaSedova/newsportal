package newsportal

import (
	"newsportal/pkg/db"
)

func NewsFromDb(in *db.News) (out *News) {
	if in != nil {
		out = &News{
			ID:          in.ID,
			Title:       in.Title,
			CategoryID:  in.CategoryID,
			Foreword:    in.Foreword,
			Content:     in.Content,
			TagIDs:      in.TagIDs,
			Author:      in.Author,
			PublishedAt: in.PublishedAt,
			StatusID:    in.StatusID,
		}
	}
	return
}

func SomeNewsFromDb(in []db.News) (out []News) {
	for _, news := range in {
		newNews := NewsFromDb(&news)
		out = append(out, *newNews)
	}
	return
}
