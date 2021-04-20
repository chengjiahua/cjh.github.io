package model

import "time"

type User struct {
	ID        int64     `xorm:"id"` // xorm默认自动递增
	Name      string    `xorm:"name"`
	Age       string    `xorm:"age"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
