package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"goframe-demo/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (c *cHello) CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error) {
	db := g.DB()
	insert, err := db.Model("users").Data(g.Map{"name": "goframe"}).Insert()
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.Writeln(insert)
	return
}

func (c *cHello) User(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error) {
	db := g.DB()

	all, err := db.Model("users").All()
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
	g.RequestFromCtx(ctx).Response.Writeln(result)
	return
}
