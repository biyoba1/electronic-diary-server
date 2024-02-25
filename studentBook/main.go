package main

import (
	"gilab.com/pragmacreviews/golang-gin-poc/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	routes.SubjectsRoutes(router)
	routes.UsersRoutes(router)
	routes.GradesRoutes(router)
	router.Run(":7080")
}
