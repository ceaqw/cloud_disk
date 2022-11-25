package services

import "CouldDisk/models"

type TransferService struct {
	StorageModel models.StorageOrm
}

func (t TransferService) GetStorageInfosById(id int) interface{} {
	result := make(map[string]interface{})
	storage := t.StorageModel.GetStorageInfosById(id)
	if storage == nil {
		result["storageSize"] = 0
		result["totalStorageSize"] = 0
	} else {
		result["storageSize"] = storage.StorageSpace
		result["totalStorageSize"] = storage.UseSpace
	}
	return result
}
