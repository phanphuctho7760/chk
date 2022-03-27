package services

import (
	"chk/models"
	"chk/repository"
)

type PriceHistoryService struct {
}

func (receiver PriceHistoryService) Create(params interface{}) models.Result {
	priceHistoryRepository := new(repository.PriceHistoryRepository)
	return priceHistoryRepository.Create(params)
}

func (receiver PriceHistoryService) Get(params models.GetParam) models.ResultGet {
	priceHistoryRepository := new(repository.PriceHistoryRepository)
	return priceHistoryRepository.Get(params)
}
