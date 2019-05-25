package v1

import (
	"ginblog/models"
	"ginblog/pkg/e"
	"ginblog/pkg/setting"
	"ginblog/pkg/util"
	"github.com/Unknwon/com"
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
}

func EditTag(c *gin.Context)  {
}

func DeleteTag(c *gin.Context)  {
}
