package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pokt-foundation/cherry-picker-api/database"
	"github.com/pokt-foundation/cherry-picker-api/shared/environment"
)

var (
	port               = ":" + environment.GetString("PORT", "4200")
	dbConnectionString = environment.GetString("DB_CONNECTION_STRING", "")
)

func main() {
	_, err := database.NewPostgresDriverFromConnectionString(dbConnectionString)
	if err != nil {
		log.Fatalf("cannot connect to database: %s", err.Error())
	}

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "si",
		})
	})

	fmt.Printf("Server live and running on port %s\n", port)
	err = router.Run(port)
	if err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
