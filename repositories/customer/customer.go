package customer

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"itmx/repositories"
	"itmx/types"
)

type CustomerRepository struct {
	TableName   []types.Customer
	Table       *types.Customer
	HandlerName func(fn string) string
}

func NewCustomerRepository() repositories.CustomerRepository {
	return &CustomerRepository{
		TableName:   nil,
		Table:       &types.Customer{},
		HandlerName: func(fn string) string { return fmt.Sprintf("Customer %s Repository", fn) },
	}
}

func (d *CustomerRepository) List(db *gorm.DB, payload *types.Customer) *[]types.Customer {
	db.Find(&d.TableName)
	return &d.TableName
}
func (d *CustomerRepository) Show(db *gorm.DB, _id uint64) *types.Customer {
	uintID := uint(_id)
	result := db.First(d.Table, uintID)
	if result.Error != nil {
		return nil
	}
	return d.Table
}

func (d *CustomerRepository) Create(db *gorm.DB, payload *types.Customer) (error, string, uint) {
	db.Create(&payload)
	return nil, "CREATE_SUCCESS", payload.ID
}

func (d *CustomerRepository) Update(db *gorm.DB, _id uint64, payload *types.Customer) (error, string, interface{}) {
	// Get model if exist
	uintID := uint(_id)
	if err := db.First(d.Table, uintID).Error; err != nil {
		return nil, "DATA_NOT_FOUND", _id
	}
	db.Model(&d.Table).Updates(payload)
	return nil, "UPDATE_SUCCESS", _id
}

func (d *CustomerRepository) Delete(db *gorm.DB, _id uint64) (error, string, interface{}) {
	// Get model if exist
	uintID := uint(_id)
	if err := db.Where("id = ?", uintID).First(&d.Table).Error; err != nil {
		return nil, "DATA_NOT_FOUND", _id
	}
	db.Delete(&d.Table)
	return nil, "DELETE_SUCCESS", _id
}

// >>> AUTO_GEN_NEW_METHOD <<< //
