package main

import (
	"log"
	"net/http"
	"tienda/handlers"
	"tienda/routes"
	"tienda/storage"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// 1. Inicializa la capa de almacenamiento
	store := storage.NewMemoryStore()

	// 2. Crea las instancias de los manejadores
	productHandlers := handlers.NewProductHandlers(store)
	userHandlers := handlers.NewUserHandlers(store)
	cartHandlers := handlers.NewCartHandlers(store, store, store)
	reportHandlers := handlers.NewReportHandlers(store, store)

	// 3. Crea el enrutador principal
	r := mux.NewRouter()

	// 4. Se elimina r.Use(CORSMiddleware). La configuración se hará de otra forma.

	// 5. Registra todas las rutas de la API (sin cambios).
	routes.RegisterRoutes(r, productHandlers, cartHandlers, userHandlers, reportHandlers)

	// 6. Configura CORS usando la librería 'rs/cors'.
	//    Esto es más seguro que usar "*", ya que solo permite tu frontend.
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8001"}, // El origen de tu app web
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		Debug:            true, // Muy útil para depurar problemas de CORS
	})

	// 7. Crea el manejador final envolviendo el enrutador con el middleware de CORS.
	handler := c.Handler(r)

	// 8. Inicia el servidor de la API con el manejador que incluye CORS.
	log.Println("🚀 Servidor API iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal("Error al iniciar el servidor API: ", err)
	}
}
