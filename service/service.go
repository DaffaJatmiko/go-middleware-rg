package service

import (
	"middleware/model"
	"middleware/repository"
)

type IService interface {
	Get() []model.Data
	GetById(id int) model.Data
	Update(model.Data) error
	Delete(id int) error
	Create(model.Data) error
}

type service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) IService {
	return &service{
		Repository: repository,
	}
}

func (s *service) Get() []model.Data {
	return s.Repository.Get()
}

func (s *service) GetById(id int) model.Data {
	return s.Repository.GetById(id)
}

func (s *service) Delete(id int) error {
	return s.Repository.Delete(id)
}

func (s *service) Update(data model.Data) error {
	return s.Repository.Update(data)
}

func (s *service) Create(data model.Data) error {
	return s.Repository.Create(data)
}