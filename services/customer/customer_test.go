package customer

import (
	"fmt"
	"itmx/types"
	"itmx/utilities"
	"testing"
)

func TestCreate(t *testing.T) {
	utilities.SetTesterEnvironment()
	x := NewCustomerService()
	err, message, _id := x.Create(&types.Profile{
		Name: "Profile",
	})
	fmt.Print(err, message, _id)
	if message != "CREATE_SUCCESS" {
		t.Error("Profile CMS TestCreate Fail")
	}
	if _id == nil {
		t.Error("Profile CMS TestCreate Fail")
	}
}

func TestUpdate(t *testing.T) {
	utilities.SetTesterEnvironment()
	x := NewCustomerService()
	err, message, _id := x.Update("5fad99ee21a7ef41807e2598", &types.Profile{
		Name: "Profile-2",
	})
	fmt.Print(err, message, _id)
	if message != "UPDATE_SUCCESS" {
		t.Error("Profile CMS TestCreate Fail")
	}
}

func TestList(t *testing.T) {
	utilities.SetTesterEnvironment()
	x := NewCustomerService()
	rows, total := x.List(&types.IRequestProfileListPagination{
		Sort:   "",
		Order:  "",
		Offset: 0,
		Limit:  10,
		Filter: &types.IRequestProfileListPaginationFilter{
			Name: "Profile",
		},
	})
	fmt.Print(rows, total)
	if total <= 0 {
		t.Error("Profile CMS TestList Fail")
	}
}

func TestShow(t *testing.T) {
	utilities.SetTesterEnvironment()
	x := NewCustomerService()
	data := x.Show("5fad868919097408b64f02c9")
	fmt.Print(data)
	if data == nil {
		t.Error("Profile CMS TestShow Fail")
	}
}

func TestDelete(t *testing.T) {
	utilities.SetTesterEnvironment()
	x := NewCustomerService()
	err, message, _id := x.Delete("5fad99ee21a7ef41807e2598")
	fmt.Print(err, message, _id)
	if message != "DELETE_SUCCESS" {
		t.Error("Profile CMS TestCreate Fail")
	}
}

func TestToggle(t *testing.T) {
	utilities.SetTesterEnvironment()
	x := NewCustomerService()
	err, message, _id := x.ToggleActive("5fad99ee21a7ef41807e2598", &types.Profile{
		IsActive: false,
	})
	fmt.Print(err, message, _id)
	if message != "TOGGLE_ACTIVE_SUCCESS" {
		t.Error("Profile CMS TestToggle Fail")
	}
}

// >>> AUTO_GEN_NEW_METHOD <<< //
