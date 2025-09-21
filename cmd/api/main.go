package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"belajar-go-http/internal/handler"
	"belajar-go-http/internal/model"
)

func main() {
	// 1. Inisialisasi dependensi
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	store := model.NewTodoStore()
	todoHandler := &handler.TodoHandler{
		Store: store,
	}

	// 2. Setup router
	router := todoHandler.Routes()

	// 3. Buat dan konfigurasikan server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// 4. Implementasi Graceful Shutdown
	go func() {
		logger.Println("Server starting on port", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %s\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	logger.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server Shutdown Failed:%+v", err)
	}

	logger.Println("Server gracefully stopped")
}
