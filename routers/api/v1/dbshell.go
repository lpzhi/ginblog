package v1

import (
	"ginblog/models"
	"ginblog/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func CreateRoleTotal(c *gin.Context)  {
	var (
		code int
		err error
		)

	err = models.CreateRoleTotalDatabase();

	if err!=nil{
		code = e.ERROR
	}else {
		code = e.SUCCESS
	}

//	maps := make(map[string]interface{})
//	total := models.GetTagTotal();

	c.JSON(code, gin.H{
		"message":err ,
	})
}

func CreateEnterLog(c *gin.Context)  {
	var (
		msg interface{}
		code int
	)

	pf := c.Query("pf")
	valid := validation.Validation{}
	valid.Required(pf,"pf").Message("平台名称不能为空")

	if err := models.CreateEnterLogDatabaseTables(pf);err !=nil{
		code = e.ERROR
		msg = err
	}

	if valid.HasErrors() {
		msg = valid.Errors
		code = e.UNPROCESSABLE
	}

	c.JSON(code,gin.H{
		"msg":msg,
	})
}

func CreateEenterLog(c *gin.Context)  {
}
