package service

import "github.com/coding-standard/golang-project-layout/internal/dao"

type Service struct {
	dao dao.Interface
}

func (s *Service) DemoService() DemoService {
	return *NewDemoService()
}

func (s *Service) DemoDbService() DemoDbService {
	return *NewDemoDbService(s.dao)
}
