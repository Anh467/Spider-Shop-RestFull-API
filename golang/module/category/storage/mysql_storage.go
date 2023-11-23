package storage

import "SpiderShop-Restfull-API/common"

type mySQLStore struct {
	aptx *common.AppConext
}

func NewMySQLStorage(aptx *common.AppConext) *mySQLStore {
	return &mySQLStore{
		aptx: aptx,
	}
}
