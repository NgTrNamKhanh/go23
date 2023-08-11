package repo

import (
	"os"

	"github.com/NgTrNamKhanh/ex05/repo/cart"
	"github.com/NgTrNamKhanh/ex05/repo/product"
	"github.com/NgTrNamKhanh/ex05/repo/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgRepo struct {
	DB      *gorm.DB
	Product product.Repo
	Cart    cart.Repo
	User    user.Repo
}

func NewRepo() Repo {
	connectString := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &pgRepo{
		DB:      db,
		Product: product.NewProductRepo(db),
		Cart:    cart.NewCartRepo(db),
		User:    user.NewUserRepo(db),
	}

}
func (r pgRepo) AutoMigrate(models ...interface{}) error {
	for idx := range models {
		if err := r.DB.AutoMigrate(models[idx]); err != nil {
			return err
		}
	}
	return nil
}
func (r pgRepo) Prod() product.Repo {
	return r.Product
}
func (r pgRepo) C() cart.Repo {
	return r.Cart
}
func (r pgRepo) Usr() user.Repo {
	return r.User
}
