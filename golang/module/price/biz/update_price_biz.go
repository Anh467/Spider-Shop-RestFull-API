package biz

import (
	"SpiderShop-Restfull-API/common"
	"SpiderShop-Restfull-API/module/price/entities"
	"context"
	"net/http"
	"strings"
)

func (b *createBiz) UpdatePriceBiz(ctx context.Context, priceUpdate entities.PriceUpdate, priceid int) {
	// decalare
	var flag = false

	// check the postive of the priceid
	if priceid <= 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: PRICE_ERR_PriceID_POSITIVE,
		})
	}

	// Name      string  `json:"name" gorm:"column:Name"`
	priceUpdate.Name = strings.Trim(priceUpdate.Name, " ")

	//Price     float64 `json:"price" gorm:"column:Price"`
	if priceUpdate.Price <= 0 {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: PRICE_ERR_Price_NOT_NEGATIVE,
		})
	}
	//Desc      string  `json:"desc" gorm:"column:Desc"`
	priceUpdate.Desc = strings.Trim(priceUpdate.Desc, " ")

	//Image     string  `json:"image" gorm:"column:Image"`
	priceUpdate.Image = strings.Trim(priceUpdate.Image, " ")

	// checking status
	for _, v := range entities.Price_STATUS {
		if v == priceUpdate.Status {
			flag = true
		}
	}
	if !flag {
		panic(&common.ErrorHandler{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: PRICE_ERR_Status_NOT_APPROPIRATE,
		})
	}

	// updating
	b.store.UpdatePriceStorage(ctx, priceUpdate, priceid)

}
