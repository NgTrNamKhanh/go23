package cart

import (
	"errors"

	model "github.com/NgTrNamKhanh/ex05/models"
	"gorm.io/gorm"
)

type Repo interface {
	CreateCart(p model.Cart) (*model.Cart, error)
	AddProduct(ID string, Quantity int) error
	RemoveProduct(ID string) error
	DeleteCart(ID string) error
	GetCartByID(ID string) (*model.Cart, error)
}

func NewCartRepo(db *gorm.DB) Repo {
	return &cRepo{
		DB: db,
	}
}

type cRepo struct {
	DB *gorm.DB
}

var cart1 string

// AddProduct implements Repo.
func (p cRepo) AddProduct(ProductID string, Quantity int) error {
	cart1 = "2"
	var cart_item model.CartItem
	c_i := p.DB.Where("cart_id = ? AND product_id = ?", cart1, ProductID).First(&cart_item)
	if c_i.Error != gorm.ErrRecordNotFound {
		result := p.DB.Model(&cart_item).Update("quantity", cart_item.Quantity+Quantity)
		if result.Error != nil {
			return result.Error
		}
		return nil
	} else {
		product := model.Product{}
		p.DB.Table("products").Where("id = ?", ProductID).First(&product)
		cart_item.CartID = cart1
		cart_item.ProductID = ProductID
		cart_item.Quantity = Quantity
		err := p.DB.Model(&cart_item).Create(&cart_item).Error
		if err != nil {
			return err
		}
		return nil
	}
}

// RemoveProduct implements Repo.
func (p cRepo) RemoveProduct(ProductID string) error {
	var cart_item model.CartItem
	p.DB.Where("cart_id = ? AND product_id = ?", cart1, ProductID).First(&cart_item)
	if cart_item.ID == 0 {
		return errors.New("Item not in cart")
	} else if cart_item.Quantity == 1 {
		err := p.DeleteCart(cart1)
		if err != nil {
			return err
		}
		return nil
	} else {
		result := p.DB.Model(&cart_item).Update("quantity", cart_item.Quantity-1)
		if result.Error != nil {
			return result.Error
		}
		return nil
	}
}

// CreateCart implements Repo.
func (r cRepo) CreateCart(p model.Cart) (*model.Cart, error) {
	return &p, r.DB.Create(&p).Error
}

// DeleteCart implements Repo.
func (r cRepo) DeleteCart(ID string) error {
	p := model.Cart{}
	return r.DB.Table("carts").Where("id = ?", ID).Delete(&p).Error
}

// GetCartByID implements Repo.
func (r cRepo) GetCartByID(ID string) (*model.Cart, error) {
	cart := model.Cart{}

	// Load cart data along with associated cart items
	result := r.DB.Table("carts").Where("id = ?", ID).Preload("CartItem").First(&cart)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cart, nil
}
