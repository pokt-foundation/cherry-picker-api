package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pokt-foundation/cherry-picker-api/shared/environment"
)

var (
	port = ":" + environment.GetString("port", "4200")
)

func main() {

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "si",
		})
	})
	err := router.Run(port)
	if err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
