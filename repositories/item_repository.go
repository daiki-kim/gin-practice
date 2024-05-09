package repositories

import (
	"errors"
	"gin-practice-api/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemId uint) error
}

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
