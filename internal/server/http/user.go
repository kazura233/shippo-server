package http

import (
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type UserServer struct {
	*Server
	router *box.RouterGroup
}

func NewUserServer(server *Server) *UserServer {
	var s = &UserServer{
		Server: server,
		router: server.router.Group("user"),
	}
	s.initRouter()
	return s
}

func (t *UserServer) initRouter() {
	t.router.POST("login", t.UserLogin)
	t.router.POST("findAll", t.FindAll)
	t.router.POST("updateUserRole", t.UpdateUserRole)
}

func (t *UserServer) UserLogin(c *box.Context) {
	var param model.UserLoginParam
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	data, err := t.service.User.UserLogin(param, c.Req.Passport)
	c.JSON(data, err)
}

func (t *UserServer) FindAll(c *box.Context) {
	var param model.UserFindAllReq
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	data, err := t.service.User.FindAll(param)
	c.JSON(data, err)
}

func (t *UserServer) UpdateUserRole(c *box.Context) {
	var param model.User
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	err := t.service.User.UpdateUserRole(param)
	c.JSON(nil, err)
}
