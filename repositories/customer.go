package repositories

import (
	"github.com/jinzhu/gorm"
	"itmx/types"
)

type ProfileRepository interface {
	List(db *gorm.DB, payload *types.Customer) *[]types.Customer
	Show(db *gorm.DB, _id int) *types.Customer
	Create(db *gorm.DB, payload *types.Customer) (error, string)
	Update(db *gorm.DB, _id int, payload *types.Customer) (error, string, interface{})
	Delete(db *gorm.DB, _id int) (error, string, interface{})
	// >>> AUTO_GEN_NEW_METHOD <<< //
}
