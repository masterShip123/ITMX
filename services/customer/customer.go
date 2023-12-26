package customer

import (
	"fmt"
	"itmx/repositories"
	repository "itmx/repositories/customer"
	"itmx/services"
	"itmx/types"
	"itmx/utilities"
)

type CustomerService struct {
	Repository  repositories.CustomerRepository
	HandlerName func(fn string) string
}

func NewCustomerService() services.CustomerService {
	return &CustomerService{
		Repository:  repository.NewCustomerRepository(),
		HandlerName: func(fn string) string { return fmt.Sprintf("Customer %s Service", fn) },
	}
}

func (d *CustomerService) List(payload *types.Customer) *[]types.Customer {

	// Connect Database
	db := utilities.SetupDB()

	// Get List & Total
	rows := d.Repository.List(db, payload)
	return rows
}

func (d *CustomerService) Show(id uint64) *types.Customer {

	// Connect Database
	db := utilities.SetupDB()
	data := d.Repository.Show(db, id)
	return data
}

func (d *CustomerService) Create(payload *types.Customer) (error, string, uint) {
	// Connect Database
	db := utilities.SetupDB()
	return d.Repository.Create(db, payload)
}

func (d *CustomerService) Update(id uint64, payload *types.Customer) (error, string, interface{}) {
	// Connect Database
	db := utilities.SetupDB()
	return d.Repository.Update(db, id, payload)
}

func (d *CustomerService) Delete(id uint64) (error, string, interface{}) {
	// Connect Database
	db := utilities.SetupDB()
	return d.Repository.Delete(db, id)
}

// >>> AUTO_GEN_NEW_METHOD <<< //
