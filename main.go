package main

import (
	"github.com/051mym/goapi/initializers"
	"github.com/051mym/goapi/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	routes.PostRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:3000 -> set in env
}
