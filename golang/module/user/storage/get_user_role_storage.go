package storage

import "context"

func (s *sqlserverStore) GetUserRoleStorage(c context.Context, account string) string {
	// declare
	var role string
	// get role column
	if err := s.aptx.GormDB.Where("Account = ?", account).Pluck("Role", role); err != nil {
		return ""
	}
	// return
	return role
}
