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
	router.POST("/item", controller.CreateItem)
	router.GET("/item/:id/stock", controller.CheckItemStock)
}

func RegisterPermissionRoutes(router *gin.Engine,
	controller *controllers.PermissionController) {

	router.GET("/permission", controller.GetAllPermissions)
	router.GET("/permission/:id", controller.GetPermissionByID)
	router.GET("/permission/searchByID", controller.SearchPermissionsByID)
	router.GET("/permission/searchByName", controller.SearchPermissionsByName)
}

func RegisterRoleRoutes(router *gin.Engine, controller *controllers.RoleController) {
	router.GET("/role/:id", controller.GetRoleByID)
	router.GET("/role/:id/permission", controller.GetAllPermissionsOfRole)
	router.GET("/role/:id/exist", controller.ExistRole)
	router.GET("/role", controller.GetAllRoles)
	router.GET("/role/searchByID", controller.SearchRolesByID)
	router.GET("/role/searchByName", controller.SearchRolesByName)
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
	router.GET("/user-state-type", controller.GetAllUserStateTypes)
	router.GET("/user-state-type/:id", controller.GetUserStateTypeByID)
}

func RegisterIdentifierTypeRoutes(router *gin.Engine, controller *controllers.IdentifierTypeController) {
	router.GET("/identifier-type", controller.GetAllIdentifierTypes)
	router.GET("/identifier-type/:id", controller.GetIdentifierTypeByID)
}

func RegisterUserRoutes(router *gin.Engine,
	controller *controllers.UserController) {
	router.GET("/user", controller.GetAllUsers)
	router.GET("/user/:id", controller.GetUserByID)
	router.GET("/user/searchByID", controller.SearchUsersByID)
	router.GET("/user/searchByEmail", controller.SearchUsersByEmail)
	router.PATCH("/user/:id/state", controller.UpdateUserState)
	router.PUT("/user/:id", controller.UpdateUser)
	router.POST("/user", controller.CreateUser)
}

func RegisterEmployeeRoutes(router *gin.Engine, controller *controllers.EmployeeController) {
	router.GET("/employee/:id", controller.GetEmployeeByID)
	router.GET("/employee", controller.GetAllEmployees)
	router.GET("/employee/searchByID", controller.SearchEmployeesByID)
	router.GET("/employee/searchByName", controller.SearchEmployeesByName)
	router.POST("/employee", controller.CreateEmployee)
	router.PUT("/employee/:id", controller.UpdateEmployee)
}

func RegisterAdditionalExpenseRoutes(router *gin.Engine,
	controller *controllers.AdditionalExpenseController) {
	router.GET("/additional-expense", controller.GetAllAdditionalExpenses)
	router.GET("/additional-expense/:id", controller.GetAdditionalExpenseByID)
	router.POST("/additional-expense", controller.CreateAdditionalExpense)
	router.PUT("/additional-expense/:id", controller.UpdateAdditionalExpense)
	router.DELETE("/additional-expense/:id", controller.DeleteAdditionalExpense)
}

func RegisterHistoricalItemPriceRoutes(router *gin.Engine, controller *controllers.HistoricalItemPriceController) {
	router.GET("/historical-item-price/:id", controller.GetHistoricalItemPrice)
}

func RegisterCommentRoutes(router *gin.Engine,
	controller *controllers.CommentController) {
	router.GET("/comment/:id", controller.GetCommentByID)
	router.GET("/comments", controller.GetAllComments)
	router.GET("/comments/searchByID", controller.SearchCommentsByID)
	router.GET("/comments/searchByName", controller.SearchCommentsByName)
	router.GET("/comments/searchByEmail", controller.SearchCommentsByEmail)
	router.POST("/comment", controller.CreateComment)
	router.PUT("/comment/:id", controller.UpdateComment)
}

func RegisterAuthorizationRoutes(router *gin.Engine, controller *controllers.AuthorizationController) {
	router.GET("/auth/check-permission", controller.CheckUserPermission)
}

func RegisterUserLogRoutes(router *gin.Engine, controller *controllers.UserLogController) {
	router.GET("/user-log/:id", controller.GetAllLogsFromUser)
}

func RegisterAppointmentRoutes(router *gin.Engine, controller *controllers.AppointmentController) {
	router.GET("/appointment/:id", controller.GetAppointmentByID)
	router.GET("/appointments", controller.GetAllAppointments)
	router.GET("/appointments/searchByID", controller.SearchAppointmentsByID)
	router.GET("/appointments/searchByCustomerID", controller.SearchAppointmentsByCustomerID)
	router.GET("/appointments/searchByState", controller.SearchAppointmentsByState)
	router.GET("/appointment/customer/:customerID", controller.GetAppointmentsByCustomerID)
	router.POST("/appointment", controller.CreateAppointment)
	router.PUT("/appointment/:id", controller.UpdateAppointment)
	router.GET("/appointments/byCustomerAndDate", controller.GetAppointmentByCustomerIDAndDate)
}

func RegisterCustomerRoutes(router *gin.Engine, controller *controllers.CustomerController) {
	router.GET("/customer/:id", controller.GetCustomerByID)
	router.GET("/customer/customerID/:customerID", controller.GetCustomerByCustomerID)
	router.GET("/customers", controller.GetAllCustomers)
	router.GET("/customer/email/:email", controller.GetCustomerByEmail)
	router.GET("/customer/searchByID", controller.SearchCustomersByID)
	router.GET("/customer/searchByName", controller.SearchCustomersByName)
	router.GET("/customer/searchByLastName", controller.SearchCustomersByLastName)
	router.POST("/customer", controller.CreateCustomer)
	router.PUT("/customer/:id", controller.UpdateCustomer)
}

func RegisterOrderStateTypeRoutes(router *gin.Engine, controller *controllers.OrderStateTypeController) {
	router.GET("/order-state-type", controller.GetAllOrderStateTypes)
	router.GET("/order-state-type/:id", controller.GetOrderStateTypeByID)
}

func RegisterPurchaseOrderRoutes(router *gin.Engine, controller *controllers.PurchaseOrderController) {
	router.GET("/purchaseorder/:id", controller.GetPurchaseOrderByID)
	router.GET("/purchaseorder", controller.GetAllPurchaseOrders)
	router.GET("/purchaseorder/searchByID", controller.SearchPurchaseOrdersByID)
	router.GET("/purchaseorder/customer/:customerID", controller.GetPurchaseOrdersByCustomerID)
	router.GET("/purchaseorder/seller/:sellerID", controller.GetPurchaseOrdersBySellerID)
	router.POST("/purchaseorder", controller.CreatePurchaseOrder)
	router.PATCH("/purchaseorder/:id/state", controller.UpdatePurchaseOrderState)
	//router.PUT("/purchaseorder/:id", controller.UpdatePurchaseOrder)

}

func RegisterDiscountTypeRoutes(router *gin.Engine, controller *controllers.DiscountTypeController) {
	router.GET("/discount-type", controller.GetAllDiscountTypes)
	router.GET("/discount-type/:id", controller.GetDiscountTypeByID)
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
	router.GET("/invoice/:id", controller.GetInvoiceByID)
	router.GET("/invoice", controller.GetAllInvoices)
	router.GET("/invoice/searchById", controller.SearchInvoiceByID)
	router.GET("/invoice/searchByPersonalId", controller.SearchInvoiceByCustomerPersonalId)
	router.POST("/invoice", controller.CreateInvoice)
}
