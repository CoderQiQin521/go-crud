package service

type sUser struct {
}

var insUser = sUser{}

func User() *sUser {
	return &insUser
}

//func (u *sUser) GetList(ctx context.Context) {
//	var m = dao.Users.Ctx(ctx)
//
//	m = m.Where(dao.Users.Columns().Name)
//	return m
//}
