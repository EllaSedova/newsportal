package newsportal

import (
	"newsportal/pkg/db"
)

func NewsFromDb(in *db.News, tags []Tag) (out *News) {
	if in != nil {
		out = &News{
			News: *in,
		}
		out.Category = CategoryFromDb(in.Category)
		out.Tags = tags
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
