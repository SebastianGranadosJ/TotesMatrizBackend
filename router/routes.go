package routes

import (
	"totesbackend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterItemTypeRoutes(router *gin.Engine, controller *controllers.ItemTypeController) {
	router.GET("/item-type", controller.GetItemTypes)
	router.GET("/item-type/:id", controller.GetItemTypeByID)
}

func RegisterItemRoutes(router *gin.Engine, controller *controllers.ItemController) {
	router.GET("/item/:id", controller.GetItemByID)
	router.GET("/item", controller.GetAllItems)
	router.GET("/item/searchById", controller.SearchItemsByID)
	router.GET("/item/searchByName", controller.SearchItemsByName)
	router.PATCH("/item/:id/state", controller.UpdateItemState)
	router.PUT("/item/:id", controller.UpdateItem)
}

func RegisterPermissionRoutes(router *gin.Engine, controller *controllers.PermissionController) {
	router.GET("/permissions/", controller.GetAllPermissions)
	router.GET("/permissions/:id", controller.GetPermissionByID)
}

// //
func RegisterRoleRoutes(router *gin.Engine, controller *controllers.RoleController) {
	router.GET("/roles/:id", controller.GetRoleByID)
	router.GET("/roles/:id/permissions", controller.GetAllPermissionsOfRole)
	router.GET("/roles/:id/exists", controller.ExistRole)
	router.GET("/roles/", controller.GetAllRoles)
}

func RegisterUserTypeRoutes(router *gin.Engine, controller *controllers.UserTypeController) {
	router.GET("/user-types/", controller.ObtainAllUserTypes)
	router.GET("/user-types/:id", controller.ObtainUserTypeByID)
	router.GET("/user-types/:id/exists", controller.Exists)
}

func RegisterUserStateTypeRoutes(router *gin.Engine, controller *controllers.UserStateTypeController) {
	router.GET("/user-state-type", controller.GetUserStateTypes)
	router.GET("/user-state-type/:id", controller.GetUserStateTypeByID)
}

func RegisterIdentifierTypeRoutes(router *gin.Engine, controller *controllers.IdentifierTypeController) {
	router.GET("/identifier-type", controller.GetIdentifierTypes)
	router.GET("/identifier-type/:id", controller.GetIdentifierTypeByID)
}

func RegisterUserRoutes(router *gin.Engine, controller *controllers.UserController) {
	router.GET("/users", controller.GetAllUsers)
	router.GET("/users/:id", controller.GetUserByID)
	router.GET("/users/:id/exists", controller.SearchUsersByID)
	router.GET("/users/search", controller.SearchUsersByEmail)
	router.PATCH("/users/:id/state", controller.UpdateUserState)
	router.PUT("/users/:id", controller.UpdateUser)
}

func RegisterEmployeeRoutes(router *gin.Engine, controller *controllers.EmployeeController) {
	router.GET("/employees/", controller.GetAllEmployees)
	router.GET("/employees/:id", controller.GetEmployeeByID)
	router.GET("/employees/search", controller.SearchEmployeesByName)
	router.POST("/employees/", controller.CreateEmployee)
	router.PUT("/employees/:id", controller.UpdateEmployee)
	router.DELETE("/employees/:id", controller.DeleteEmployee)
}
