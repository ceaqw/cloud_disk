package controller

import (
	"CouldDisk/models/resp"
	"CouldDisk/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transfer struct {
	TransferService services.TransferService
}

func (t Transfer) GetStorage(c *gin.Context) {
	userid, _ := c.Get("userId")
	userid_int := int(userid.(uint64))
	data := t.TransferService.GetStorageInfosById(userid_int)
	c.JSON(http.StatusOK, resp.Success(data))
}
