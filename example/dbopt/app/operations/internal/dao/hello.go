package dao

import (
	"dbopt/app/operations/internal/model"
	"time"

	"github.com/go-kratos/kratos/pkg/log"
)

// UserInsert .
func (d *Dao) UserInsert(name string, age string) (err error) {
	User := model.User{
		Name:      name,
		Age:       age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = d.Sqlite.Insert(&User)
	if err != nil {
		log.Error("UserInsert Insert err : %v", err)
		return
	}
	return
}
