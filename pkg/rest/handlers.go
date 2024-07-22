package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"newsportal/pkg/newsportal"
)

type FilterParams struct {
	CategoryID *int `query:"categoryID"`
	TagID      *int `query:"tagID"`
	Page       *int `query:"page"`
	PageSize   *int `query:"pageSize"`
}

type NewsService struct {
	m *newsportal.Manager
}

func NewNewsService(m *newsportal.Manager) *NewsService {
	return &NewsService{m: m}
}

// NewsByID получение новости по id
func (ss *NewsService) NewsByID(c echo.Context) error {
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
func (ss *NewsService) NewsWithFilters(c echo.Context) error {
	var params FilterParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid query parameters")
	}
	newsResponse, err := ss.m.News(params.CategoryID, params.TagID, params.Page, params.PageSize)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var newNewsList []NewsSummary
	for _, summary := range newsResponse {
		newNews := NewsSummaryFromManager(&summary)
		newNewsList = append(newNewsList, *newNews)
	}

	return c.JSON(http.StatusOK, newNewsList)
}

// NewsCountWithFilters получение количества новостей с фильтрами
func (ss *NewsService) NewsCountWithFilters(c echo.Context) error {
	var params FilterParams
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid query parameters")
	}
	count, err := ss.m.NewsCount(params.CategoryID, params.TagID, params.Page, params.PageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, count)
}

// Categories получение всех категорий
func (ss *NewsService) Categories(c echo.Context) error {
	categories, err := ss.m.Categories()
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	newCategories := CategoriesFromManager(categories)
	return c.JSON(http.StatusOK, newCategories)
}

// Tags получение всех тегов
func (ss *NewsService) Tags(c echo.Context) error {
	tags, err := ss.m.Tags()
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	newTags := TagsFromManager(tags)
	return c.JSON(http.StatusOK, newTags)
}
