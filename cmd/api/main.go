package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fabricioviapiana/orders-app/internal/handler"
	"github.com/fabricioviapiana/orders-app/internal/repository"
	"github.com/fabricioviapiana/orders-app/internal/service"
)

const ServerPort = ":8080"

// aqui o ponteiro de request é passado para favorecer a performance,
// visto que é um struc gigante e é custo passar para todas request
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"status":  "ok",
		"message": "running",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "error encoding the response", http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	productRepository := repository.NewInMemoryProductRepository()
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	mux.HandleFunc("/products", productHandler.HandleProducts)

	log.Println("Server running on port", ServerPort)
	if err := http.ListenAndServe(ServerPort, mux); err != nil {
		log.Fatal(err)
	}
}
