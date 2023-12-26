package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/price/entities"
	"context"
	"net/http"
	"strings"
)

func (b *createBiz) CreatePriceStorage(ctx context.Context, priceCreate entities.PriceCreate) *entities.PriceGet {
	// declare
	var priceGet *entities.PriceGet
	// priceID
	priceCreate.PriceID = 0
	// ProductID int     `json:"ProductID" gorm:"column:ProductID"`
	if priceCreate.ProductID <= 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: PRICE_ERR_PRODUCTID_POSITIVE,
		})
	}

	// Name      string  `json:"name" gorm:"column:Name"`
	priceCreate.Name = strings.Trim(priceCreate.Name, " ")

	//Price     float64 `json:"price" gorm:"column:Price"`
	if priceCreate.Price <= 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: PRICE_ERR_Price_NOT_NEGATIVE,
		})
	}
	//Desc      string  `json:"desc" gorm:"column:Desc"`
	priceCreate.Desc = strings.Trim(priceCreate.Desc, " ")

	//Image     string  `json:"image" gorm:"column:Image"`
	priceCreate.Image = strings.Trim(priceCreate.Image, " ")

	// creating
	priceGet = b.store.CreatePriceStorage(ctx, priceCreate)
	return priceGet
}
