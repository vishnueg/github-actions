package main

import (
	"initSetupScripts/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "0.0.0.0:8082"
	log.Println("Initializing the server")
	router := gin.Default()

	router.GET("/hostname/:macaddress", controller.GetHostname)
	router.GET("/usecase/:macaddress", controller.GetAssetUsecase)

	// API paths for raspberry pi
	router.GET("/raspinitscript", controller.GetRaspInit)
	router.GET("/raspiusecase/:macaddress", controller.GetRaspiUseCase)

	// API paths for Jetson
	router.GET("/jetsoninitscript", controller.GetJetsonInit)
	router.GET("/jetsonusecase/:macaddress", controller.GetJetsonUseCase)

	// API paths for creating TAR file
	tarRoutes := router.Group("/tar")

	tarRoutes.POST("/raspi", controller.TarRaspInit)
	tarRoutes.POST("/jetson", controller.TarJetsonInit)
	tarRoutes.POST("/usecase/raspi/regular", controller.TarRaspiRegularUseCase)
	tarRoutes.POST("/usecase/raspi/football", controller.TarRaspiFootballUseCase)
	tarRoutes.POST("/usecase/jetson/:usecase", controller.TarJetsonUseCase)

	log.Println("Listening at port")
	router.Run(port)

}
