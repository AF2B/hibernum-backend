package routers

import (
	"hibernum-backend/main/handlers/user"

	"github.com/gin-gonic/gin"
)

type SetupRouters struct {
	Router *gin.Engine
}

func NewSetupRouters() *SetupRouters {
	router := gin.Default()

	setupRouters := &SetupRouters{
		Router: router,
	}

	setupRouters.setupRoutes()

	return setupRouters
}

func (s *SetupRouters) setupRoutes() {
	s.SetupUserRouter()
}

func (s *SetupRouters) SetupUserRouter() {
	userGroup := s.Router.Group("/api/v1/users")
	{
		userGroup.POST("/", func(c *gin.Context) {
			user.UserPOSTHandler(c)
		})
		userGroup.GET("/:id", func(c *gin.Context) {
			user.UserGETIDHandler(c)
		})
		userGroup.GET("/", func(c *gin.Context) {
			user.UserGETHandler(c)
		})
	}
}

// func (s *SetupRouters) SetupProductRouter(router *gin.Engine) {
// 	productGroup := router.Group("/api/v1/products")
// 	{
// 		productGroup.POST("/", SetupProduct)
// 		productGroup.GET("/:id", GetProductByID)
// 	}
// }

// func (s *SetupRouters) SetupReservationRouter(router *gin.Engine) {
// 	reservationGroup := router.Group("/api/v1/reservations")
// 	{
// 		reservationGroup.POST("/", SetupReservation)
// 		reservationGroup.GET("/:id", GetReservationByID)
// 	}
// }
