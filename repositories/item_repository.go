package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (r *ItemRepository) GetItemByID(id string) (*models.Item, error) {
	var item models.Item
	err := r.DB.Preload("ItemType").Preload("AdditionalExpenses").First(&item, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) HasEnoughStock(id string, quantity int) (bool, error) {
	var stock int
	err := r.DB.Model(&models.Item{}).Select("stock").Where("id = ?", id).Scan(&stock).Error
	if err != nil {
		return false, err
	}
	return stock >= quantity, nil
}

func (r *ItemRepository) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	err := r.DB.Preload("ItemType").Preload("AdditionalExpenses").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ItemRepository) SearchItemsByID(query string) ([]models.Item, error) {
	var items []models.Item
	err := r.DB.Preload("ItemType").Preload("AdditionalExpenses").
		Where("CAST(id AS TEXT) LIKE ?", query+"%").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ItemRepository) SearchItemsByName(query string) ([]models.Item, error) {
	var items []models.Item
	err := r.DB.Preload("ItemType").Preload("AdditionalExpenses").
		Where("LOWER(name) LIKE LOWER(?)", query+"%").
		Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ItemRepository) UpdateItemState(id string, state bool) (*models.Item, error) {
	var item models.Item
	if err := r.DB.Preload("ItemType").Preload("AdditionalExpenses").First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}

	item.ItemState = state

	if err := r.DB.Save(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) UpdateItem(item *models.Item) (bool, error) {

	var existingItem models.Item
	if err := r.DB.Preload("ItemType").Preload("AdditionalExpenses").First(&existingItem, "id = ?", item.ID).Error; err != nil {
		return false, err
	}

	priceChanged := existingItem.SellingPrice != item.SellingPrice

	existingItem.ItemTypeID = item.ItemTypeID

	if err := r.DB.Model(&existingItem).Updates(item).Error; err != nil {
		return false, err
	}

	if err := r.DB.Model(&existingItem).Select("ItemState").Updates(item).Error; err != nil {
		return false, err
	}

	return priceChanged, nil
}

func (r *ItemRepository) CreateItem(item *models.Item) (*models.Item, error) {

	if err := r.DB.Create(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *ItemRepository) SubtractItemsFromInventory(itemID string, amount int) error {
	if err := r.DB.Model(&models.Item{}).
		Where("id = ?", itemID).
		UpdateColumn("stock", gorm.Expr("stock - ?", amount)).Error; err != nil {
		return err
	}
	return nil
}

func (r *ItemRepository) ReturnItemsToInventory(itemID string, amount int) error {
	if err := r.DB.Model(&models.Item{}).
		Where("id = ?", itemID).
		UpdateColumn("stock", gorm.Expr("stock + ?", amount)).Error; err != nil {
		return err
	}
	return nil
}
