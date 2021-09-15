package usecases

import (
	"exchange_currency/domain/ocurrency"
	"exchange_currency/dto"
	"exchange_currency/infra/repository"
)

type OcurrencyService struct {
	repository ocurrency.IOccurrency
}

func NewOcurrencyService(repository *repository.OcurrencyRepository) *OcurrencyService {
	return &OcurrencyService{
		repository: repository,
	}
}

func (o *OcurrencyService) Exchange(d *dto.OcurrencyDTO) (interface{}, error) {

	dmn, err := d.Convert2Entity()

	if err != nil {
		return nil, err
	}

	err = o.repository.Store(dmn)

	if err != nil {
		return nil, err
	}

	return dmn.ToJSON(), nil
}

func (o *OcurrencyService) FindAll() ([]interface{}, error) {

	dmn, err := o.repository.FindAll()

	if err != nil {
		return nil, err
	}

	var res []interface{}

	for _, v := range dmn {
		res = append(res, v.DomainToJSON())
	}

	return res, nil
}

func (o *OcurrencyService) FindByID(id uint) (interface{}, error) {

	res, err := o.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	return res.DomainToJSON(), nil
}
