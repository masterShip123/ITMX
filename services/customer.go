package services

import (
	"itmx/types"
)

type ProfileService interface {
	List(payload *types.Customer) *[]types.Customer
	Show(id int) *types.Customer
	Create(payload *types.Customer) (error, string)
	Update(id int, payload *types.Customer) (error, string, interface{})
	Delete(id int) (error, string, interface{})
	// >>> AUTO_GEN_NEW_METHOD <<< //
}
