package controllers

import (
	models "github.com/NgTrNamKhanh/ex05/models"
	"github.com/NgTrNamKhanh/ex05/repo"
	"github.com/gin-gonic/gin"
)

func ProductsCreate(c *gin.Context) {
	//get data from body
	var body struct {
		Name  string
		Price int
	}
	c.Bind(&body)
	//Create a product
	db := repo.NewRepo()

	p, err := db.Prod().CreateProduct(models.Product{
		Product_Name: body.Name,
		Price:        body.Price,
	})
	if err != nil {
		c.Status(400)
		return
	}
	// return
	c.JSON(200, gin.H{
		"product": p,
	})
}

func AllProducts(c *gin.Context) {
	//Get all products
	db := repo.NewRepo()

	ps, err := db.Prod().GetProductList()
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"products": ps,
	})
	// Respond with them
}
func GetProduct(c *gin.Context) {
	//Get ID of product
	id := c.Param("id")
	//Get Product
	db := repo.NewRepo()

	p, err := db.Prod().GetProductByID(id)
	//Respond
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"product": p,
	})
	// Respond with them
}
func UpdateProduct(c *gin.Context) {
	//Get ID of product
	id := c.Param("id")
	//get data from body
	var body struct {
		Name  string
		Price int
	}
	c.Bind(&body)
	newP := models.Product{
		Product_Name: body.Name,
		Price:        body.Price,
	}
	//update
	db := repo.NewRepo()
	result, err := db.Prod().UpdateProduct(id, newP)
	//Respond
	if err != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"product": result,
	})
	// Respond with them
}
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	db := repo.NewRepo()
	err := db.Prod().DeleteProduct(id)
	if err != nil {
		c.Status(400)
		return
	}
	// return
	c.Status(200)
}
