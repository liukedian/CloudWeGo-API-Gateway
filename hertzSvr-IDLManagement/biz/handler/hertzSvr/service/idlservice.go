// Code generated by hertz generator.

package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	service "hertzSvr-IDLManagement/biz/model/hertzSvr/service"
)

// AddIDL .
// @router /idlManager/add [POST]
func AddIDL(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.IDLInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(service.IDLResponse)

	var svrInfo IDLMessage
	fmt.Println(svrInfo.ID)
	DB.First(&svrInfo, "svr_name = ?", req.Name)
	if svrInfo.ID != 0 {
		resp.Success = false
		resp.Message = "ERROR:SERVICE ALREADY EXISTED"
		c.JSON(consts.StatusOK, resp)
		return
	}

	DB.Create(&IDLMessage{
		SvrName: req.Name,
		IDL:     req.Idl,
	})
	resp.Success = true
	resp.Message = ""
	c.JSON(consts.StatusOK, resp)
}

// DeleteIDL .
// @router /idlManager/delete [POST]
func DeleteIDL(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.IDLMessage
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(service.IDLResponse)
	var svrInfo IDLMessage
	DB.First(&svrInfo, "svr_name = ?", req.Name)
	if svrInfo.ID == 0 {
		resp.Success = false
		resp.Message = "ERROR:NO SUCH SERVICE"
		c.JSON(consts.StatusOK, resp)
		return
	}

	DB.Delete(&svrInfo, svrInfo.ID)
	resp.Success = true
	resp.Message = ""
	c.JSON(consts.StatusOK, resp)
}

// UpdateIDL .
// @router /idlManager/update [POST]
func UpdateIDL(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.IDLInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(service.IDLResponse)

	var svrInfo IDLMessage
	DB.First(&svrInfo, "svr_name = ?", req.Name)
	if svrInfo.ID == 0 {
		resp.Success = false
		resp.Message = "ERROR:NO SUCH SERVICE"
		c.JSON(consts.StatusOK, resp)
		return
	}
	DB.Model(&svrInfo).Update("id_l", req.Idl)

	resp.Success = true

	c.JSON(consts.StatusOK, resp)
}

// QueryIDL .
// @router /idlManager/query [GET]
func QueryIDL(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.IDLQueryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(service.IDLInfo)

	var svrInfo IDLMessage
	DB.First(&svrInfo, "svr_name = ?", req.Name)
	if svrInfo.ID == 0 {
		resp.Idl = "ERROR:NO SUCH SERVICE"
		resp.Name = req.Name
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp = &service.IDLInfo{
		Name: req.Name,
		Idl:  svrInfo.IDL,
	}

	c.JSON(consts.StatusOK, resp)
}
