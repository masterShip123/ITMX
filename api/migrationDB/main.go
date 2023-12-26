package main

import (
	"fmt"
	"itmx/types"
	"itmx/utilities"
)

func main() {
	db := utilities.SetupDB()

	// Insert sample data
	customers := []types.Customer{
		{Name: "John Doe", Age: 30},
		{Name: "Jane Smith", Age: 25},
		{Name: "Bob Johnson", Age: 35},
	}

	for _, customer := range customers {
		result := db.Create(&customer)
		if result.Error != nil {
			fmt.Printf("Error inserting customer %v: %v\n", customer, result.Error)
		} else {
			fmt.Printf("Customer inserted successfully: %v\n", customer)
		}
	}
}
