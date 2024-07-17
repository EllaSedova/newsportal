package newsportal

import (
	"fmt"
	"newsportal/pkg/db"
)

type Manager struct {
	nr db.NewsRepo
}

func NewManager(db db.NewsRepo) *Manager {
	return &Manager{nr: db}
}
func (m Manager) NewsByID(id int) (*News, error) {
	news, err := m.nr.NewsByID(id)
	if err != nil {
		return nil, fmt.Errorf("[NewsByID]: %s\n", err)
	}
	return NewsFromDb(news), nil
}

// NewsByTagID возвращает все новости с заданным тегом
func (m Manager) NewsByTagID(tagID int) ([]News, error) {
	// checking for the existence of a tag
	_, err := m.nr.TagByID(tagID)
	if err != nil {
		return nil, err
	}
	// get news by tagId
	news, err := m.nr.NewsByTagID(tagID)
	if err != nil {
		return nil, fmt.Errorf("[NewsByTagID]: %s\n", err)
	}
	return SomeNewsFromDb(news), nil
}

//// NewsByCategoryID возвращает все новости в заданной категории
//func (nr *NewsRepo) NewsByCategoryID(categoryID int) ([]News, error) {
//	var news []News
//	err := nr.Model(&news).Where("\"categoryId\" = ?", categoryID).Select()
//	return news, err
//}
//
//// NewsByTagIDWithLimit возвращает ограниченное количество новостей с заданным тегом
//func (nr *NewsRepo) NewsByTagIDWithLimit(tagID, limit int) ([]News, error) {
//	var news []News
//	err := nr.Model(&news).Where("? = ANY (\"tagIds\")", tagID).Limit(limit).Select()
//	return news, err
//}
//
//// NewsByCategoryIDWithLimit возвращает ограниченное количество новостей в заданной категории
//func (nr *NewsRepo) NewsByCategoryIDWithLimit(categoryID, limit int) ([]News, error) {
//	var news []News
//	err := nr.Model(&news).Where("\"categoryId\" = ?", categoryID).Limit(limit).Select()
//	return news, err
//}
//
//// NewsByTagIDWithPagination возвращает новости с заданным тегом с пагинацией
//func (nr *NewsRepo) NewsByTagIDWithPagination(tagID, page, pageSize int) ([]News, error) {
//	var news []News
//	offset := (page - 1) * pageSize
//	err := nr.Model(&news).
//		Where("? = ANY (\"tagIds\")", tagID).
//		Offset(offset).
//		Limit(pageSize).
//		Select()
//	return news, err
//}
//
//// NewsByCategoryIDWithPagination возвращает новости в заданной категории с пагинацией
//func (nr *NewsRepo) NewsByCategoryIDWithPagination(categoryID, page, pageSize int) ([]News, error) {
//	var news []News
//	offset := (page - 1) * pageSize
//	err := nr.Model(&news).
//		Where("\"categoryId\" = ?", categoryID).
//		Offset(offset).
//		Limit(pageSize).
//		Select()
//	return news, err
//}
//
//// CategoryByID возвращает Category по id из бд
//func (nr *NewsRepo) CategoryByID(id int) (*Category, error) {
//	category := &Category{ID: id}
//	err := nr.Model(category).WherePK().Select()
//	return category, err
//}
//
//// TagByID возвращает Tag по id из бд
//func (nr *NewsRepo) TagByID(id int) (*Tag, error) {
//	tag := &Tag{ID: id}
//	err := nr.Model(tag).WherePK().Select()
//	return tag, err
//}
