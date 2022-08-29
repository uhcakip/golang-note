package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.String(http.StatusOK, "welcome gin server")
	})

	service := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 丟到背景執行避免阻斷後續流程
	go func() {
		if err := service.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("listen error:", err)
		}
	}()

	// 等待關閉訊號
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("caught signal to quit")

	// 等待 5 秒後關閉所有連線
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := service.Shutdown(ctx); err != nil {
		log.Fatalln("service shutdown error:", err)
	}

	log.Println("service exit")
}
