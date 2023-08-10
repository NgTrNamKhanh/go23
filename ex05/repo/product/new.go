package product

import (
	model "github.com/NgTrNamKhanh/ex05/models"
	"gorm.io/gorm"
)

type Repo interface {
	CreateProduct(p model.Product) (*model.Product, error)
	UpdateProduct(ID string, newP model.Product) (*model.Product, error)
	GetProductByID(ID string) (*model.Product, error)
	GetProductList() ([]model.Product, error)
	DeleteProduct(ID string) error
}

func NewProductRepo(db *gorm.DB) Repo {
	return &pdRepo{
		DB: db,
	}
}

type pdRepo struct {
	DB *gorm.DB
}

// UpdateProduct implements Repo.
func (r pdRepo) UpdateProduct(ID string, newp model.Product) (*model.Product, error) {
	var p model.Product
	r.DB.First(&p, ID)
	return &p, r.DB.Model(&p).Updates(model.Product{
		Product_Name: newp.Product_Name,
		Price:        newp.Price,
	}).Error
}

// CreateProduct implements Repo.
func (r pdRepo) CreateProduct(p model.Product) (*model.Product, error) {
	return &p, r.DB.Create(&p).Error
}

// DeleteProduct implements Repo.
func (r pdRepo) DeleteProduct(ID string) error {
	p := model.Product{}
	return r.DB.Table("products").Where("id = ?", ID).Delete(&p).Error
}

// GetProductByID implements Repo.
func (r pdRepo) GetProductByID(ID string) (*model.Product, error) {
	p := model.Product{}
	return &p, r.DB.Table("products").Where("id = ?", ID).First(&p).Error
}

// GetProductList implements Repo.
func (r pdRepo) GetProductList() ([]model.Product, error) {
	rs := []model.Product{}
	return rs, r.DB.Table("products").Find(&rs).Error
}
