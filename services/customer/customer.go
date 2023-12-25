package customer

import (
	"fmt"
	"itmx/repositories"
	repository "itmx/repositories/customer"
	"itmx/services"
	"itmx/types"
	"itmx/utilities"
)

type profileService struct {
	Repository  repositories.ProfileRepository
	HandlerName func(fn string) string
}

func NewCustomerService() services.ProfileService {
	return &profileService{
		Repository:  repository.NewProfileRepository(),
		HandlerName: func(fn string) string { return fmt.Sprintf("Profile CMS %s Service", fn) },
	}
}

func (d *profileService) List(payload *types.Customer) *[]types.Customer {

	// Connect Database
	db := utilities.SetupDB()

	// Get List & Total
	rows := d.Repository.List(db, payload)
	return rows
}

func (d *profileService) Show(id int) *types.Customer {

	// Connect Database
	db := utilities.SetupDB()
	data := d.Repository.Show(db, id)
	return data
}

func (d *profileService) Create(payload *types.Customer) (error, string) {
	// Connect Database
	db := utilities.SetupDB()
	return d.Repository.Create(db, payload)
}

func (d *profileService) Update(id int, payload *types.Customer) (error, string, interface{}) {
	// Connect Database
	db := utilities.SetupDB()
	return d.Repository.Update(db, id, payload)
}

func (d *profileService) Delete(id int) (error, string, interface{}) {
	// Connect Database
	db := utilities.SetupDB()
	return d.Repository.Delete(db, id)
}

// >>> AUTO_GEN_NEW_METHOD <<< //
