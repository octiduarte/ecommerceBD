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

type StoresHandler struct {
	storesService interfaces.StoresService
}

func NewStoresHandler(storesService interfaces.StoresService) StoresHandler {
	return StoresHandler{storesService: storesService}
}

func (h StoresHandler) GetStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stores, err := h.storesService.GetStores()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(stores)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (h StoresHandler) GetStoreByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	storeID, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}

	storeIDInt, err := strconv.Atoi(storeID)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	store, err := h.storesService.GetStoreByID(storeIDInt)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(store)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (h StoresHandler) SetStores(w http.ResponseWriter, r *http.Request) {
	store := &entities.Store{}
	err := json.NewDecoder(r.Body).Decode(store)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	storeID, err := h.storesService.SetStores(*store)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.NewEncoder(w).Encode(storeID)
	if err != nil {
		fmt.Println(err)
		return
	}
}
