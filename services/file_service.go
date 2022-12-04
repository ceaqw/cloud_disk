package services

import (
	"CouldDisk/models"
)

type FileService struct {
	fileModel models.FileOrm
}

func (f FileService) GetFileListByType(filetype, page, num, userid int) (int, []interface{}) {
	var results []interface{}
	var total int = 0
	fb := f.fileModel.GetFileListByType(filetype, page, num, userid)
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

func (f FileService) GetFileList(path string, page, pagesize, userid int) (int, []interface{}) {
	var results []interface{}
	var total int = 0
	fb := f.fileModel.GetFileByPath(path, page, pagesize, userid)
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

func (f FileService) GetRecoveryFileList(userid int) []interface{} {
	var results []interface{}
	fb := f.fileModel.GetRecoveryFileList(userid)
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
	return results
}

func (f FileService) GetShareFileList(path string, page, pagesize, userid int) (int, []interface{}) {
	var results []interface{}
	var total int = 0
	fb := f.fileModel.GetShareFileByPath(path, page, pagesize, userid)
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
func (f FileService) GetFileTree(userid int) interface{} {
	results := make(map[int]map[string]interface{})
	fb := f.fileModel.GetAllFile(userid)
	i := 1
	for _, value := range *fb {
		result := make(map[string]interface{})
		result["id"] = value.Id
		result["label"] = value.Name
		result["ext"] = value.Ext
		result["parentpath"] = value.ParentPath
		result["isdir"] = value.IsDir
		results[i] = result
		i += 1
	}
	children := f.getChildren("/", results)
	treeRes := make(map[string]interface{})
	treeRes["id"] = 0
	treeRes["label"] = "/"
	treeRes["children"] = children
	return treeRes

}

func (f FileService) getChildren(path interface{}, results map[int]map[string]interface{}) []interface{} {
	var ress []interface{}
	for _, result := range results {
		if result["parentpath"] == path {
			res := make(map[string]interface{})
			res["id"] = result["id"]
			res["label"] = result["label"]
			if result["isdir"] == uint8(1) {
				res["children"] = f.getChildren(result["label"], results)
			} else {
				res["children"] = nil
			}
			ress = append(ress, res)
		}
	}
	return ress
}
