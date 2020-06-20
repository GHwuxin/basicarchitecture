package router

import (
	"io"
	"os"
	"yangjian.com/basicarchitecture/api/controllers"

	_ "yangjian.com/basicarchitecture/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router() (router *gin.Engine) {

	// set server start mode
	// debug or release
	mode := gin.DebugMode
	switch viper.GetString("server.mode") {
	case "debug":
		mode = gin.DebugMode
	case "release":
		mode = gin.ReleaseMode
	case "test":
		mode = gin.TestMode
	default:
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	// disable logger trans to console with color
	gin.DisableConsoleColor()
	// set writer
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	// get router
	router = gin.New()
	// set middleware, normal
	router.Use(gin.Logger(), gin.Recovery())
	// set cors
	router.Use(cors.Default())
	// register routes
	registerRoutes(router)

	return
}

func registerRoutes(router *gin.Engine) {

	// swagger init
	//url := ginSwagger.URL("http://127.0.0.1:8092/swagger/doc.json")
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// favicon
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	// assets file
	router.Static("/assets", "./assets")
	// Static assets like js and css files
	router.StaticFS("/download", gin.Dir("./resources", true))
	//JSON-REST API Version 1
	v1 := router.Group("/api/v1")
	{
		v1.POST("/station/dic", (&controllers.SpaceStationController{}).AddDicSpaceStation)
		v1.DELETE("/station/dic/:short_name", (&controllers.SpaceStationController{}).DestroyDicSpaceStation)
		v1.PUT("/station/dic", (&controllers.SpaceStationController{}).UpdateDicSpaceStation)
		v1.GET("/station/dic/:short_name", (&controllers.SpaceStationController{}).SelectDicSpaceStation)
		v1.GET("/station/dic", (&controllers.SpaceStationController{}).ListDicSpaceStation)
		v1.GET("/station/search", (&controllers.SpaceStationController{}).SearchDicSpaceStation)
	}

	// Default HTML page (client-side routing implemented via Vue.js)
	// router.NoRoute(func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", conf.ClientConfig())
	// })
}
