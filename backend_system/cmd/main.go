package main

import (
	"context"
	"github.com/wqh/smart/school/system/internal/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wqh/smart/school/system/internal/initiate"
)

/**
* description: 程序主入口
* author: wqh
* date: 2025/1/8
 */

func main() {
	engine := gin.Default()

	engine.LoadHTMLGlob("web/*")
	initiate.InitConfig()

	engine.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	engine.Use(middleware.Cors())
	engine.Use(middleware.ErrorHandler())

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
