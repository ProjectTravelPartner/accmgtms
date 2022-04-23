package api

import (
	"strconv"

	"github.com/ProjectTravelPartner/accmgtms/dao"
	"github.com/ProjectTravelPartner/accmgtms/models"
	"github.com/ProjectTravelPartner/commonclient"
	"github.com/gin-gonic/gin"
)

func AccountCreate(c *gin.Context) {
	var data models.Account
	c.BindJSON(&data)
	id, _ := dao.AccountCreate(data)
	commonclient.RespOK(c, id)
}

func AccountGet(c *gin.Context) {
	var out models.Account
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	out, _ = dao.AccountGet(id)
	commonclient.RespOK(c, out)
}
