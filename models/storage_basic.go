package models

import (
	"time"

	"github.com/druidcaesa/gotool"
)

type StorageBasic struct {
	UserId       uint64    `xorm:"int(10) pk" json:"user_id"`
	StorageSpace uint64    `xorm:"" json:"storage_space"`
	UseSpace     uint64    `xorm:"" json:"use_space"`
	Updated      time.Time `xorm:"created default current_timestamp" json:"updated_at"`
}

func (StorageBasic) TableName() string {
	return "storage_basic"
}

type StorageOrm struct {
}

func (s StorageOrm) GetStorageInfosById(id int) *StorageBasic {
	userStorage := StorageBasic{}
	ok, err := MainDb.Where("user_id = ?", id).Get(&userStorage)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
	}
	if ok {
		return &userStorage
	}
	return nil
}
