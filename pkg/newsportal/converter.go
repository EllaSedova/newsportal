package newsportal

import (
	"newsportal/pkg/db"
)

func NewsFromDB(in *db.News) *News {
	if in == nil {
		return nil
	}
	return &News{
		News:     in,
		Category: CategoryFromDb(in.Category),
	}
}

func NewNewsListFromDB(in []db.News) (out []News) {
	for i := range in {
		out = append(out, *NewsFromDB(&in[i]))
	}

	return
}

func CategoryFromDb(in *db.Category) *Category {
	if in == nil {
		return nil
	}
	return &Category{
		ID:          in.ID,
		Title:       in.Title,
		OrderNumber: in.OrderNumber,
		Alias:       in.Alias,
		StatusID:    in.StatusID,
	}
}

func CategoriesFromDb(in []db.Category) (out []Category) {
	for _, category := range in {
		newCategories := CategoryFromDb(&category)
		out = append(out, *newCategories)
	}
	return
}

func TagFromDb(in *db.Tag) *Tag {
	if in == nil {
		return nil
	}

	return &Tag{
		ID:       in.ID,
		Title:    in.Title,
		StatusID: in.StatusID,
	}
}

func TagsFromDb(in []db.Tag) (out []Tag) {
	for _, tag := range in {
		newTags := TagFromDb(&tag)
		out = append(out, *newTags)
	}
	return
}
