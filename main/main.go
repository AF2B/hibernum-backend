package main

import (
	"hibernum-backend/main/routers"
)

func main() {
	stup := routers.NewSetupRouters()
	stup.Router.Run(":8080")
}
