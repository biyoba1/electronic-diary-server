package main

import (
	"gilab.com/pragmacreviews/golang-gin-poc/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SubjectsRoutes(router)
	routes.UsersRoutes(router)

	router.Run(":7080")

}
