package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"totesbackend/dtos"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeController struct {
	Service *services.EmployeeService
}

func NewEmployeeController(service *services.EmployeeService) *EmployeeController {
	return &EmployeeController{Service: service}
}

func (ec *EmployeeController) GetEmployeeByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Param("id")

	employee, err := ec.Service.SearchEmployeeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	employeeDTO := dtos.GetEmployeeDTO{
		ID:               employee.ID,
		Names:            employee.Names,
		LastNames:        employee.LastNames,
		PersonalID:       employee.PersonalID,
		Address:          employee.Address,
		PhoneNumbers:     employee.PhoneNumbers,
		UserID:           employee.UserID,
		IdentifierTypeID: employee.IdentifierTypeID,
	}

	c.JSON(http.StatusOK, employeeDTO)
}

func (ec *EmployeeController) GetAllEmployees(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	employees, err := ec.Service.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employees"})
		return
	}

	var employeesDTO []dtos.GetEmployeeDTO
	for _, employee := range employees {
		employeeDTO := dtos.GetEmployeeDTO{
			ID:               employee.ID,
			Names:            employee.Names,
			LastNames:        employee.LastNames,
			PersonalID:       employee.PersonalID,
			Address:          employee.Address,
			PhoneNumbers:     employee.PhoneNumbers,
			UserID:           employee.UserID,
			IdentifierTypeID: employee.IdentifierTypeID,
		}
		employeesDTO = append(employeesDTO, employeeDTO)
	}

	c.JSON(http.StatusOK, employeesDTO)
}

func (ec *EmployeeController) SearchEmployeesByName(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	query := c.Query("names")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	employees, err := ec.Service.SearchEmployeesByName(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employees"})
		return
	}

	if len(employees) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No employees found"})
		return
	}

	var employeesDTO []dtos.GetEmployeeDTO
	for _, employee := range employees {
		employeeDTO := dtos.GetEmployeeDTO{
			ID:               employee.ID,
			Names:            employee.Names,
			LastNames:        employee.LastNames,
			PersonalID:       employee.PersonalID,
			Address:          employee.Address,
			PhoneNumbers:     employee.PhoneNumbers,
			UserID:           employee.UserID,
			IdentifierTypeID: employee.IdentifierTypeID,
		}
		employeesDTO = append(employeesDTO, employeeDTO)
	}

	c.JSON(http.StatusOK, employeesDTO)
}

func (ec *EmployeeController) CreateEmployee(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	var dto dtos.UpdateEmployeeDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	employee, err := ec.Service.CreateEmployee(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating employee"})
		return
	}

	employeeDTO := dtos.GetEmployeeDTO{
		ID:               employee.ID,
		Names:            employee.Names,
		LastNames:        employee.LastNames,
		PersonalID:       employee.PersonalID,
		Address:          employee.Address,
		PhoneNumbers:     employee.PhoneNumbers,
		UserID:           employee.UserID,
		IdentifierTypeID: employee.IdentifierTypeID,
	}

	c.JSON(http.StatusCreated, employeeDTO)
}

func (ec *EmployeeController) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var dto dtos.UpdateEmployeeDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	employee, err := ec.Service.UpdateEmployee(id, dto)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating employee"})
		return
	}

	employeeDTO := dtos.GetEmployeeDTO{
		ID:               employee.ID,
		Names:            employee.Names,
		LastNames:        employee.LastNames,
		PersonalID:       employee.PersonalID,
		Address:          employee.Address,
		PhoneNumbers:     employee.PhoneNumbers,
		UserID:           employee.UserID,
		IdentifierTypeID: employee.IdentifierTypeID,
	}

	c.JSON(http.StatusOK, employeeDTO)
}

func (ec *EmployeeController) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	err := ec.Service.DeleteEmployeeById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
