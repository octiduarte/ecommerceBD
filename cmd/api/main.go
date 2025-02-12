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
	// Obtener el puerto desde Render
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Valor por defecto en local
	}

	// Obtener las variables de entorno de la base de datos
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Verificar que todas las variables de entorno estén configuradas
	if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || dbName == "" {
		fmt.Println("ERROR: Faltan variables de entorno para la base de datos")
		os.Exit(1)
	}

	// Conectar a la base de datos
	database, err := db.GetDB()
	if err != nil {
		fmt.Println("Failed to connect to database with error:", err)
		panic(err)
	}
	defer database.Close()

	fmt.Println("Connected to database")

	// Inicializar dependencias
	dm := dependencies.NewDependencyManager(database)

	// Crear manejadores
	productsHandlers := dm.ProductsHandler()
	storesHandler := dm.StoresHandler()
	mainStoresHandler := dm.MainStoresHandler()
	productsImageHandler := dm.ProductsImageHandler()

	// Configurar rutas
	r := routes.InitRoutes(productsHandlers, storesHandler, mainStoresHandler, productsImageHandler)

	// Configurar CORS para permitir solicitudes desde cualquier origen
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Permitir todas las solicitudes
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	// Servir imágenes desde la carpeta interna
	fs := http.FileServer(http.Dir("./internal/image"))
	r.PathPrefix("/image/").Handler(http.StripPrefix("/image/", fs))

	// Iniciar servidor en el puerto configurado
	fmt.Println("Starting server on port:", port)
	err = http.ListenAndServe(":"+port, corsHandler)
	if err != nil {
		fmt.Println("Error starting server with error:", err)
		panic(err)
	}
}
