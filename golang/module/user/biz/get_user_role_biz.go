package biz

import "golang.org/x/net/context"

func (b *createBiz) GetUserRoleBiz(c context.Context, account string) string {
	// declare
	var role string
	// get role column
	role = b.store.GetUserRoleStorage(c, account)
	return role
}
