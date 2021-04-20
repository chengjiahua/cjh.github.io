package dao

import (
	"context"
	"dbopt/app/operations/conf"
	"dbopt/app/operations/internal/model"
	"dbopt/library/database/sqlite3"

	"github.com/go-kratos/kratos/pkg/log"
	"xorm.io/xorm"
)

// Dao .
type Dao struct {
	Sqlite *xorm.Engine
}

// New new a instance.
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		Sqlite: sqlite3.NewSqlite(c.SQLITE),
	}
	d.initORM()

	return
}

func (d *Dao) initORM() {
	if err := d.Sqlite.Sync2(&model.User{}); err != nil {
		log.Error("d.Sync2 error (%v)", err)
		panic(err)
	}
	d.Sqlite.ShowSQL(true)
}

// Ping check connection of db , mc .
func (d *Dao) Ping(c context.Context) (err error) {
	if d.Sqlite != nil {
		if err = d.Sqlite.DB().PingContext(c); err != nil {
			log.Error("d.PingContext error (%v)", err)
			return
		}
	}

	return
}

// Close close connection of db.
func (d *Dao) Close() {
	if d.Sqlite != nil {
		d.Sqlite.Close()
	}
}
