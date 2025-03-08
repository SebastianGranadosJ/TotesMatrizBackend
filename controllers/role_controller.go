package controllers

import (
	"fmt"
	"net/http"
	"totesbackend/dtos"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	Service *services.RoleService
}

func NewRoleController(service *services.RoleService) *RoleController {
	return &RoleController{Service: service}
}

func (rc *RoleController) GetAllRoles(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	roles, err := rc.Service.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving roles"})
		return
	}

	c.JSON(http.StatusOK, roles)
}

func (rc *RoleController) GetRoleByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := rc.Service.GetRoleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	permissionIDs, err := rc.Service.GetRolePermissions(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving role permissions"})
		return
	}

	// Convertir a DTO
	roleDTO := dtos.RoleDTO{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		Permissions: make([]string, len(permissionIDs)),
	}

	for i, permissionID := range permissionIDs {
		roleDTO.Permissions[i] = fmt.Sprintf("%d", permissionID)
	}

	c.JSON(http.StatusOK, roleDTO)
}

func (rc *RoleController) GetAllPermissionsOfRole(c *gin.Context) {
	roleIDParam := c.Param("id")
	var roleID uint
	if _, err := fmt.Sscanf(roleIDParam, "%d", &roleID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	permissions, err := rc.Service.GetAllPermissionsOfRole(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permissions for role"})
		return
	}

	c.JSON(http.StatusOK, permissions)
}

func (rc *RoleController) ExistRole(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	exists, err := rc.Service.ExistRole(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking role existence"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": exists})
}
