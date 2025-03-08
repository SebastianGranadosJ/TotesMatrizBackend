package repositories

import (
	"totesbackend/dtos"
	"totesbackend/models"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) SearchEmployeeByID(id string) (*models.Employee, error) {
	var employee models.Employee
	err := r.DB.Preload("User").Preload("IdentifierType").First(&employee, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepository) SearchEmployeesByName(names string) ([]models.Employee, error) {
	var employees []models.Employee
	err := r.DB.Preload("User").Preload("IdentifierType").
		Where("LOWER(names) LIKE LOWER(?)", names+"%").
		Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *EmployeeRepository) GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.DB.Preload("User").Preload("IdentifierType").Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *EmployeeRepository) UpdateEmployee(id string, employeeDTO dtos.UpdateEmployeeDTO) (*models.Employee, error) {
	var employee models.Employee
	err := r.DB.First(&employee, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	employee.Names = employeeDTO.Names
	employee.LastNames = employeeDTO.LastNames
	employee.PersonalID = employeeDTO.PersonalID
	employee.Address = employeeDTO.Address
	employee.PhoneNumbers = employeeDTO.PhoneNumbers
	employee.UserID = employeeDTO.UserID
	employee.IdentifierTypeID = employeeDTO.IdentifierTypeID

	if err := r.DB.Save(&employee).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

// agregar un empleado
func (r *EmployeeRepository) CreateEmployee(employeeDTO dtos.UpdateEmployeeDTO) (*models.Employee, error) {

	employee := models.Employee{
		Names:            employeeDTO.Names,
		LastNames:        employeeDTO.LastNames,
		PersonalID:       employeeDTO.PersonalID,
		Address:          employeeDTO.Address,
		PhoneNumbers:     employeeDTO.PhoneNumbers,
		UserID:           employeeDTO.UserID,
		IdentifierTypeID: employeeDTO.IdentifierTypeID,
	}

	err := r.DB.Create(&employee).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepository) DeleteEmployeeById(id string) error {

	err := r.DB.Delete(&models.Employee{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
