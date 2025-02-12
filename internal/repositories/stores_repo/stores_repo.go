package stores_repo

import (
	"database/sql"
	"simi/internal/domain/model"
	"simi/internal/domain/model/entities"
)

type StoresRepository struct {
	db *sql.DB
}

func NewStoresRepository(db *sql.DB) StoresRepository {
	return StoresRepository{db: db}
}

func (r StoresRepository) SetStores(store entities.Store) (storeID int64, err error) {
	result, err := r.db.Exec("INSERT INTO Store (store_id, name, logo, banner, address) VALUES (?, ?, ?, ?, ?)",
		store.StoreID, store.Name, store.Logo, store.Banner, store.Address)
	if err != nil {
		return storeID, err
	}
	storeID, err = result.LastInsertId()
	if err != nil {
		return storeID, err
	}
	return storeID, nil
}

func (r StoresRepository) GetStores() (storesResponse []entities.Store, err error) {
	rows, err := r.db.Query("SELECT store_id, name, logo, banner, address FROM Store")
	if err != nil {
		return storesResponse, err
	}

	for rows.Next() {
		var store entities.Store
		err = rows.Scan(&store.StoreID, &store.Name, &store.Logo, &store.Banner, &store.Address)
		if err != nil {
			return storesResponse, err
		}
		storesResponse = append(storesResponse, store)
	}

	return storesResponse, nil
}

func (r StoresRepository) GetStoreById(id int) (store entities.Store, err error) {
	row := r.db.QueryRow("SELECT store_id, name, logo, banner, address FROM Store WHERE store_id = ?", id)
	err = row.Scan(&store.StoreID, &store.Name, &store.Logo, &store.Banner, &store.Address)
	if err != nil {
		return store, err
	}
	return store, nil
}

func (r StoresRepository) GetMainStoreById(id int) (store model.MainStore, err error) {
	row := r.db.QueryRow("SELECT name, logo, banner FROM Store WHERE store_id = ?", id)
	err = row.Scan(&store.Name, &store.Logo, &store.Banner)
	if err != nil {
		return store, err
	}

	rows, err := r.db.Query("SELECT name, url FROM Social_media WHERE store_id = ?", id)
	if err != nil {
		return store, err
	}
	for rows.Next() {
		var socialMediaBD entities.SocialMedia
		err = rows.Scan(&socialMediaBD.Name, &socialMediaBD.URL)
		store.SocialMedia = append(store.SocialMedia, socialMediaBD)
	}

	return store, nil
}

func (r StoresRepository) SetStoreImage(storeID int, imageLogo, imageBanner string) error {
	_, err := r.db.Exec("UPDATE Store SET logo = ?, banner = ? WHERE store_id = ?", imageLogo, imageBanner, storeID)
	if err != nil {
		return err
	}

	return nil
}
