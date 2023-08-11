package main

import (
	"github.com/NgTrNamKhanh/ex05/controllers"
	"github.com/NgTrNamKhanh/ex05/initializers"
	"github.com/NgTrNamKhanh/ex05/middleware"
	model "github.com/NgTrNamKhanh/ex05/models"
	models "github.com/NgTrNamKhanh/ex05/models"
	"github.com/NgTrNamKhanh/ex05/repo"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}
func main() {

	db := repo.NewRepo()

	db.AutoMigrate(&models.Product{}, &models.Cart{}, &models.CartItem{}, &model.User{})
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	productPortal := r.Group("/product", middleware.RequireAuth)
	{
		productPortal.POST("/create", controllers.ProductsCreate)
		productPortal.GET("/all", controllers.AllProducts)
		productPortal.GET("/:id", controllers.GetProduct)
		productPortal.PUT("/:id", controllers.UpdateProduct)
		productPortal.DELETE("/:id", controllers.DeleteProduct)
	}

	cartPortal := r.Group("cart", middleware.RequireAuth)
	{
		cartPortal.POST("/create", controllers.CartCreate)
		cartPortal.GET("/:id", controllers.GetCart)
		cartPortal.PUT("/add/:id", controllers.AddItemToCart)
		cartPortal.PUT("/remove/:id", controllers.RemoveItemFromCart)
		cartPortal.DELETE("/:id", controllers.DeleteProduct)
	}

	r.Run()

}
