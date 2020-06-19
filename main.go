package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"yangjian.com/basicarchitecture/api/controllers"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"yangjian.com/basicarchitecture/api/router"
)

// @title Swagger basic architecture API
// @version 1.0
// @description This is a visual data query service for meteorological data.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8090
// @BasePath /api/v1
func main() {
	// init controller
	if err := controllers.InitController(); err != nil {
		log.Errorf("init is error:%s", err.Error())
		return
	}
	// get port
	port := viper.GetInt("server.port")
	router.Run(fmt.Sprintf(":%d", port), router.Router())
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	router.Stop()
}
