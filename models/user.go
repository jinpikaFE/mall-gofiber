package models

import (
	// "time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Model

	// query tag是query参数别名，json xml，form适合post
	Username  string `validate:"" gorm:"unique" validate:"unique" query:"username" json:"username" xml:"username" form:"username"`
	Password  string `validate:"" query:"password" json:"password" xml:"password" form:"password"`
	Mobile    string `validate:"" gorm:"unique" validate:"unique" query:"mobile" json:"mobile" xml:"mobile" form:"mobile"`
	AvatarUrl string `validate:"" query:"avatarUrl" json:"avatarUrl" xml:"avatarUrl" form:"avatarUrl"`
	NickName  string `validate:"" query:"nickName" json:"nickName" xml:"nickName" form:"nickName"`
	Gender    int    `validate:"" query:"gender" json:"gender" xml:"gender" form:"gender"`
	Province  string `validate:"" query:"province" json:"province" xml:"province" form:"province"`
	City      string `validate:"" query:"city" json:"city" xml:"city" form:"city"`
	Country   string `validate:"" query:"country" json:"country" xml:"country" form:"country"`
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetUserTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&User{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*User, error) {
	var users []*User
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

func GetUser(maps interface{}) (*User, error) {
	var user User
	err := db.Where(maps).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func AddUser(user *User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func EditUser(id int, data interface{}) error {
	if err := db.Model(&User{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	if err := db.Where("id = ?", id).Delete(User{}).Error; err != nil {
		return err
	}

	return nil
}

// 根据id判断test 对象是否存在
func ExistUserByID(id int) bool {
	var user User
	db.Select("id").Where("id = ?", id).First(&user)

	return user.ID > 0
}

// // gorm所支持的回调方法：

// // 创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
// // 更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
// // 删除：BeforeDelete、AfterDelete
// // 查询：AfterFind

// func (user *User) BeforeCreate(scope *gorm.Scope) error {
// 	scope.SetColumn("CreatedOn", time.Now().Unix())

// 	return nil
// }

// func (user *User) BeforeUpdate(scope *gorm.Scope) error {
// 	scope.SetColumn("ModifiedOn", time.Now().Unix())

// 	return nil
// }
