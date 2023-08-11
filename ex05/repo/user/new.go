package user

import (
	model "github.com/NgTrNamKhanh/ex05/models"
	"gorm.io/gorm"
)

type Repo interface {
	CreateUser(p model.User) (*model.User, error)
	GetUserByID(ID interface{}) (*model.User, error)
	GetUserByEmail(Email string) (*model.User, error)
	GetUserList() ([]model.User, error)
	DeleteUser(ID string) error
}

func NewUserRepo(db *gorm.DB) Repo {
	return &usrRepo{
		DB: db,
	}
}

type usrRepo struct {
	DB *gorm.DB
}

// CreateUser implements Repo.
func (r usrRepo) CreateUser(p model.User) (*model.User, error) {
	return &p, r.DB.Create(&p).Error
}

// DeleteUser implements Repo.
func (r usrRepo) DeleteUser(ID string) error {
	p := model.User{}
	return r.DB.Table("users").Where("id = ?", ID).Delete(&p).Error
}

// GetUserByID implements Repo.
func (r usrRepo) GetUserByID(ID interface{}) (*model.User, error) {
	p := model.User{}
	return &p, r.DB.Table("users").First(&p, ID).Error
}
func (r usrRepo) GetUserByEmail(Email string) (*model.User, error) {
	p := model.User{}
	return &p, r.DB.Table("users").Where("email = ?", Email).First(&p).Error
}

// GetUserList implements Repo.
func (r usrRepo) GetUserList() ([]model.User, error) {
	rs := []model.User{}
	return rs, r.DB.Table("users").Find(&rs).Error
}
