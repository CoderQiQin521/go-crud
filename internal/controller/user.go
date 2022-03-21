package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "goframe-demo/api/v1"
	"goframe-demo/internal/model"
	"goframe-demo/internal/service"
)

var User = cUser{}

type cUser struct {
}

// Index 用户列表
func (c *cUser) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	all, err := service.User().GetList(ctx)
	if err != nil {
		return nil, err
	}
	//res = &all
	res = &v1.IndexRes{
		Result: all,
	}
	return
}

// Create 创建用户
func (c *cUser) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	data := model.CreateUserInput{
		Name:  req.Name,
		Email: req.Email,
	}
	err = service.User().CreateUser(ctx, data)
	return
}

func (c *cUser) Show(ctx context.Context, req *v1.ShowReq) (res *v1.ShowRes, err error) {
	//id := r.Get("id")
	//db := g.DB()
	//one, err := db.Model("users").Where("id", id).One()
	//if err != nil {
	//	return nil, err
	//}
	//
	//g.RequestFromCtx(ctx).Response.Writeln(v1.Res{
	//	Code:    200,
	//	Message: "success",
	//	Data:    one})
	g.RequestFromCtx(ctx).Response.Writeln(ctx.Value("id"))
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

// 用户删除
func (c *cUser) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	id := req.Id

	result, err := service.User().Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	res.Code = 0
	res.Message = "删除成功"
	res.Data = result
	return
}
