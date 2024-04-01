package main

import (
	"fmt"
	"go_api/controllers"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

// ----------------------- main -------------------------------//
func main() {
	controller.Iconnect_toDB()
	fmt.Println("api transfer")
	router := gin.Default()

	router.GET("api/v1/bplus/transfers", controller.ListTransfers)
	router.GET("api/v1/bplus/transferBypennding", controller.GetTransferBypennding)
	router.GET("api/v1/bplus/transfer/:orderid", controller.GetTransferByOrderId)
	router.GET("pi/v1/bplus/transfer/:buyerid",controller.GetTransferBybuyer_no)
	router.POST("api/v1/bplus/transferCreate", controller.CreateTransfer)
	//router.POST("/Transfer", createTransfer)

	router.Run("127.0.0.1:3000")
	controller.CloseDB()
}
