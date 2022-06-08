package models

import (
	// "time"
	"github.com/jinzhu/gorm"
)

type User struct {
	Model

	// query tag是query参数别名，json xml，form适合post
	Username  *string `validate:"" gorm:"unique;comment:'用户名'" query:"username" json:"username" xml:"username" form:"username"`
	Password  string  `validate:"" gorm:"comment:'密码'" query:"password" json:"password" xml:"password" form:"password"`
	Mobile    *string `validate:"" gorm:"unique;comment:'手机号'" query:"mobile" json:"mobile" xml:"mobile" form:"mobile"`
	AvatarUrl *string `validate:"" gorm:"comment:'头像'" query:"avatarUrl" json:"avatarUrl" xml:"avatarUrl" form:"avatarUrl"`
	NickName  *string `validate:"" gorm:"comment:'昵称'" query:"nickName" json:"nickName" xml:"nickName" form:"nickName"`
	Gender    *string `validate:"" sql:"type:enum('0', '1', '2')" gorm:"comment:'性别,0未知 1男 2女';default:'0'" query:"gender" json:"gender" xml:"gender" form:"gender"`
	Province  *string `validate:"" gorm:"comment:'省'" query:"province" json:"province" xml:"province" form:"province"`
	City      *string `validate:"" gorm:"comment:'市'" query:"city" json:"city" xml:"city" form:"city"`
	Country   *string `validate:"" gorm:"comment:'区'" query:"country" json:"country" xml:"country" form:"country"`
	Unionid   *string `validate:"" gorm:"comment:'小程序unionid'" query:"unionid" json:"unionid" xml:"unionid" form:"unionid"`
	Openid    *string `validate:"" gorm:"comment:'小程序openid'" query:"openid" json:"openid" xml:"openid" form:"openid"`
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

func AddUser(user User) error {
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
