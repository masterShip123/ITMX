package customer

import (
	"fmt"
	"itmx/types"
	"testing"
)

var ids uint64

func TestCreate(t *testing.T) {
	x := NewCustomerService()
	err, message, id := x.Create(&types.Customer{
		Name: "Shippy",
		Age:  15,
	})
	fmt.Print(err, message, id)
	ids = uint64(id)
	if message != "CREATE_SUCCESS" {
		t.Error("Customer CMS TestCreate Fail")
	}
}

func TestUpdate(t *testing.T) {
	fmt.Println("Check ID : ", ids)
	x := NewCustomerService()
	err, message, _id := x.Update(ids, &types.Customer{
		Name: "ShipUpdate",
		Age:  55,
	})
	fmt.Print(err, message, _id)
	if message != "UPDATE_SUCCESS" {
		t.Error("Profile CMS TestCreate Fail")
	}
}

func TestList(t *testing.T) {
	x := NewCustomerService()
	rows := x.List(&types.Customer{})
	fmt.Print(rows)
}

func TestShow(t *testing.T) {
	x := NewCustomerService()
	data := x.Show(ids)
	fmt.Print(data)
	if data == nil {
		t.Error("Profile CMS TestShow Fail")
	}
}

func TestDelete(t *testing.T) {
	x := NewCustomerService()
	err, message, _id := x.Delete(ids)
	fmt.Print(err, message, _id)
	if message != "DELETE_SUCCESS" {
		t.Error("Profile CMS TestCreate Fail")
	}
}

// >>> AUTO_GEN_NEW_METHOD <<< //
