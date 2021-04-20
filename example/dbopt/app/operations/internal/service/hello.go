package service

import (
	"context"
	v1 "dbopt/app/operations/api/v1"

	"github.com/go-kratos/kratos/pkg/ecode"
	"github.com/go-kratos/kratos/pkg/log"
)

// SayHelloURL .
func (s *Service) SayHelloURL(ctx context.Context, req *v1.HelloReq) (rsp *v1.HelloResp, err error) {
	if req.Age == "" {
		return nil, ecode.Error(v1.NameIsEmptyErr, v1.NameIsEmptyErr.Message())
	}
	err = s.dao.UserInsert(req.Name, req.Age)
	if err != nil {
		log.Error("s.dao.UserInsert err : %v", err)
	}
	
	rsp = &v1.HelloResp{
		Content: "name : " + req.Name + " age : " + req.Age,
	}

	return
}
