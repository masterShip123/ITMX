package routers

import (
	"github.com/gin-gonic/gin"
	service "itmx/services/customer"
	"itmx/types"
	"net/http"
	"strconv"
)

func CreateCustomer(c *gin.Context) {
	sv := service.NewCustomerService()
	var payload *types.Customer
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, _, _id := sv.Create(payload)

	c.JSON(http.StatusCreated, _id)
}

func GetCustomer(c *gin.Context) {
	sv := service.NewCustomerService()
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	data := sv.Show(id)
	if data == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetCustomers(c *gin.Context) {
	sv := service.NewCustomerService()
	var payload *types.Customer
	defer func() { payload = nil }()
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list := sv.List(payload)

	c.JSON(http.StatusOK, list)
}

func UpdateCustomer(c *gin.Context) {
	sv := service.NewCustomerService()
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var payload *types.Customer
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, message, _id := sv.Update(id, payload)
	if message == "DATA_NOT_FOUND" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, _id)
}

func DeleteCustomer(c *gin.Context) {
	sv := service.NewCustomerService()
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	_, message, _id := sv.Delete(id)
	if message == "DATA_NOT_FOUND" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, _id)
}
