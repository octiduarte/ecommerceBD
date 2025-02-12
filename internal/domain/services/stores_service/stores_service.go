package stores_service

import (
	"simi/internal/domain/interfaces"
	"simi/internal/domain/model/entities"
)

type StoresService struct {
	storesRepository interfaces.StoresRepository
}

func NewStoresService(storesRepository interfaces.StoresRepository) StoresService {
	return StoresService{storesRepository: storesRepository}
}

func (s StoresService) GetStores() (storesResponse []entities.Store, err error) {
	storesResponse, err = s.storesRepository.GetStores()
	if err != nil {
		return nil, err
	}
	return storesResponse, nil
}

func (s StoresService) GetStoreByID(id int) (storeResponse entities.Store, err error) {
	storeResponse, err = s.storesRepository.GetStoreById(id)
	if err != nil {
		return storeResponse, err
	}
	return storeResponse, nil
}

func (s StoresService) SetStores(newStore entities.Store) (storeID int64, err error) {
	storeID, err = s.storesRepository.SetStores(newStore)
	if err != nil {
		return storeID, err
	}
	return storeID, nil
}
