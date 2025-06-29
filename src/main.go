package main

import (
	"context"
	"github.com/NewChakrit/golang_gin_vuejs_full-stack/controller"
	"github.com/NewChakrit/golang_gin_vuejs_full-stack/repository"
	"github.com/NewChakrit/golang_gin_vuejs_full-stack/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server ...")

	db, err := initDB()
	if err != nil {
		log.Fatalf("Unable to intitialize database: %v\n", err)
	}

	router := gin.Default()

	transactionRepository := repository.NewTransactionRepository(db.DB)
	transactionService := services.NewTransactionService(transactionRepository)

	controller.NewController(&controller.Config{
		router,
		transactionService,
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//Graceful service shutdown

	//Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default sent syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	if err = db.close(); err != nil {
		log.Fatalf("A problem occurred gracefully shutting dowm the database connection: %v\n", err)
	}

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
