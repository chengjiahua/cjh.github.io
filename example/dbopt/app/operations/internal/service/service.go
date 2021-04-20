package service

import (
	"context"
	"dbopt/app/operations/conf"
	"dbopt/app/operations/internal/dao"
	"dbopt/app/operations/internal/service/ansible"
)

// Service biz service def.
type Service struct {
	c   *conf.Config
	dao *dao.Dao
	cmd *ansible.Command
}

// New new a Service and return.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
		cmd: ansible.NewCmd(),
	}
	return
}

// Ping check dao health.
func (s *Service) Ping(c context.Context) (err error) {
	return s.dao.Ping(c)
}

// Close close all dao.
func (s *Service) Close() {
	s.dao.Close()

}

// GetDao return dao instance.
func (s *Service) GetDao() *dao.Dao {
	return s.dao
}
