package service

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/database/gdb"
	"goframe-demo/internal/model"
	"goframe-demo/internal/service/internal/dao"
	"goframe-demo/internal/service/internal/do"
)

var insUser = sUser{}

type sUser struct {
}

func User() *sUser {
	return &insUser
}

// 用户列表
func (u *sUser) GetList(ctx context.Context) (gdb.Result, error) {
	res, err := dao.Users.Ctx(ctx).All()
	if err != nil {
		return nil, err
	}
	return res, err
}

// 添加用户
func (u *sUser) CreateUser(ctx context.Context, in model.CreateUserInput) (err error) {
	_, err = dao.Users.Ctx(ctx).Data(do.Users{
		Name:  in.Name,
		Email: in.Email,
	}).Insert()
	return err
}

// 用户删除
func (u *sUser) Delete(ctx context.Context, id int) (result sql.Result, err error) {
	result, err = dao.Users.Ctx(ctx).Delete("id", id)
	if err != nil {
		return nil, err
	}
	return
}
