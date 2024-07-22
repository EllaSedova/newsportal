package rpc

import (
	"newsportal/pkg/newsportal"
)

func NewsFromManager(in *newsportal.News) *News {
	if in == nil {
		return nil
	}

	return &News{
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

func NewsSummaryFromManager(in *newsportal.News) *NewsSummary {
	if in == nil {
		return nil
	}

	return &NewsSummary{
		ID:          in.ID,
		Title:       in.Title,
		Foreword:    in.Foreword,
		Author:      in.Author,
		PublishedAt: in.PublishedAt,
		Tags:        TagsFromManager(in.Tags),
		Category:    *CategoryFromManager(in.Category),
	}

}

func CategoryFromManager(in *newsportal.Category) *Category {
	if in == nil {
		return nil
	}

	return &Category{
		ID:          in.ID,
		Title:       in.Title,
		OrderNumber: in.OrderNumber,
		Alias:       in.Alias,
	}
}

func CategoriesFromManager(in []newsportal.Category) (out []Category) {
	for _, category := range in {
		newCategories := CategoryFromManager(&category)
		out = append(out, *newCategories)
	}
	return
}

func TagFromManager(in *newsportal.Tag) *Tag {
	if in == nil {
		return nil
	}

	return &Tag{
		ID:    in.ID,
		Title: in.Title,
	}
}

func TagsFromManager(in []newsportal.Tag) (out []Tag) {
	for _, tag := range in {
		newTags := TagFromManager(&tag)
		out = append(out, *newTags)
	}
	return
}
