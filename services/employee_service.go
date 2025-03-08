package services

import (
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/repositories"
)

type EmployeeService struct {
	Repo *repositories.EmployeeRepository
}

func NewEmployeeService(repo *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{Repo: repo}
}

func (s *EmployeeService) SearchEmployeeByID(id string) (*models.Employee, error) {
	return s.Repo.SearchEmployeeByID(id)
}

func (s *EmployeeService) SearchEmployeesByName(names string) ([]models.Employee, error) {
	return s.Repo.SearchEmployeesByName(names)
}

func (s *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.Repo.GetAllEmployees()
}

func (s *EmployeeService) UpdateEmployee(id string, employeeDTO dtos.UpdateEmployeeDTO) (*models.Employee, error) {
	return s.Repo.UpdateEmployee(id, employeeDTO)
}

func (s *EmployeeService) CreateEmployee(employeeDTO dtos.UpdateEmployeeDTO) (*models.Employee, error) {
	return s.Repo.CreateEmployee(employeeDTO)
}

func (s *EmployeeService) DeleteEmployeeById(id string) error {
	return s.Repo.DeleteEmployeeById(id)
}
