package main

import (
	"github.com/051mym/goapi/initializers"
	"github.com/051mym/goapi/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
