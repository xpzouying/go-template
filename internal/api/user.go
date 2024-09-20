package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xpzouying/go-cmd-project-template/internal/controller"
)

type (
	ReqGetUser struct {
		Uid int `json:"uid"`
	}

	RespGetUser struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}
)

func ParseReqGetUser(c *gin.Context) (*ReqGetUser, error) {
	var req ReqGetUser
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func HandleGetUser(c *gin.Context) {

	req, err := ParseReqGetUser(c)
	if err != nil {
		WriteErrorResponseByErrorRequest(c)
		return
	}

	user, err := controller.GetUser(c.Request.Context(), req.Uid)
	if err != nil {
		WriteErrorResponse(c, "获取用户失败")
		return
	}

	WriteSuccessResponse(c, &RespGetUser{
		Uid:  user.Uid,
		Name: user.Name,
	})
}

type (
	ReqCreateUser struct {
		Name string `json:"name"`
	}

	RespCreateUser struct {
		Uid  int    `json:"uid"`
		Name string `json:"name"`
	}
)

func ParseReqCreateUser(c *gin.Context) (*ReqCreateUser, error) {
	var req ReqCreateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func HandleCreateUser(c *gin.Context) {

	req, err := ParseReqCreateUser(c)
	if err != nil {
		WriteErrorResponseByErrorRequest(c)
		return
	}

	user, err := controller.CreateUser(c.Request.Context(), req.Name)
	if err != nil {
		WriteErrorResponse(c, "创建用户失败")
		return
	}

	WriteSuccessResponse(c, &RespCreateUser{
		Uid:  user.Uid,
		Name: user.Name,
	})
}
