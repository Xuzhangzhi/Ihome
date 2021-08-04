package controllers
import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"Ihome/models"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) RetData(resp map[string]interface{}) {
	c.Data["json"] = &resp
	c.ServeJSON()
}

func (c *UserController) Reg() {
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &resp)
	fmt.Println("resp_mobile:",resp["mobile"])
	fmt.Println("resp_pwd:",resp["password"])
	fmt.Println("resp_code:",resp["sms_code"])
	user := models.User{}
	o := orm.NewOrm()
	user.Name = resp["mobile"].(string)
	user.Mobile = resp["mobile"].(string)
	user.Password_hash = resp["password"].(string)
	id, err := o.Insert(&user)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	fmt.Println("reg success, id=", id)
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	c.SetSession("name", user.Name)
}