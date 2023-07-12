package service

import (
	"errors"
	"github.com/Walk2future/bi-chatgpt-golang-python/common/requests"
	"github.com/Walk2future/bi-chatgpt-golang-python/models"
	"github.com/Walk2future/bi-chatgpt-golang-python/models/serializers"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct{}

// Login 用户登录业务
func (userService *UserService) Login(request *requests.LoginRequest) (NewUser *serializers.CurrentUser, err error) {
	userAccount := request.UserAccount
	userPassword := request.UserPassword
	if strutil.IsBlank(userAccount) {
		return nil, errors.New("用户名为空")
	}
	if strutil.IsBlank(userPassword) {
		return nil, errors.New("密码为空")
	}
	var user *models.User
	err = models.BI_DB.Where("userAccount = ? AND userPassword = ?", userAccount, userPassword).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名或密码错误")
	}
	return &serializers.CurrentUser{
		ID:          user.ID,
		UserAccount: user.UserAccount,
		UserName:    user.UserName,
		UserAvatar:  user.UserAvatar,
		UserRole:    user.UserRole,
	}, err
}

// Register 用户注册业务
func (userService *UserService) Register(request *requests.RegisterRequest) (res interface{}, err error) {
	userAccount := request.UserAccount
	userPassword := request.UserPassword
	checkPassword := request.CheckPassword
	if strutil.IsBlank(userAccount) {
		return nil, errors.New("用户名为空")
	}
	if strutil.IsBlank(userPassword) {
		return nil, errors.New("密码为空")
	}
	if strutil.IsBlank(checkPassword) {
		return nil, errors.New("检查密码为空")
	}
	if len(userAccount) < 6 || len(userAccount) > 18 {
		return nil, errors.New("用户名长度不能低于6位并且不超过18位")
	}
	if len(userPassword) < 8 || len(userPassword) > 18 {
		return nil, errors.New("密码长度不能低于8位并且不能超过18位")
	}
	if userPassword != checkPassword {
		return nil, errors.New("两次输入密码不一致")
	}
	// TODO 检查非法字符
	var count int64
	models.BI_DB.Model(&models.User{}).Where("userAccount = ?", userAccount).Count(&count)
	if count != 0 {
		return nil, errors.New("用户名已存在")
	}
	user := &models.User{
		UserAccount:  userAccount,
		UserPassword: userPassword,
	}
	err = models.BI_DB.Model(&models.User{}).Select("userAccount", "userPassword").Create(&user).Error
	return user.UserAccount, err
}

// Current 获取当前用户业务
func (userService *UserService) Current(c *gin.Context) (*serializers.CurrentUser, error) {
	value, exists := c.Get("id")
	if !exists {
		return nil, errors.New("ID not found")
	}

	return value.(*serializers.CurrentUser), nil
}

// Add 新增用户操作
func (userService *UserService) Add(requests.AddUserRequest) (id int, err error) {
	var user *models.User
	err = models.BI_DB.Model(&models.User{}).Select("userAccount", "userPassword", "userName",
		"userAvatar", "userRole", "freeCount").Create(&user).Error
	if err != nil {
		return -1, errors.New("插入数据库失败")
	}
	return user.ID, nil
}

// RemoveById 根据主键删除用户
func (userService *UserService) RemoveById(id string) (bool, error) {
	var user *models.User
	tx := models.BI_DB.Model(&user).Delete(&user, id)
	if tx.Error != nil {
		return false, errors.New("根据id删除用户失败")
	}
	return true, nil
}

// BatchRemove 根据id切片批量删除用户
func (userService *UserService) BatchRemove(ids []string) (bool, error) {
	var user models.User
	if err := models.BI_DB.Model(&user).Where("id IN (?)", ids).Delete(&user).Error; err != nil {
		return false, errors.New("批量删除用户失败")
	}
	return true, nil
}

// Update 修改用户信息
func (userService *UserService) Update(newUser *models.User) (bool, error) {
	var user *models.User
	if tx := models.BI_DB.Model(&user).Updates(newUser); tx.Error != nil {
		return false, errors.New("更新用户信息失败")
	}
	return true, nil
}

// List 查询用户信息
func List(page *requests.Page) ([]serializers.CurrentUser, error) {
	var userList []models.User
	if tx := models.BI_DB.Model(&models.User{}).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&userList); tx.Error != nil {
		return nil, errors.New("查询用户信息失败")
	}
	var list []serializers.CurrentUser
	for _, user := range userList {
		currentUser := serializers.CurrentUser{
			ID:          user.ID,
			UserAccount: user.UserAccount,
			UserName:    user.UserName,
			UserAvatar:  user.UserAvatar,
			UserRole:    user.UserRole,
		}
		list = append(list, currentUser)
	}
	return list, nil
}
