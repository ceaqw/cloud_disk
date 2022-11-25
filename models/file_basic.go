package models

import (
	"time"

	"github.com/druidcaesa/gotool"
)

type FileBasic struct {
	Id         uint64    `xorm:"int(10) pk autoincr" json:"id"`
	UserId     uint64    `xorm:"int(10)" json:"user_id"`
	Hash       string    `xorm:"varchar(255)" json:"hash"`
	Name       string    `xorm:"varchar(255)" json:"name"`
	Ext        string    `xorm:"varchar(30)" json:"ext"`
	Size       uint64    `xorm:"int(10)" json:"size"`
	FullPath   string    `xorm:"varchar(255)" json:"full_path"`
	ParentPath string    `xorm:"varchar(255)" json:"parent_path"`
	IsDir      uint8     `xorm:"tinyint(1) " json:"is_dir"`
	Created    time.Time `xorm:"created default current_timestamp" json:"created_at"`
	Updated    time.Time `xorm:"created default current_timestamp" json:"updated_at"`
	Deleted    time.Time `xorm:"created " json:"deleted_at"`
	ShareType  uint8     `xorm:"tinyint(1) default 1" json:"share_type"` //1公共 0私密
	ShareTime  time.Time `xorm:"created " json:"share_time"`
}

func (FileBasic) TableName() string {
	return "file_basic"
}

type FileOrm struct {
}

func (f FileOrm) GetFileByPath(path string, page, pageSize, userid int) *[]FileBasic {
	var files []FileBasic
	err := MainDb.Where("parent_path = ?", path).And("user_id = ?", userid).Limit(pageSize, page*pageSize).Find(&files)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
	}
	return &files
}
