package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"newsportal/pkg/newsportal"
)

type ServerService struct {
	m *newsportal.Manager
}

func NewServerService(m *newsportal.Manager) *ServerService {
	return &ServerService{m: m}
}

func Respond(w http.ResponseWriter, data any) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// NewsByID получение новости по id
func (ss *ServerService) NewsByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	fmt.Println(id)
	news, err := ss.m.NewsByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	Respond(w, news)
}

// NewsWithFilters получение новости с фильтрами
func (ss *ServerService) NewsWithFilters(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	var id *int
	if newsId := queryParams.Get("id"); newsId != "" {
		idInt, err := strconv.Atoi(newsId)
		if err != nil {
			http.Error(w, "Invalid newsId parameter", http.StatusBadRequest)
			log.Fatal(err)
			return
		}
		id = &idInt
	}
	var categoryID *int
	if categoryId := queryParams.Get("categoryID"); categoryId != "" {
		categoryInt, err := strconv.Atoi(categoryId)
		if err != nil {
			http.Error(w, "Invalid categoryId parameter", http.StatusBadRequest)
			log.Fatal(err)
			return
		}
		categoryID = &categoryInt
	}
	var tagID *int
	if tagId := queryParams.Get("tagID"); tagId != "" {
		tagInt, err := strconv.Atoi(tagId)
		if err != nil {
			http.Error(w, "Invalid tagId parameter", http.StatusBadRequest)
			log.Fatal(err)
			return
		}
		tagID = &tagInt
	}
	var page *int
	if pageNumber := queryParams.Get("page"); pageNumber != "" {
		pageInt, err := strconv.Atoi(pageNumber)
		if err != nil {
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			log.Fatal(err)
			return
		}
		page = &pageInt
	}
	var pageSize *int
	if pageSizeNumber := queryParams.Get("pageSize"); pageSizeNumber != "" {
		pageSizeInt, err := strconv.Atoi(pageSizeNumber)
		if err != nil {
			http.Error(w, "Invalid pageSize parameter", http.StatusBadRequest)
			log.Fatal(err)
			return
		}
		pageSize = &pageSizeInt
	}
	var sortTitle *bool
	if sort := queryParams.Get("sortTitle"); sort != "" {
		sortBool := sort == "true"
		sortTitle = &sortBool
	}
	news, err := ss.m.News(id, categoryID, tagID, page, pageSize, sortTitle)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	Respond(w, news)

}

// Categories получение всех категорий
func (ss *ServerService) Categories(w http.ResponseWriter, r *http.Request) {
	categories, err := ss.m.Categories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	Respond(w, categories)
}

// Tags получение всех тегов
func (ss *ServerService) Tags(w http.ResponseWriter, r *http.Request) {
	tags, err := ss.m.Tags()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	Respond(w, tags)
}
