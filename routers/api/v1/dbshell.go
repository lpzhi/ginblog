package v1

import (
	"ginblog/models"
	"ginblog/pkg/e"
	"github.com/gin-gonic/gin"
)

func CreateRoleTotal(c *gin.Context)  {
	models.CreateRoleTotalDatabase()
//	maps := make(map[string]interface{})
	total := models.GetTagTotal();
	c.JSON(e.SUCCESS, gin.H{
		"total":total,
	})
}

func CreateEenterLog(c *gin.Context)  {
}
