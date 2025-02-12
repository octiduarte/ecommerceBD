package main

import (
	"fmt"
	"net/http"
	"os"
	"simi/cmd/api/dependencies"
	"simi/cmd/api/routes"
	"simi/internal/utils/db"

	"github.com/gorilla/handlers"
)

func main() {
	database, err := db.GetDB()
	if err != nil {
		fmt.Println("Failed to connect to database with error: ", err)
		panic(err)
	}
	defer database.Close()

	fmt.Println("Connected to database")

	dm := dependencies.NewDependencyManager(database)

	productsHandlers := dm.ProductsHandler()
	storesHandler := dm.StoresHandler()
	mainStoresHandler := dm.MainStoresHandler()
	productsImageHandler := dm.ProductsImageHandler()

	r := routes.InitRoutes(productsHandlers, storesHandler, mainStoresHandler, productsImageHandler)

	// Configurar CORS para permitir solicitudes desde localhost:3000
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	fs := http.FileServer(http.Dir("./internal/image"))
	r.PathPrefix("/image/").Handler(http.StripPrefix("/image/", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Usa 8000 por defecto si no está definido
	}

	err = http.ListenAndServe(":"+port, corsHandler)

	if err != nil {
		fmt.Println("Error starting server with error: ", err)
		panic(err)
	}

}
