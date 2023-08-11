package controllers

import (
	models "github.com/NgTrNamKhanh/ex05/models"
	"github.com/NgTrNamKhanh/ex05/repo"
	"github.com/gin-gonic/gin"
)

func CartCreate(c *gin.Context) {
	//Create a cart
	db := repo.NewRepo()

	p, err := db.C().CreateCart(models.Cart{})
	if err != nil {
		c.Status(400)
		return
	}
	// return
	c.JSON(200, gin.H{
		"cart": p,
	})
}

func GetCart(c *gin.Context) {
	//Get ID of cartuct
	id := c.Param("id")
	//Get Cartuct
	db := repo.NewRepo()
	p, err := db.C().GetCartByID(id)
	//Respond
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"cart": p,
		
	})
	// Respond with them
}
func AddItemToCart(c *gin.Context) {
	//Get ID of cartuct
	id := c.Param("id")
	quantity := 1
	//update
	db := repo.NewRepo()
	err := db.C().AddProduct(id, quantity)
	//Respond
	if err != nil {
		c.Status(400)
		return
	}
	c.Status(200)
	// Respond with them
}
func RemoveItemFromCart(c *gin.Context) {
	//Get ID of cartuct
	id := c.Param("id")
	//update
	db := repo.NewRepo()
	err := db.C().RemoveProduct(id)
	//Respond
	if err != nil {
		c.Status(400)
		return
	}
	c.Status(200)
	// Respond with them
}
func DeleteCart(c *gin.Context) {
	id := c.Param("id")
	db := repo.NewRepo()
	err := db.C().DeleteCart(id)
	if err != nil {
		c.Status(400)
		return
	}
	// return
	c.Status(200)
}
