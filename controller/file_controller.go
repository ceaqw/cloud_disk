package controller

import (
	"CouldDisk/models/resp"
	"CouldDisk/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type File struct {
	fileService services.FileService
}

func (f File) GetFileList(c *gin.Context) {
	filePath := c.Query("filePath")
	page := c.Query("currentPage")
	pageSize := c.Query("pageCount")
	page_int, _ := strconv.Atoi(page)
	pageSize_int, _ := strconv.Atoi(pageSize)
	userid, _ := c.Get("userId")
	userid_int := int(userid.(uint64))
	page_int = page_int - 1
	total, data := f.fileService.GetFileList(filePath, page_int, pageSize_int, userid_int)
	result := make(map[string]interface{})
	result["total"] = total
	result["list"] = data
	c.JSON(http.StatusOK, resp.Success(result))
}
