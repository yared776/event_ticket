package initiator

import (
	"context"
	"event_ticket/routing"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Initiate() {
	server := gin.Default()
	v1 := server.Group("v1")
	server.LoadHTMLGlob("public/*.html")
	server.Static("/public", "./public")
	routing.RegisterRoutes(v1, routing.Routes)
	srv := &http.Server{
		Addr:        "localhost:9000",
		ReadTimeout: 1 * time.Second,
		Handler:     server,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	go func() {
		srv.ListenAndServe()
	}()

	_ = <-quit
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	log.Printf("sever is going to shut down %+V", srv.Shutdown(ctx))

}
