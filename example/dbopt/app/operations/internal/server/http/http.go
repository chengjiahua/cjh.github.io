package http

import (
	"net/http"

	v1 "dbopt/app/operations/api/v1"
	"dbopt/app/operations/conf"
	"dbopt/app/operations/internal/service"

	"github.com/go-kratos/kratos/pkg/log"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

var (
	operationsSvc *service.Service
	scriptSvc     *conf.Script
)

// New ...
func New(c *conf.Config, s *service.Service) (b *bm.Engine) {
	operationsSvc = s
	initService(c)
	// init inner router
	engineIn := bm.DefaultServer(c.Web)
	innerRouter(engineIn)
	// init inner server
	if err := engineIn.Start(); err != nil {
		log.Error("engineInner.Start error (%v)", err)
		panic(err)
	}

	return engineIn
}

func initService(c *conf.Config) {
}

// innerRouter .
func innerRouter(e *bm.Engine) {
	e.Ping(ping)
	v1.RegisterKoOperationsBMServer(e, service.New(conf.Conf))
}

// ping check server ok.
func ping(c *bm.Context) {
	if err := operationsSvc.Ping(c); err != nil {
		log.Error("service ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
