package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"goframe-demo/internal/service/internal/dao"
	"goframe-demo/internal/service/internal/do"
)

type sUser struct {
}

var insUser = sUser{}

func User() *sUser {
	return &insUser
}

func (u *sUser) GetList(ctx context.Context) (gdb.Result, error) {
	res, err := dao.Users.Ctx(ctx).All()
	if err != nil {
		return nil, err
	}
	return res, err
}

func (u *sUser) CreateUser(ctx context.Context, name string, email string) (err error) {
	_, err = dao.Users.Ctx(ctx).Data(do.Users{Name: name, Email: email}).Insert()
	return err
}
