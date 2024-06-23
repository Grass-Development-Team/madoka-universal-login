package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"projects.gdt.im/yuanzui-cf/madoka-universal-login/internal/config"
	"projects.gdt.im/yuanzui-cf/madoka-universal-login/internal/utils"
	"projects.gdt.im/yuanzui-cf/madoka-universal-login/routers"
)

func main() {
	// Set Log Writer
	utils.Log().Writer = io.MultiWriter(os.Stdout)

	// Init Config
	if !config.Conf().DebugMode {
		utils.Level = utils.LevelInformation
	}

	// Init Router
	r := routers.InitRouter()
	srv := &http.Server{
		Addr:    ":9000",
		Handler: r,
	}

	// Start Server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Log().Panic("Failed to start server: %s", err)
		}
	}()

	// Process Server Interrupt
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	<-quit
	shutdown(srv)
}

func shutdown(srv *http.Server) {
	utils.Log().Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		utils.Log().Panic("Failed to shutdown server: %s", err)
	}
	utils.Log().Info("Server shutdown")
}
