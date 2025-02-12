package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"simi/internal/domain/interfaces"
	"strconv"
)

type MainStoresHandler struct {
	mainStoresService interfaces.MainStoresService
}

func NewMainStoresHandler(mainStoresService interfaces.MainStoresService) MainStoresHandler {
	return MainStoresHandler{mainStoresService: mainStoresService}
}

func (h MainStoresHandler) GetMainStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	storeID, ok := vars["store_id"]
	if !ok {
		fmt.Println("store_id is missing in parameters")
	}

	storeIDInt, err := strconv.Atoi(storeID)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	mainStore, err := h.mainStoresService.GetMainStoreByID(storeIDInt)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode(mainStore)
	if err != nil {
		fmt.Println(err)
		return
	}
}
