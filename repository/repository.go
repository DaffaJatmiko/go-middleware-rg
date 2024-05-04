package repository

import (
	"errors"
	"middleware/model"
)

type IRepository interface {
	Get() []model.Data
	GetById(id int) model.Data
	Update(model.Data) error
	Delete(id int) error
	Create(model.Data) error
}

type repository struct {
	data map[int]model.Data
}


func NewRepository() IRepository {
	return &repository{
		data: make(map[int]model.Data),
	}
}

func (r *repository) Get() []model.Data {
	var datas []model.Data

	for _, data := range r.data {
		datas = append(datas, data)
	}

	return datas
}

func (r *repository) GetById(id int) model.Data {
	return r.data[id]
}

func (r *repository) Update(data model.Data) error {
	dataById := r.data[data.ID]

	if dataById.ID == 0 {
		return errors.New("data is empty")
	}

	dataById.Name = data.Name
	dataById.Address = data.Address
	dataById.Email = data.Email
	dataById.PhoneNumber = data.PhoneNumber

	r.data[data.ID] = dataById
	return nil
}

// Create implements IRepository.
func (r *repository) Create(data model.Data) error {
	if r.data[data.ID].ID != 0 {
		return errors.New("data already exists")
	}
	r.data[data.ID] = data
	return nil
}

// Delete implements IRepository.
func (*repository) Delete(id int) error {
	panic("unimplemented")
}
