package repo

import (
	"github.com/NgTrNamKhanh/ex05/repo/cart"
	"github.com/NgTrNamKhanh/ex05/repo/product"
)

//"gorm.io/gorm"

type Repo interface {
	AutoMigrate(models ...interface{}) error
	Prod() product.Repo
	C() cart.Repo
}
