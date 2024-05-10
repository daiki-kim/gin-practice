package services

import (
	"flea-market/dto"
	"flea-market/models"
	"flea-market/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItemInput dto.CreateItemInput) (*models.Item, error)
	Update(updateItemInput dto.UpdateItemInput, itemId uint) (*models.Item, error)
	Delete(deleteItemId uint) error
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(itemId uint) (*models.Item, error) {
	return s.repository.FindById(itemId)
}

func (s *ItemService) Create(newItemInput dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        newItemInput.Name,
		Price:       newItemInput.Price,
		Description: newItemInput.Description,
		SoldOut:     false,
	}
	return s.repository.Create(newItem)
}

func (s *ItemService) Update(updateItemInput dto.UpdateItemInput, updataItemId uint) (*models.Item, error) {
	targetItem, err := s.FindById(updataItemId)
	if err != nil {
		return nil, err
	}
	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}
	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}
	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}
	if updateItemInput.SoldOut != nil {
		targetItem.SoldOut = *updateItemInput.SoldOut
	}
	return s.repository.Update(*targetItem)
}

func (s *ItemService) Delete(deleteItemId uint) error {
	return s.repository.Delete(deleteItemId)
}
