package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-practice-api/dto"
	"gin-practice-api/services"
)

type IItemController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ItemController struct {
	service services.IItemService
}

func NewItemController(service services.IItemService) IItemController {
	return &ItemController{service: service}
}

func (c *ItemController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

func (c *ItemController) FindById(ctx *gin.Context) {
	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	item, err := c.service.FindById(uint(itemId))
	if err != nil {
		if err.Error() == "item not found" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": item})
}

func (c *ItemController) Create(ctx *gin.Context) {
	var newItemInput dto.CreateItemInput
	if err := ctx.ShouldBindJSON(&newItemInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newItem, err := c.service.Create(newItemInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newItem})
}

func (c *ItemController) Update(ctx *gin.Context) {
	updateItemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var updateItemInput dto.UpdateItemInput
	if err := ctx.ShouldBindJSON(&updateItemInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateItem, err := c.service.Update(updateItemInput, uint(updateItemId))
	if err != nil {
		if err.Error() == "item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updateItem})
}

func (c *ItemController) Delete(ctx *gin.Context) {
	deleteItemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := c.service.Delete(uint(deleteItemId)); err != nil {
		if err.Error() == "item not found" {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
		return
	}

	ctx.Status(http.StatusOK)
}
