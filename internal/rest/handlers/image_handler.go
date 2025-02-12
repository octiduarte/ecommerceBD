package handlers

import (
	"fmt"
	"net/http"
	"simi/internal/domain/interfaces"
)

type ImageHandler struct {
	imageService interfaces.ImageService
}

func NewProductsImageHandler(imageService interfaces.ImageService) ImageHandler {
	return ImageHandler{imageService: imageService}
}

func (h ImageHandler) UploadProductImage(w http.ResponseWriter, r *http.Request) {
	// Parse the form to retrieve file
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	productID := r.FormValue("product_id")
	if productID == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	err = h.imageService.UploadProductsImage(file, handler, productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

func (h ImageHandler) UploadStoreImage(w http.ResponseWriter, r *http.Request) {
	// Parse the form to retrieve file
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	fileLogo, handlerLogo, err := r.FormFile("image_logo")
	if err != nil {
		http.Error(w, "Failed to retrieve file logo", http.StatusBadRequest)
		return
	}
	defer fileLogo.Close()

	fileBanner, handlerBanner, err := r.FormFile("image_banner")
	if err != nil {
		http.Error(w, "Failed to retrieve file banner", http.StatusBadRequest)
		return
	}
	defer fileBanner.Close()

	storeID := r.FormValue("store_id")
	if storeID == "" {
		http.Error(w, "Store ID is required", http.StatusBadRequest)
		return
	}

	err = h.imageService.UploadStoreImage(fileLogo, handlerLogo, fileBanner, handlerBanner, storeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", handlerLogo.Filename)
	fmt.Fprintf(w, "File uploaded successfully: %s", handlerBanner.Filename)

}
