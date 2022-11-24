package controller

import (
	"CouldDisk/models/resp"
	"CouldDisk/services"
	"fmt"
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
	// value, _ := c.Get("userId")
	total, data := f.fileService.GetFileList(filePath, page_int, pageSize_int)
	result := make(map[string]interface{})
	result["total"] = total
	result["list"] = data
	fmt.Println(result)
	c.JSON(http.StatusOK, resp.Success(result))
}
