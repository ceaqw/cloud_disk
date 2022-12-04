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
	pagesize_int, _ := strconv.Atoi(pageSize)
	userid, _ := c.Get("userId")
	userid_int := int(userid.(uint64))
	page_int = page_int - 1
	total, data := f.fileService.GetFileList(filePath, page_int, pagesize_int, userid_int)
	result := make(map[string]interface{})
	result["total"] = total
	result["list"] = data
	c.JSON(http.StatusOK, resp.Success(result))
}

func (f File) GetFileListByType(c *gin.Context) {
	fileType := c.Query("fileType")
	currentPage := c.Query("currentPage")
	pageCount := c.Query("pageCount")
	filetype_int, _ := strconv.Atoi(fileType)
	page_int, _ := strconv.Atoi(currentPage)
	pageSize_int, _ := strconv.Atoi(pageCount)
	userid, _ := c.Get("userId")
	userid_int := int(userid.(uint64))
	page_int = page_int - 1
	total, data := f.fileService.GetFileListByType(filetype_int, page_int, pageSize_int, userid_int)
	result := make(map[string]interface{})
	result["total"] = total
	result["list"] = data
	c.JSON(http.StatusOK, resp.Success(result))
}

func (f File) GetRecoveryFileList(c *gin.Context) {
	userid, _ := c.Get("userId")
	userid_int := int(userid.(uint64))
	data := f.fileService.GetRecoveryFileList(userid_int)
	result := make(map[string]interface{})
	result["list"] = data
	c.JSON(http.StatusOK, resp.Success(result))
}

func (f File) GetShareFileList(c *gin.Context) {
	filePath := c.Query("shareFilePath")
	currentPage := c.Query("currentPage")
	pageCount := c.Query("pageCount")
	page_int, _ := strconv.Atoi(currentPage)
	pagesize_int, _ := strconv.Atoi(pageCount)
	userid, _ := c.Get("userId")
	userid_int := int(userid.(uint64))
	page_int = page_int - 1
	total, data := f.fileService.GetShareFileList(filePath, page_int, pagesize_int, userid_int)
	result := make(map[string]interface{})
	result["total"] = total
	result["list"] = data
	c.JSON(http.StatusOK, resp.Success(result))
}

func (f File) GetFileTree(c *gin.Context) {
	userid, _ := c.Get("userId")
	userid_int := int(userid.(uint64))
	m := f.fileService.GetFileTree(userid_int)
	c.JSON(http.StatusOK, resp.Success(m))
}
