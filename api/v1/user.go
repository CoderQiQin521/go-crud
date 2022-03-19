package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type IndexReq struct {
	g.Meta `path:"/users" tags:"Index" method:"get" summary:"用户列表"`
}

type IndexRes struct {
	Res
	g.Meta `mime:"application/json" example:"[{"id": 1,"name": "张三","age": 22,"email": "zhang@qq.com"},{"id": 2,"name": "李四","age": 23,"email": "lisi@qq.com"}]"`
}

type ShowReq struct {
	g.Meta `path:"/users/:id" tags:"Show" method:"get" summary:"用户详情"`
}

type ShowRes struct {
	Res
	g.Meta `mime:"application/json" example:"[{"id": 1,"name": "张三","age": 22,"email": "zhang@qq.com"},{"id": 2,"name": "李四","age": 23,"email": "lisi@qq.com"}]"`
}

type UpdateReq struct {
	g.Meta `path:"/users/:id" tags:"Update" method:"put" summary:"用户更新"`
}

type UpdateRes struct {
	Res
	g.Meta `mime:"application/json" example:"[{"id": 1,"name": "张三","age": 22,"email": "zhang@qq.com"},{"id": 2,"name": "李四","age": 23,"email": "lisi@qq.com"}]"`
}

type CreateReq struct {
	g.Meta `path:"/users" tags:"Create" method:"post" summary:"创建用户"`
	Name   string
	Age    int
	Email  string
}

type CreateRes struct {
	Res
	g.Meta `mime:"application/json" example:"成功"`
}

type DeleteReq struct {
	g.Meta `path:"/users" tags:"Delete" method:"delete" summary:"删除用户"`
	Id     int `json:"id"`
}

type DeleteRes struct {
	Res
	g.Meta `mime:"application/json" example:"true"`
}
