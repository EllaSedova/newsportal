package newsportal

import (
	"newsportal/pkg/db"
)

func NewsSummaryFromDb(in *db.News) (out *NewsSummary) {
	if in != nil {
		out = &NewsSummary{
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

func NewsFromDb(in []db.News) (out []NewsSummary) {
	for _, news := range in {
		newNews := NewsSummaryFromDb(&news)
		out = append(out, *newNews)
	}
	return
}

func CategoryFromDb(in *db.Category) (out *Category) {
	if in != nil {
		out = &Category{
			ID:          in.ID,
			Title:       in.Title,
			OrderNumber: in.OrderNumber,
			Alias:       in.Alias,
			StatusID:    in.StatusID,
		}
	}
	return
}

func CategoriesFromDb(in []db.Category) (out []Category) {
	for _, category := range in {
		newCategories := CategoryFromDb(&category)
		out = append(out, *newCategories)
	}
	return
}

func TagFromDb(in *db.Tag) (out *Tag) {
	if in != nil {
		out = &Tag{
			ID:       in.ID,
			Title:    in.Title,
			StatusID: in.StatusID,
		}
	}
	return
}

func TagsFromDb(in []db.Tag) (out []Tag) {
	for _, tag := range in {
		newTags := TagFromDb(&tag)
		out = append(out, *newTags)
	}
	return
}
