package storage

import (
	"SpiderShop-Restfull-API/module/user/entities"
	"context"
)

func (s *mySQLStore) GetUserRoleStorage(c context.Context, account string) string {
	// declare
	var role string
	// get role column
	if err := s.aptx.GormDB.Table(entities.USER_TABLE).Where("Account = ?", account).Pluck("Role", &role).Error; err != nil {
		panic(err)
	}
	// return
	return role
}
