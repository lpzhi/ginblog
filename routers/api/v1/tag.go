package v1

import (
	"ginblog/models"
	"ginblog/pkg/e"
	"ginblog/pkg/setting"
	"ginblog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTags(c *gin.Context)  {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name !="" {
		maps["name"] = name
	}

	if arg := c.Query("state");arg!="" {
		maps["state"] = com.StrTo(arg).MustInt()
	}

	data["lists"] = models.GetTags(	util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal()

	c.JSON(http.StatusOK,gin.H{
		"code":e.SUCCESS,
		"data":data,
	})
}


func AddTag(c *gin.Context)  {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	creatdBy := c.Query("creatd_by")

	valid := validation.Validation{}

	valid.Required(name,"name").Message("名字不能为空")
	valid.MaxSize(name,100,"name").Message("长度不能操作100个字符")
	valid.Required(creatdBy,"creatd_by").Message("创建人不能为空")
	valid.Range(state,0,1,"state").Message("state 只能是 0或者1")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		if !models.ExitsByName(name) {
			code = e.SUCCESS
			models.AddTag(name,creatdBy,state)
		}else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":"",
		"data":make(map[string]string),
	})
}

func EditTag(c *gin.Context)  {
}

func DeleteTag(c *gin.Context)  {
}


