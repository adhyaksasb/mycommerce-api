package main

import (
	"github.com/adhyaksasb/mycommerce-api/initializers"
	model "github.com/adhyaksasb/mycommerce-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.Post{})
}