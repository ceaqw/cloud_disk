package models

import (
	"fmt"
	"time"

	"github.com/druidcaesa/gotool"
)

// DROP TABLE IF EXISTS `user_basic`;
// CREATE TABLE `user_basic` (
//   `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
//   `identity` varchar(36) DEFAULT NULL,
//   `name` varchar(60) DEFAULT NULL,
//   `password` varchar(32) DEFAULT NULL,
//   `email` varchar(100) DEFAULT NULL,
//   `created_at` datetime DEFAULT NULL,
//   `updated_at` datetime DEFAULT NULL,
//   `deleted_at` datetime DEFAULT NULL,
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

type UserBasic struct {
	Id       uint64    `xorm:"pk autoincr" json:"id"`
	Name     string    `xorm:"varchar(255) notnull " json:"name"`
	Password string    `xorm:"varchar(255)" json:"password"`
	Email    string    `xorm:"varchar(255)" json:"email"`
	Identity string    `xorm:"varchar(255)" json:"identity"`
	Created  time.Time `xorm:"created default current_timestamp" json:"created_at"`
	Updated  time.Time `xorm:"created default current_timestamp" json:"updated_at"`
	Deleted  time.Time `xorm:"created " json:"deleted_at"`
}

func (UserBasic) TableName() string {
	return "user_basic"
}

type UserOrm struct {
}

func (u UserOrm) GetUserByEmail(email string) *UserBasic {
	user := UserBasic{}
	ok, err := MainDb.Where("email = ?", email).Get(&user)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
	}
	if ok {
		return &user
	}
	return nil
}
func (u UserOrm) AddUser(user UserBasic) error {
	_, err := MainDb.InsertOne(user)
	return err
}

func (u UserOrm) UpdatePwdByEmail(email, pwd string) bool {

	affected, err := MainDb.Exec("update user_basic set password = ? where email = ?", pwd, email)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return false
	}
	fmt.Println(affected)
	return true
}
