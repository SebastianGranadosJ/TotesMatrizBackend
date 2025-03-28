package database

import (
	"errors"
	"log"
	"os"
	"totesbackend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDB devuelve la instancia de la base de datos
func GetDB() *gorm.DB {
	return db
}

// StartPostgres inicia la conexión con PostgreSQL
func StartPostgres() error {
	// Obtener la URI desde la variable de entorno
	dsn := os.Getenv("POSTGRES_URI")
	if dsn == "" {
		return errors.New("you must set your 'POSTGRES_URI' environmental variable")
	}

	// Conectar con PostgreSQL usando GORM
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("failed to connect to PostgreSQL")
	}

	// Verificar la conexión
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		return errors.New("can't verify a connection")
	}

	return nil
}

// ClosePostgres cierra la conexión con la base de datos
func ClosePostgres() {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}

}

func MigrateDB() {

	err := db.AutoMigrate(&models.Item{}, &models.ItemType{},
		&models.AdditionalExpense{}, &models.Permission{}, &models.Role{},
		&models.UserType{}, &models.IdentifierType{}, &models.UserStateType{}, &models.Employee{}, &models.HistoricalItemPrice{},
		&models.Comment{}, models.User{}, models.UserLog{}, &models.Customer{}, &models.Appointment{}, models.OrderStateType{}, &models.PurchaseOrder{},
		&models.DiscountType{}, &models.TaxType{}, &models.Invoice{}, &models.InvoiceItem{}, &models.PurchaseOrderItem{})
	if err != nil {
		log.Fatal("Error en la migración de la base de datos:", err)
	}
}
