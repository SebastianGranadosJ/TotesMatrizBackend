package app

import (
	"time"
	"totesbackend/config"
	"totesbackend/controllers"
	"totesbackend/controllers/utilities"
	"totesbackend/database"
	"totesbackend/repositories"
	routes "totesbackend/router"
	"totesbackend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB
var router *gin.Engine
var authUtil *utilities.AuthorizationUtil

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
	authUtil = utilities.NewAuthorizationUtil(services.NewAuthorizationService(repositories.NewAuthorizationRepository(db)))

	router = gin.Default()
	database.MigrateDB() // recordar descomentar para inicializar la base de datos

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:5503", "http://127.0.0.1:5500", "http://127.0.0.1:5501"}, // Especifica los orígenes permitidos
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Username"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	setUpUserRouter()
	setUpItemTypeRouter()
	setUpItemRouter()
	setUpPermissionRouter()
	setUpRoleRouter()
	setUpUserTypeRouter()
	setUpIdentifierTypeRouter()
	setUpUserStateTypeRouter()
	setUpEmployeeRouter()
	setUpAdditionalExpenseRouter()
	setUpHistoricalItemPriceRouter()
	setUpCommentRouter()
	setUpAuthRouter()
	setUpUserLogRouter()
	setUpAppointmentRouter()
	setUpCustomerRouter()
	setUpOrderStateTypeRouter()
	setUpPurchaseOrderRouter()
	setUpDiscountTypeRouter()
	setUpUserCredentialValidationRouter()
	setUpTaxTypeRouter()
	setUpBillingRouter()
	setUpInvoice()
	router.Run("localhost:8080")

	return nil
}

func setUpPermissionRouter() {
	permissionRepo := repositories.NewPermissionRepository(db)
	permissionService := services.NewPermissionService(permissionRepo)
	permissionController := controllers.NewPermissionController(permissionService, authUtil)
	routes.RegisterPermissionRoutes(router, permissionController)
}

func setUpEmployeeRouter() {
	employeeRepo := repositories.NewEmployeeRepository(db)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeController := controllers.NewEmployeeController(employeeService, authUtil)
	routes.RegisterEmployeeRoutes(router, employeeController)
}

func setUpRoleRouter() {
	roleRepo := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepo)
	roleController := controllers.NewRoleController(roleService, authUtil)
	routes.RegisterRoleRoutes(router, roleController)
}

func setUpItemTypeRouter() {
	itemTypeRepo := repositories.NewItemTypeRepository(db)
	itemTypeService := services.NewItemTypeService(itemTypeRepo)
	itemTypeController := controllers.NewItemTypeController(itemTypeService, authUtil)
	routes.RegisterItemTypeRoutes(router, itemTypeController)
}

func setUpUserTypeRouter() {
	userTypeRepo := repositories.NewUserTypeRepository(db)
	userTypeService := services.NewUserTypeService(userTypeRepo)
	userTypeController := controllers.NewUserTypeController(userTypeService, authUtil)
	routes.RegisterUserTypeRoutes(router, userTypeController)
}

func setUpItemRouter() {
	itemRepo := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepo)
	itemController := controllers.NewItemController(itemService, authUtil)
	routes.RegisterItemRoutes(router, itemController)
}

func setUpUserStateTypeRouter() {
	userStateTypeRepo := repositories.NewUserStateTypeRepository(db)
	userStateTypeService := services.NewUserStateTypeService(userStateTypeRepo)
	userStateTypeController := controllers.NewUserStateTypeController(userStateTypeService, authUtil)
	routes.RegisterUserStateTypeRoutes(router, userStateTypeController)
}

func setUpIdentifierTypeRouter() {
	identifierTypeRepo := repositories.NewIdentifierTypeRepository(db)
	identifierTypeService := services.NewIdentifierTypeService(identifierTypeRepo)
	identifierTypeController := controllers.NewIdentifierTypeController(identifierTypeService, authUtil)
	routes.RegisterIdentifierTypeRoutes(router, identifierTypeController)
}

func setUpUserRouter() {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService, authUtil)
	routes.RegisterUserRoutes(router, userController)
}

func setUpAdditionalExpenseRouter() {
	addRepo := repositories.NewAdditionalExpenseRepository(db)
	addService := services.NewAdditionalExpenseService(addRepo)
	addController := controllers.NewAdditionalExpenseController(addService, authUtil)
	routes.RegisterAdditionalExpenseRoutes(router, addController)
}

func setUpHistoricalItemPriceRouter() {
	hisRepo := repositories.NewHistoricalItemPriceRepository(db)
	hisService := services.NewHistoricalItemPriceService(hisRepo)
	hisController := controllers.NewHistoricalItemPriceController(hisService, authUtil)
	routes.RegisterHistoricalItemPriceRoutes(router, hisController)
}

func setUpCommentRouter() {
	commentRepo := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)
	commentController := controllers.NewCommentController(commentService, authUtil)
	routes.RegisterCommentRoutes(router, commentController)
}

func setUpAuthRouter() {
	authRepo := repositories.NewAuthorizationRepository(db)
	authService := services.NewAuthorizationService(authRepo)
	authController := controllers.NewAuthorizationController(authService)
	routes.RegisterAuthorizationRoutes(router, authController)
}

func setUpUserLogRouter() {
	userLogRepo := repositories.NewUserLogRepository(db)
	userLogService := services.NewUserLogService(userLogRepo)
	userLogController := controllers.NewUserLogController(userLogService, authUtil)
	routes.RegisterUserLogRoutes(router, userLogController)
}

func setUpAppointmentRouter() {
	appointmentRepo := repositories.NewAppointmentRepository(db)
	appointmentService := services.NewAppointmentService(appointmentRepo)
	appointmentController := controllers.NewAppointmentController(appointmentService, authUtil)
	routes.RegisterAppointmentRoutes(router, appointmentController)
}

func setUpCustomerRouter() {
	customerRepo := repositories.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepo)
	customerController := controllers.NewCustomerController(customerService, authUtil)
	routes.RegisterCustomerRoutes(router, customerController)

}

func setUpOrderStateTypeRouter() {
	orderStateTypeRepo := repositories.NewOrderStateTypeRepository(db)
	orderStateTypeService := services.NewOrderStateTypeService(orderStateTypeRepo)
	orderStateTypeController := controllers.NewOrderStateTypeController(orderStateTypeService, authUtil)
	routes.RegisterOrderStateTypeRoutes(router, orderStateTypeController)
}

func setUpPurchaseOrderRouter() {
	purchaseOrderRepo := repositories.NewPurchaseOrderRepository(db)
	itemRepo := repositories.NewItemRepository(db)
	billingRepo := repositories.NewItemRepository(db)
	discountRepo := repositories.NewDiscountTypeRepository(db)
	taxRepo := repositories.NewTaxTypeRepository(db)

	billingService := services.NewBillingService(billingRepo, discountRepo, taxRepo)
	purchaseOrderService := services.NewPurchaseOrderService(purchaseOrderRepo, itemRepo, billingService)
	purchaseOrderController := controllers.NewPurchaseOrderController(purchaseOrderService, authUtil)

	routes.RegisterPurchaseOrderRoutes(router, purchaseOrderController)
}

func setUpDiscountTypeRouter() {
	discountTypeRepo := repositories.NewDiscountTypeRepository(db)
	discountTypeService := services.NewDiscountTypeService(discountTypeRepo)
	discountTypeController := controllers.NewDiscountTypeController(discountTypeService, authUtil)
	routes.RegisterDiscountTypeRoutes(router, discountTypeController)
}

func setUpUserCredentialValidationRouter() {
	userCredentialValidationRepo := repositories.NewUserCredentialValidationRepository(db)
	userCredentialValidationService := services.NewUserCredentialValidationService(userCredentialValidationRepo)
	userCredentialValidationController := controllers.NewUserCredentialValidationController(userCredentialValidationService, authUtil)
	routes.RegisterUserCredentialValidationRoutes(router, userCredentialValidationController)
}

func setUpTaxTypeRouter() {
	taxTypeRepo := repositories.NewTaxTypeRepository(db)
	taxTypeService := services.NewTaxTypeService(taxTypeRepo)
	taxTypeController := controllers.NewTaxTypeController(taxTypeService, authUtil)
	routes.RegisterTaxTypeRoutes(router, taxTypeController)
}

func setUpBillingRouter() {
	billingRepo := repositories.NewItemRepository(db)
	discountRepo := repositories.NewDiscountTypeRepository(db)
	taxRepo := repositories.NewTaxTypeRepository(db)

	billingService := services.NewBillingService(billingRepo, discountRepo, taxRepo)
	billingController := controllers.NewBillingController(billingService, authUtil)

	routes.RegisterBillingRoutes(router, billingController)
}

func setUpInvoice() {
	invoiceRepo := repositories.NewInvoiceRepository(db)
	itemRepo := repositories.NewItemRepository(db)
	billingRepo := repositories.NewItemRepository(db)
	discountRepo := repositories.NewDiscountTypeRepository(db)
	taxRepo := repositories.NewTaxTypeRepository(db)

	billingService := services.NewBillingService(billingRepo, discountRepo, taxRepo)
	invoiceService := services.NewInvoiceService(invoiceRepo, itemRepo, billingService)
	invoiceController := controllers.NewInvoiceController(invoiceService, authUtil)

	routes.RegisterInvoice(router, invoiceController)
}
