package app

import (
	"totesbackend/config"
	"totesbackend/controllers"
	"totesbackend/database"
	"totesbackend/repositories"
	routes "totesbackend/router"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB
var router *gin.Engine

func SetupAndRunApp() error {
	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// start database
	err = database.StartPostgres()
	if err != nil {
		return err
	}

	// defer closing database
	defer database.ClosePostgres()

	db = database.GetDB()
	router = gin.Default()
	database.MigrateDB()

	setUpUserRouter()
	setUpItemTypeRouter()
	setUpItemRouter()
	setUpPermissionRouter()
	setUpRoleRouter()
	setUpUserTypeRouter()
	setUpIdentifierTypeRouter()
	setUpUserStateTypeRouter()
	setUpEmployeeRouter()

	router.Run("localhost:8080")

	return nil
}
func setUpPermissionRouter() {
	permissionRepo := repositories.NewPermissionRepository(db)
	permissionService := services.NewPermissionService(permissionRepo)
	permissionController := controllers.NewPermissionController(permissionService)
	routes.RegisterPermissionRoutes(router, permissionController)
}

func setUpEmployeeRouter() {
	employeeRepo := repositories.NewEmployeeRepository(db)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeController := controllers.NewEmployeeController(employeeService)
	routes.RegisterEmployeeRoutes(router, employeeController)
}

func setUpRoleRouter() {
	roleRepo := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepo)
	roleController := controllers.NewRoleController(roleService)
	routes.RegisterRoleRoutes(router, roleController)
}

func setUpItemTypeRouter() {
	itemTypeRepo := repositories.NewItemTypeRepository(db)
	itemTypeService := services.NewItemTypeService(itemTypeRepo)
	itemTypeController := controllers.NewItemTypeController(itemTypeService)
	routes.RegisterItemTypeRoutes(router, itemTypeController)
}

func setUpUserTypeRouter() {
	userTypeRepo := repositories.NewUserTypeRepository(db)
	userTypeService := services.NewUserTypeService(userTypeRepo)
	userTypeController := controllers.NewUserTypeController(userTypeService)
	routes.RegisterUserTypeRoutes(router, userTypeController)
}

func setUpItemRouter() {
	itemRepo := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepo)
	itemController := controllers.NewItemController(itemService)
	routes.RegisterItemRoutes(router, itemController)
}

func setUpUserStateTypeRouter() {
	userStateTypeRepo := repositories.NewUserStateTypeRepository(db)
	userStateTypeService := services.NewUserStateTypeService(userStateTypeRepo)
	userStateTypeController := controllers.NewUserStateTypeController(userStateTypeService)
	routes.RegisterUserStateTypeRoutes(router, userStateTypeController)
}

func setUpIdentifierTypeRouter() {
	identifierTypeRepo := repositories.NewIdentifierTypeRepository(db)
	identifierTypeService := services.NewIdentifierTypeService(identifierTypeRepo)
	identifierTypeController := controllers.NewIdentifierTypeController(identifierTypeService)
	routes.RegisterIdentifierTypeRoutes(router, identifierTypeController)
}

func setUpUserRouter() {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	routes.RegisterUserRoutes(router, userController)
}
