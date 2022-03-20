package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "goframe-demo/api/v1"
	"goframe-demo/internal/service"
)

var User = cUser{}

type cUser struct {
}

// Index 用户列表
func (c *cUser) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	all, err := service.User().GetList(ctx)
	if err != nil {
		return
	}
	res = &v1.IndexRes{
		Result: all,
	}
	return
}

// Create 创建用户
func (c *cUser) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = service.User().CreateUser(ctx, req.Name, req.Email)
	return
}

func (c *cUser) Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error) {
	ctx.Value("id")

	db := g.DB()
	one, err := db.Model("users").Where("id", 1).One()
	if err != nil {
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.Writeln(v1.Res{
		Code:    200,
		Message: "success",
		Data:    one})
	return
}

func (c *cUser) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	db := g.DB()
	all, err := db.Model("users").Data("name", "嘿嘿嘿").Where("id", 1).Update()
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.Writeln(all)
	return
}

func (c *cUser) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	db := g.DB()
	id := req.Id
	result, err := db.Model("users").Delete("id", id)
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.Writeln(v1.Res{
		Code:    200,
		Message: "success",
		Data:    result})
	return
}
