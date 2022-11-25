package services

import (
	"CouldDisk/models"
)

type FileService struct {
	fileModel models.FileOrm
}

func (f FileService) GetFileList(path string, page, pageSize, userid int) (int, []interface{}) {
	var results []interface{}
	var total int = 0
	fb := f.fileModel.GetFileByPath(path, page, pageSize, userid)
	for _, value := range *fb {
		result := make(map[string]interface{})
		if value.IsDir == 1 {
			result["isDir"] = true
		} else {
			result["isDir"] = false
		}
		result["fileName"] = value.Name
		result["filePath"] = value.FullPath
		result["fileSize"] = value.Size
		result["extendName"] = value.Ext
		result["shareType"] = value.ShareType
		result["uploadTime"] = value.Created
		result["deleteTime"] = value.Deleted
		result["shareTime"] = value.ShareTime
		result["endTime"] = value.ShareTime
		results = append(results, result)
	}
	total = len(results)
	return total, results
}
