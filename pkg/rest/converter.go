package rest

import (
	"newsportal/pkg/newsportal"
)

func NewsFromManager(in *newsportal.News) (out *News) {
	if in != nil {
		out = &News{
			ID:          in.ID,
			Title:       in.Title,
			Foreword:    in.Foreword,
			Content:     in.Content,
			Author:      in.Author,
			PublishedAt: in.PublishedAt,
			Tags:        TagsFromManager(in.Tags),
			Category:    *CategoryFromManager(in.Category),
		}

	}
	return
}

func NewsSummaryFromManager(in *newsportal.News) (out *NewsSummary) {
	if in != nil {
		out = &NewsSummary{
			ID:          in.ID,
			Title:       in.Title,
			Foreword:    in.Foreword,
			Author:      in.Author,
			PublishedAt: in.PublishedAt,
			Tags:        TagsFromManager(in.Tags),
			Category:    *CategoryFromManager(in.Category),
		}
	}
	return
}

func CategoryFromManager(in *newsportal.Category) (out *Category) {
	if in != nil {
		out = &Category{
			ID:          in.ID,
			Title:       in.Title,
			OrderNumber: in.OrderNumber,
			Alias:       in.Alias,
		}
	}
	return
}

func CategoriesFromManager(in []newsportal.Category) (out []Category) {
	for _, category := range in {
		newCategories := CategoryFromManager(&category)
		out = append(out, *newCategories)
	}
	return
}

func TagFromManager(in *newsportal.Tag) (out *Tag) {
	if in != nil {
		out = &Tag{
			ID:    in.ID,
			Title: in.Title,
		}
	}
	return
}

func TagsFromManager(in []newsportal.Tag) (out []Tag) {
	for _, tag := range in {
		newTags := TagFromManager(&tag)
		out = append(out, *newTags)
	}
	return
}
