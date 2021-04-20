package orm

import (
	"time"

	xtime "github.com/go-kratos/kratos/pkg/time"

	"github.com/go-kratos/kratos/pkg/log"

	// database driver
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
	"xorm.io/xorm"
)

// Config mysql config.
type Config struct {
	DSN         string         // data source name.
	Active      int            // pool
	Idle        int            // pool
	IdleTimeout xtime.Duration // connect max life time.
}

//type ormLog struct{}
//
//func (l ormLog) Print(v ...interface{}) {
//	log.Info(strings.Repeat("%v ", len(v)), v...)
//}

func init() {
	//gorm.ErrRecordNotFound = ecode.NothingFound
}

// NewMySQL .
func NewMySQL(c *Config) (db *xorm.Engine) {
	db, err := xorm.NewEngine("mysql", c.DSN)
	if err != nil {
		log.Error("db dsn(%s) error: %v", c.DSN, err)
		panic(err)
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)
	//db.ShowSQL(true)
	return

}
