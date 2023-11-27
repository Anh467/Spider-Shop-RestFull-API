package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
	"net/http"
	"strings"
)

func (b *createBiz) CreateProductBiz(ctx context.Context, productCreate entities.ProductCreate) *entities.ProductGet {
	// CateID int    `json:"cateid" gorm:"column:CateID"`
	if productCreate.CateID < 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: PRODUCT_ERR_CateID_POSITIVE,
		})
	}
	// Name   string `json:"name" gorm:"column:Name"`
	productCreate.Name = strings.Trim(productCreate.Name, " ")
	// regex
	CheckValidName(productCreate.Name)

	// Desc   string `json:"desc" gorm:"column:Desc"`
	productCreate.Desc = strings.Trim(productCreate.Desc, " ")

	// Image  string `json:"image" gorm:"column:Image"`

	// Status string `json:"status" gorm:"column:Status;default:Normal"`
	productCreate.Status = strings.Trim(productCreate.Status, " ")
	if productCreate.Status != "" {
		flag := false
		for _, status := range entities.PRODUCT_STATUS {
			if productCreate.Status == status {
				flag = true
			}
		}
		if !flag {
			panic(&common.ErrorHandler{
				ErrorMessage: CATE_ERR_Status_EXIST,
				ErrorCode:    http.StatusBadRequest,
			})
		}
	}

	// creating
	productGet := b.store.CreateProductStorage(ctx, productCreate)

	// return
	return productGet
}
