package main

import (
	"server/models"
	"server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.InitDB()
	router.InitRouter(r)
	r.Run("0.0.0.0:8080")
}
