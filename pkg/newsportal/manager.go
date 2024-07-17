package newsportal

import (
	"newsportal/pkg/db"
)

type Manager struct {
	nr db.NewsRepo
}

func NewManager(db db.NewsRepo) *Manager {
	return &Manager{nr: db}
}
func (m Manager) NewsByID(id int) {

}
