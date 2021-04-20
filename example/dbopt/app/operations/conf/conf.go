package conf

import (
	"flag"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"dbopt/library/database/orm"
	"dbopt/library/database/sqlite3"
)

// global var
var (
	Conf = &Config{}
)

//Script represent web script
type Script struct {
	Dir           string   `yaml:"dir"`
	OriginScripts []string `yaml:"origin_scripts"`
}

// Config config set
type Config struct {
	Version string           `yaml:"version"`
	Web     *bm.ServerConfig `yaml:"web"`
	Log     *log.Config      `yaml:"log"`
	ORM     *orm.Config      `yaml:"orm"`
	SQLITE  *sqlite3.Config  `yaml:"sqlite3"`
}

// Init init conf
func Init() error {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}

	if err := paladin.Get("app.yml").UnmarshalYAML(&Conf); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	return nil
}
