package main_stores_service

import (
	"simi/internal/domain/interfaces"
	"simi/internal/domain/model"
)

type MainStoresService struct {
	mainStoresRepo interfaces.MainStoresRepository
}

func NewMainStoresService(mainStoresRepo interfaces.MainStoresRepository) MainStoresService {
	return MainStoresService{mainStoresRepo: mainStoresRepo}
}

func (s MainStoresService) GetMainStoreByID(storeID int) (mainStoreResponse model.MainResponse, err error) {
	return s.mainStoresRepo.GetMainStore(storeID)
}
