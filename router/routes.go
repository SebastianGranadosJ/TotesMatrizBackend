package routes

import (
	"totesbackend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterItemTypeRoutes(router *gin.Engine, controller *controllers.ItemTypeController) {
	router.GET("/item-types", controller.GetItemTypes)
	router.GET("/item-types/:id", controller.GetItemTypeByID)
}

func RegisterItemRoutes(router *gin.Engine, controller *controllers.ItemController) {
	router.GET("/items/:id", controller.GetItemByID)
	router.GET("/items", controller.GetAllItems)
	router.GET("/items/searchById", controller.SearchItemsByID)
	router.GET("/items/searchByName", controller.SearchItemsByName)
	router.PATCH("/items/:id/state", controller.UpdateItemState)
	router.PUT("/items/:id", controller.UpdateItem)
	router.POST("/items", controller.CreateItem)
	router.GET("/items/:id/stock", controller.CheckItemStock)
}

func RegisterPermissionRoutes(router *gin.Engine,
	controller *controllers.PermissionController) {

	router.GET("/permissions", controller.GetAllPermissions)
	router.GET("/permissions/:id", controller.GetPermissionByID)
	router.GET("/permissions/searchByID", controller.SearchPermissionsByID)
	router.GET("/permissions/searchByName", controller.SearchPermissionsByName)
}

func RegisterRoleRoutes(router *gin.Engine, controller *controllers.RoleController) {
	router.GET("/roles/:id", controller.GetRoleByID)
	router.GET("/roles/:id/permission", controller.GetAllPermissionsOfRole)
	router.GET("/roles/:id/exist", controller.ExistRole)
	router.GET("/roles", controller.GetAllRoles)
	router.GET("/roles/searchByID", controller.SearchRolesByID)
	router.GET("/roles/searchByName", controller.SearchRolesByName)
}

func RegisterUserTypeRoutes(router *gin.Engine,
	controller *controllers.UserTypeController) {
	router.GET("/user-types", controller.GetAllUserTypes)
	router.GET("/user-types/:id", controller.GetUserTypeByID)
	router.GET("/user-types/:id/exists", controller.ExistsUserType)
	router.GET("/user-types/searchByID", controller.SearchUserTypesByID)
	router.GET("/user-types/searchByName", controller.SearchUserTypesByName)
}

func RegisterUserStateTypeRoutes(router *gin.Engine,
	controller *controllers.UserStateTypeController) {
	router.GET("/user-state-types", controller.GetAllUserStateTypes)
	router.GET("/user-state-types/:id", controller.GetUserStateTypeByID)
}

func RegisterIdentifierTypeRoutes(router *gin.Engine, controller *controllers.IdentifierTypeController) {
	router.GET("/identifier-types", controller.GetAllIdentifierTypes)
	router.GET("/identifier-types/:id", controller.GetIdentifierTypeByID)
}

func RegisterUserRoutes(router *gin.Engine,
	controller *controllers.UserController) {
	router.GET("/users", controller.GetAllUsers)
	router.GET("/users/:id", controller.GetUserByID)
	router.GET("/users/searchByID", controller.SearchUsersByID)
	router.GET("/users/searchByEmail", controller.SearchUsersByEmail)
	router.PATCH("/users/:id/state", controller.UpdateUserState)
	router.PUT("/users/:id", controller.UpdateUser)
	router.POST("/users", controller.CreateUser)
}

func RegisterEmployeeRoutes(router *gin.Engine, controller *controllers.EmployeeController) {
	router.GET("/employees/:id", controller.GetEmployeeByID)
	router.GET("/employees", controller.GetAllEmployees)
	router.GET("/employees/searchByID", controller.SearchEmployeesByID)
	router.GET("/employees/searchByName", controller.SearchEmployeesByName)
	router.POST("/employees", controller.CreateEmployee)
	router.PUT("/employees/:id", controller.UpdateEmployee)
}

func RegisterAdditionalExpenseRoutes(router *gin.Engine,
	controller *controllers.AdditionalExpenseController) {
	router.GET("/additional-expenses", controller.GetAllAdditionalExpenses)
	router.GET("/additional-expenses/:id", controller.GetAdditionalExpenseByID)
	router.POST("/additional-expenses", controller.CreateAdditionalExpense)
	router.PUT("/additional-expenses/:id", controller.UpdateAdditionalExpense)
	router.DELETE("/additional-expenses/:id", controller.DeleteAdditionalExpense)
}

func RegisterHistoricalItemPriceRoutes(router *gin.Engine, controller *controllers.HistoricalItemPriceController) {
	router.GET("/historical-item-prices/:id", controller.GetHistoricalItemPrice)
}

func RegisterCommentRoutes(router *gin.Engine,
	controller *controllers.CommentController) {
	router.GET("/comments/:id", controller.GetCommentByID)
	router.GET("/comments", controller.GetAllComments)
	router.GET("/comments/searchByID", controller.SearchCommentsByID)
	router.GET("/comments/searchByName", controller.SearchCommentsByName)
	router.GET("/comments/searchByEmail", controller.SearchCommentsByEmail)
	router.POST("/comments", controller.CreateComment)
	router.PUT("/comments/:id", controller.UpdateComment)
}

func RegisterAuthorizationRoutes(router *gin.Engine, controller *controllers.AuthorizationController) {
	router.GET("/auth/check-permission", controller.CheckUserPermission)
}

func RegisterAppointmentRoutes(router *gin.Engine, controller *controllers.AppointmentController) {
	router.GET("/appointments/:id", controller.GetAppointmentByID)
	router.GET("/appointments", controller.GetAllAppointments)
	router.GET("/appointments/searchByID", controller.SearchAppointmentsByID)
	router.GET("/appointments/searchByCustomerID", controller.SearchAppointmentsByCustomerID)
	router.GET("/appointments/searchByState", controller.SearchAppointmentsByState)
	router.GET("/appointments/customer/:customerID", controller.GetAppointmentsByCustomerID)
	router.POST("/appointments", controller.CreateAppointment)
	router.PUT("/appointments/:id", controller.UpdateAppointment)
	router.GET("/appointments/byCustomerAndDate", controller.GetAppointmentByCustomerIDAndDate)
}

func RegisterCustomerRoutes(router *gin.Engine, controller *controllers.CustomerController) {
	router.GET("/customers/:id", controller.GetCustomerByID)
	router.GET("/customers/customerID/:customerID", controller.GetCustomerByCustomerID)
	router.GET("/customers", controller.GetAllCustomers)
	router.GET("/customers/email/:email", controller.GetCustomerByEmail)
	router.GET("/customers/searchByID", controller.SearchCustomersByID)
	router.GET("/customers/searchByName", controller.SearchCustomersByName)
	router.GET("/customers/searchByLastName", controller.SearchCustomersByLastName)
	router.POST("/customers", controller.CreateCustomer)
	router.PUT("/customers/:id", controller.UpdateCustomer)
}

func RegisterOrderStateTypeRoutes(router *gin.Engine, controller *controllers.OrderStateTypeController) {
	router.GET("/order-state-types", controller.GetAllOrderStateTypes)
	router.GET("/order-state-types/:id", controller.GetOrderStateTypeByID)
}

func RegisterPurchaseOrderRoutes(router *gin.Engine, controller *controllers.PurchaseOrderController) {

	router.GET("/purchase-orders/:id", controller.GetPurchaseOrderByID)
	router.GET("/purchase-orders", controller.GetAllPurchaseOrders)
	router.GET("/purchase-orders/searchByID", controller.SearchPurchaseOrdersByID)
	router.GET("/purchase-orders/customers/:customerID", controller.GetPurchaseOrdersByCustomerID)
	router.GET("/purchase-orders/seller/:sellerID", controller.GetPurchaseOrdersBySellerID)
	router.GET("/purchase-orders/state/:stateID", controller.GetPurchaseOrdersByStateID)
	router.POST("/purchase-orders", controller.CreatePurchaseOrder)
	router.PATCH("/purchase-orders/:id/state", controller.ChangePurchaseOrderState)

}

func RegisterDiscountTypeRoutes(router *gin.Engine, controller *controllers.DiscountTypeController) {
	router.GET("/discount-types", controller.GetAllDiscountTypes)
	router.GET("/discount-types/:id", controller.GetDiscountTypeByID)
}

func RegisterUserCredentialValidationRoutes(router *gin.Engine, controller *controllers.UserCredentialValidationController) {
	router.POST("/user-credential-validation", controller.ValidateUserCredentials)
}

func RegisterTaxTypeRoutes(router *gin.Engine, controller *controllers.TaxTypeController) {
	router.GET("/tax-types", controller.GetAllTaxTypes)
	router.GET("/tax-types/:id", controller.GetTaxTypeByID)
}

func RegisterBillingRoutes(router *gin.Engine, controller *controllers.BillingController) {
	router.POST("/billing/subtotal", controller.CalculateSubtotal)
	router.POST("/billing/total", controller.CalculateTotal)
}

func RegisterInvoice(router *gin.Engine, controller *controllers.InvoiceController) {
	router.GET("/invoices/:id", controller.GetInvoiceByID)
	router.GET("/invoices", controller.GetAllInvoices)
	router.GET("/invoices/searchById", controller.SearchInvoiceByID)
	router.GET("/invoices/searchByPersonalId", controller.SearchInvoiceByCustomerPersonalId)
	router.POST("/invoices", controller.CreateInvoice)
}
