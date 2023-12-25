package customer

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"itmx/repositories"
	"itmx/types"
)

type profileRepository struct {
	TableName   []types.Customer
	Table       *types.Customer
	HandlerName func(fn string) string
}

func NewProfileRepository() repositories.ProfileRepository {
	return &profileRepository{
		TableName:   nil,
		Table:       nil,
		HandlerName: func(fn string) string { return fmt.Sprintf("Profile CMS %s Repository", fn) },
	}
}

func (d *profileRepository) List(db *gorm.DB, payload *types.Customer) *[]types.Customer {
	db.Find(&payload)

	return &d.TableName
}
func (d *profileRepository) Show(db *gorm.DB, _id int) *types.Customer {
	result := db.First(&types.Customer{}, _id)
	if result.Error != nil {
		return nil
	}
	return d.Table
}

func (d *profileRepository) Create(db *gorm.DB, payload *types.Customer) (error, string) {
	db.Create(&payload)
	return nil, "CREATE_SUCCESS"
}

func (d *profileRepository) Update(db *gorm.DB, _id int, payload *types.Customer) (error, string, interface{}) {
	// Get model if exist
	if err := db.Where("id = ?", _id).First(&d.Table).Error; err != nil {
		return nil, "DATA_NOT_FOUND", _id
	}
	db.Model(&d.Table).Updates(payload)
	return nil, "UPDATE_SUCCESS", _id
}

func (d *profileRepository) Delete(db *gorm.DB, _id int) (error, string, interface{}) {
	// Get model if exist
	if err := db.Where("id = ?", _id).First(&d.Table).Error; err != nil {
		return nil, "DATA_NOT_FOUND", _id
	}
	db.Delete(&d.Table)
	return nil, "DELETE_SUCCESS", _id
}

// >>> AUTO_GEN_NEW_METHOD <<< //
