package service

import (
	"errors"
	"github.com/Walk2future/bi-chatgpt-golang-python/common/requests"
	"github.com/Walk2future/bi-chatgpt-golang-python/models"
	"github.com/duke-git/lancet/v2/strutil"
	"gorm.io/gorm"
)

type UserService struct{}

// UserLogin 用户登录业务
func (userService *UserService) UserLogin(request *requests.UserLoginRequest) (user *models.User, err error) {
	userAccount := request.UserAccount
	userPassword := request.UserPassword
	if strutil.IsBlank(userAccount) {
		return nil, errors.New("用户名为空")
	}
	if strutil.IsBlank(userPassword) {
		return nil, errors.New("密码为空")
	}
	err = models.BI_DB.Where("userAccount = ? AND userPassword = ?", userAccount, userPassword).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名或密码错误")
	}
	return user, err
}
