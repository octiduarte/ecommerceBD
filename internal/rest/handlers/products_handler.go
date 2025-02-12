package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"simi/internal/domain/interfaces"
	"simi/internal/domain/model/entities"
	"strconv"
)

type ProductsHandler struct {
	productsService interfaces.ProductsService
}

func NewProductsHandler(productsService interfaces.ProductsService) ProductsHandler {
	return ProductsHandler{productsService: productsService}
}

func (h ProductsHandler) GetProductsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	productID, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	product, err := h.productsService.GetProductByID(productIDInt)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (h ProductsHandler) SetProducts(w http.ResponseWriter, r *http.Request) {
	product := &entities.Product{}
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	products, err := h.productsService.SetProducts(*product)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		fmt.Println(err)
		return
	}
}
