package main

import (
	"github.com/NgTrNamKhanh/ex05/controllers"
	"github.com/NgTrNamKhanh/ex05/initializers"
	models "github.com/NgTrNamKhanh/ex05/models"
	"github.com/NgTrNamKhanh/ex05/repo"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}
func main() {

	db := repo.NewRepo()

	db.AutoMigrate(&models.Product{}, &models.Cart{}, &models.CartItem{})
	r := gin.Default()
	r.POST("/product/create", controllers.ProductsCreate)
	r.GET("/product/all", controllers.AllProducts)
	r.GET("/product/:id", controllers.GetProduct)
	r.PUT("/product/:id", controllers.UpdateProduct)
	r.DELETE("/product/:id", controllers.DeleteProduct)

	r.POST("/cart/create", controllers.CartCreate)
	r.GET("/cart/:id", controllers.GetCart)
	r.PUT("/cart/add/:id", controllers.AddItemToCart)
	r.PUT("/cart/remove/:id", controllers.RemoveItemFromCart)
	r.DELETE("/cart/:id", controllers.DeleteProduct)

	r.Run()

}
