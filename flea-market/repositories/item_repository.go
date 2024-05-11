package repositories

import (
	"errors"
	"flea-market/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemId uint) error
}

/*
ItemMemoryRepository
*/
type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, v := range r.items {
		if v.ID == itemId {
			return &v, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

func (r *ItemMemoryRepository) Update(updateItem models.Item) (*models.Item, error) {
	targetItem, err := r.FindById(updateItem.ID)
	if err != nil {
		return nil, err
	}
	targetItem = &updateItem
	return targetItem, nil
}

func (r *ItemMemoryRepository) Delete(deleteItemId uint) error {
	for i, v := range r.items {
		if v.ID == deleteItemId {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

/*
ItemRepository
*/
type ItemRepository struct {
	items *gorm.DB
}

func NewItemRepository(items *gorm.DB) IItemRepository {
	return &ItemRepository{items: items}
}

func (r *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.items.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

func (r *ItemRepository) Delete(itemId uint) error {
	item, err := r.FindById(itemId)
	if err != nil {
		return err
	}
	result := r.items.Delete(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := r.items.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

func (r *ItemRepository) FindById(itemId uint) (*models.Item, error) {
	var item models.Item
	result := r.items.First(&item, itemId) //First(dest, conds): のdestに構造体、condsに欲しいデータのprimary keyを指定する
	if result.Error != nil {
		return nil, errors.New("item not found")
	}
	return &item, nil
}

func (r *ItemRepository) Update(updateItem models.Item) (*models.Item, error) {
	result := r.items.Save(updateItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateItem, nil
}
