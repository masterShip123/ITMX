package services

import (
	"itmx/types"
)

type CustomerService interface {
	List(payload *types.Customer) *[]types.Customer
	Show(id uint64) *types.Customer
	Create(payload *types.Customer) (error, string, uint)
	Update(id uint64, payload *types.Customer) (error, string, interface{})
	Delete(id uint64) (error, string, interface{})
	// >>> AUTO_GEN_NEW_METHOD <<< //
}
