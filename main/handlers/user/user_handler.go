package user

import (
	"fmt"
	"hibernum-backend/main/config"

	"github.com/gin-gonic/gin"
)

func UserGETHandler(gin *gin.Context) {
	isWorks := gin.Query("works")
	fmt.Printf("works: %s\n", isWorks)
	config.FakeData()
}

func UserGETIDHandler(gin *gin.Context) {
	id := gin.Param("id")
	fmt.Printf("id: ", id)
}

func UserPOSTHandler(gin *gin.Context) {
	fmt.Println("works!")
}
