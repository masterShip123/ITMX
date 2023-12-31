package repositories

import (
	"github.com/jinzhu/gorm"
	"itmx/types"
)

type CustomerRepository interface {
	List(db *gorm.DB, payload *types.Customer) *[]types.Customer
	Show(db *gorm.DB, _id uint64) *types.Customer
	Create(db *gorm.DB, payload *types.Customer) (error, string, uint)
	Update(db *gorm.DB, _id uint64, payload *types.Customer) (error, string, interface{})
	Delete(db *gorm.DB, _id uint64) (error, string, interface{})
	// >>> AUTO_GEN_NEW_METHOD <<< //
}
