package router

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var server *http.Server

func Run(addr string, router *gin.Engine) {

	server = &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    time.Second * 60 * 3,
		WriteTimeout:   time.Second * 60 * 3,
		MaxHeaderBytes: 256 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Info("web server shutdown complete")
			} else {
				log.Errorf("web server closed unexpect: %s", err)
			}
		}
	}()
}

func Stop() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Errorf("web server shutdown error:%s", err.Error())
	}
	log.Infoln("web server exiting")
}
