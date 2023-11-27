package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/product/entities"
	"context"
	"net/http"
	"strings"
)

func (b *createBiz) UpdateProductBiz(ctx context.Context, productUpdate entities.ProductUpdate, productid int) *entities.ProductGet {
	// productid int
	if productid <= 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: PRODUCT_ERR_PRODUCT_ID_POSITIVE,
		})
	}

	// CateID int
	if productUpdate.CateID < 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: PRODUCT_ERR_CateID_POSITIVE,
		})
	}

	// Name
	productUpdate.Name = strings.Trim(productUpdate.Name, " ")
	// regex
	CheckValidName(productUpdate.Name)

	// Desc
	productUpdate.Desc = strings.Trim(productUpdate.Desc, " ")

	// Image

	// Status
	productUpdate.Status = strings.Trim(productUpdate.Status, " ")
	if productUpdate.Status != "" {
		flag := false
		for _, status := range entities.PRODUCT_STATUS {
			if productUpdate.Status == status {
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

	// updating
	productGet := b.store.UpdateProductStorage(ctx, productUpdate, productid)

	// return
	return productGet
}
