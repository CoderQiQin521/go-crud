package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"第一个api"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type UserReq struct {
	g.Meta `path:"/user" tags:"User" method:"get" summary:"用户列表"`
}

type UserRes struct {
	g.Meta `mime:"application/json" example:"[{"id": 1,"name": "张三","age": 22,"email": "zhang@qq.com"},{"id": 2,"name": "李四","age": 23,"email": "lisi@qq.com"}]"`
}

type CreateUserReq struct {
	g.Meta `path:"/user" tags:"CreateUser" method:"post" summary:"创建用户"`
}

type CreateUserRes struct {
	g.Meta `mime:"application/json" example:"123"`
}

type DeleteReq struct {
	g.Meta `path:"/user" tags:"Delete" method:"delete" summary:"删除用户"`
	Id     int `json:"id"`
}

type DeleteRes struct {
	g.Meta `mime:"application/json" example:"true"`
}
