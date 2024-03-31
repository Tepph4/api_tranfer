package main

import (
	"fmt"
	"go_api/controllers"
	"go_api/models"

	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

func getTransfers(c *gin.Context) {
	fmt.Print("get data")
	transfers := models.Transfer{} // Populate transfers as needed
	c.IndentedJSON(http.StatusOK, transfers)

}

// ----------------------- main -------------------------------//
func main() {
	controller.Iconnect_toDB()
	fmt.Println("api transfer")
	router := gin.Default()
	router.GET("api/v1/bplus/transferss", getTransfers)
	router.GET("api/v1/bplus/transfers", controller.ListTransfers)
	router.GET("api/v1/bplus/transferBypennding", controller.GetTransferBypennding)
	router.GET("api/v1/bplus/transfer/:orderid", controller.GetTransferByOrderId)

	router.POST("api/v1/bplus/transferCreate", controller.CreateTransfer)
	//router.POST("/Transfer", createTransfer)

	router.Run("127.0.0.1:3000")
	controller.CloseDB()
}
