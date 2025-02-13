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
	// Obtener la conexión a la base de datos
	database, err := db.GetDB()
	if err != nil {
		fmt.Println("Failed to connect to database with error: ", err)
		panic(err)
	}
	defer database.Close()

	fmt.Println("Connected to database")

	// Inicializar manejadores de dependencias
	dm := dependencies.NewDependencyManager(database)

	productsHandlers := dm.ProductsHandler()
	storesHandler := dm.StoresHandler()
	mainStoresHandler := dm.MainStoresHandler()
	productsImageHandler := dm.ProductsImageHandler()

	// Inicializar rutas
	r := routes.InitRoutes(productsHandlers, storesHandler, mainStoresHandler, productsImageHandler)

	// Configurar CORS para permitir solicitudes desde cualquier origen (ajústalo si es necesario)
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Permitir cualquier origen (cambiar si es necesario)
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	// Servir imágenes desde la carpeta interna
	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/image/").Handler(http.StripPrefix("/image/", fs))

	// Obtener el puerto desde la variable de entorno (si no está, usar 8000 por defecto)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Server running on port:", port)

	// Iniciar servidor
	err = http.ListenAndServe(":"+port, corsHandler)
	if err != nil {
		fmt.Println("Error starting server with error: ", err)
		panic(err)
	}
}
