package sqlite3

import (
	"time"

	"github.com/go-kratos/kratos/pkg/log"
	xtime "github.com/go-kratos/kratos/pkg/time"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

// Config mysql config.
type Config struct {
	DSN         string         // data source name.
	Active      int            // pool
	Idle        int            // pool
	IdleTimeout xtime.Duration 
}

// NewSqlite .
func NewSqlite(c *Config) (sqlite *xorm.Engine) {
	sqlite, err := xorm.NewEngine("sqlite3", c.DSN)
	if err != nil {
		log.Error("sqlite dsn(%s) error: %v", c.DSN, err)
		panic(err)
	}
	sqlite.DB().SetMaxIdleConns(c.Idle)
	sqlite.DB().SetMaxOpenConns(c.Active)
	sqlite.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)
	return
}
