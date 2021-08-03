package controllers

import (
	"Ihome/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

type AreaController struct {
	beego.Controller
}

func (c *AreaController) RetData(resp map[string]interface{}) {
	c.Data["json"] = &resp
	c.ServeJSON()
}
func (c *AreaController) GetArea() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	var areas []models.Area

	//connect to mysql
	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)
	if err != nil {
		resp["errno"] = 4001
		resp["errmsg"] = "query error"
		return
	}
	if num == 0 {
		resp["errno"] = 4002
		resp["errmsg"] = "there is no data"
		return
	}
	resp["errno"] = 0
	resp["errmsg"] = "OK"
	resp["data"] = areas
	fmt.Println("query data success, resp = ", resp)
}
