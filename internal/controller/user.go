package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "goframe-demo/api/v1"
)

var UserController = userController{}

type userController struct {
}

// Index 用户列表
func (c *cHello) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	db := g.DB()

	all, err := db.Model("users").All()
	if err != nil {
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.Writeln(v1.Res{
		Code:    200,
		Message: "success",
		Data:    all})
	return
}

func (c *cHello) Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error) {
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

// Create 创建用户
func (c userController) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	data := g.Map{"name": req.Name, "age": req.Age, "email": req.Email}
	db := g.DB()
	_, err = db.Model("users").Data(data).Insert()
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.Writeln(v1.Res{
		Code:    200,
		Message: "success",
		Data:    data})
	return
}

func (c *cHello) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	db := g.DB()
	all, err := db.Model("users").Data("name", "嘿嘿嘿").Where("id", 1).Update()
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.Writeln(all)
	return
}

func (c *cHello) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
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
