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
	dsn := "postgres://user:password@localhost:5432/orders_db?sslmode=disable"
	db, err := repository.NewDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := repository.RunMigrations(db); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	//Products
	productRepository := repository.NewInMemoryProductRepository()
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	mux.HandleFunc("GET /products", productHandler.HandleProducts)
	mux.HandleFunc("POST /products", productHandler.HandleProducts)
	mux.HandleFunc("GET /products/{id}", productHandler.HandleProducts)

	//User
	userRepository := repository.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	mux.HandleFunc("GET /users", userHandler.HandleUsers)
	mux.HandleFunc("POST /users", userHandler.HandleUsers)
	mux.HandleFunc("GET /users/{id}", userHandler.HandleUsers)

	//Orders
	orderRepository := repository.NewInMemoryOrderRepository()
	orderService := service.NewOrderService(orderRepository, productService, userService)
	orderHandler := handler.NewOrderHandler(orderService)
	mux.HandleFunc("/orders", orderHandler.HandleOrders)

	log.Println("Server running on port", ServerPort)
	if err := http.ListenAndServe(ServerPort, mux); err != nil {
		log.Fatal(err)
	}
}
