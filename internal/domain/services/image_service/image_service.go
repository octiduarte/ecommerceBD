package image_service

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"simi/internal/domain/interfaces"
	"strconv"
)

type ImageService struct {
	productsImageRepo interfaces.ImageRepository
	productsRepo      interfaces.ProductsRepository
	storesRepo        interfaces.StoresRepository
}

func NewImageService(productsImageRepo interfaces.ImageRepository,
	productsRepo interfaces.ProductsRepository,
	storesRepo interfaces.StoresRepository) ImageService {
	return ImageService{productsImageRepo: productsImageRepo, productsRepo: productsRepo, storesRepo: storesRepo}
}

func (s ImageService) UploadProductsImage(file multipart.File, handler *multipart.FileHeader, productID string) error {
	storeName, err := s.productsRepo.GetStoreNameByProductID(productID)

	imagePath := filepath.Join("public", "image", storeName, handler.Filename)
	destFile, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		return err
	}

	return s.productsImageRepo.UploadImageToProduct(productID, imagePath)
}

func (s ImageService) UploadStoreImage(fileLogo multipart.File, handlerLogo *multipart.FileHeader, fileBanner multipart.File,
	handlerBanner *multipart.FileHeader, StoreID string) error {
	storeIDInt, _ := strconv.Atoi(StoreID)
	store, err := s.storesRepo.GetStoreById(storeIDInt)
	if err != nil {
		return err
	}
	imagePathLogo := filepath.Join("public", "image", store.Name, handlerLogo.Filename)
	destFileLogo, err := os.Create(imagePathLogo)
	if err != nil {
		return err
	}
	defer destFileLogo.Close()

	_, err = io.Copy(destFileLogo, fileLogo)
	if err != nil {
		return err
	}

	imagePathBanner := filepath.Join("public", "image", store.Name, handlerBanner.Filename)
	destFileBanner, err := os.Create(imagePathBanner)
	if err != nil {
		return err
	}
	defer destFileBanner.Close()

	_, err = io.Copy(destFileBanner, fileLogo)
	if err != nil {
		return err
	}

	return s.storesRepo.SetStoreImage(storeIDInt, imagePathLogo, imagePathBanner)
}
