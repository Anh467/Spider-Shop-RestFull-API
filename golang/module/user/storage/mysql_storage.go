package storage

import (
	"SpiderShop-Restfull-API/common"
)

type sqlserverStore struct {
	aptx *common.AppConext
}

func NewMySQLStorage(aptx *common.AppConext) *sqlserverStore {
	return &sqlserverStore{aptx: aptx}
}
