package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"newsportal/pkg/newsportal"
)

type FilterParams struct {
	CategoryID *int  `query:"categoryID"`
	TagID      *int  `query:"tagID"`
	Page       *int  `query:"page"`
	PageSize   *int  `query:"pageSize"`
	SortTitle  *bool `query:"sortTitle"`
}

type ServerService struct {
	m *newsportal.Manager
}

func NewServerService(m *newsportal.Manager) *ServerService {
	return &ServerService{m: m}
}

// NewsByID получение новости по id
func (ss *ServerService) NewsByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid ID")
	}
	news, err := ss.m.NewsByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	newNews := NewsFromManager(news)
	return c.JSON(http.StatusOK, newNews)
}

// NewsWithFilters получение новости с фильтрами
func (ss *ServerService) NewsWithFilters(c echo.Context) error {
	var params FilterParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid query parameters")
	}
	news, _, err := ss.m.News(params.CategoryID, params.TagID, params.Page, params.PageSize, params.SortTitle, false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var newNewsList []NewsSummary
	for _, summary := range news {
		newNews := NewsSummaryFromManager(&summary)
		newNewsList = append(newNewsList, *newNews)
	}

	return c.JSON(http.StatusOK, newNewsList)
}

// NewsCountWithFilters получение количества новостей с фильтрами
func (ss *ServerService) NewsCountWithFilters(c echo.Context) error {
	var params FilterParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid query parameters")
	}
	_, count, err := ss.m.News(params.CategoryID, params.TagID, params.Page, params.PageSize, params.SortTitle, true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, count)
}

// Categories получение всех категорий
func (ss *ServerService) Categories(c echo.Context) error {
	categories, err := ss.m.Categories()
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	newCategories := CategoriesFromManager(categories)
	return c.JSON(http.StatusOK, newCategories)
}

// Tags получение всех тегов
func (ss *ServerService) Tags(c echo.Context) error {
	tags, err := ss.m.Tags()
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	newTags := TagsFromManager(tags)
	return c.JSON(http.StatusOK, newTags)
}
