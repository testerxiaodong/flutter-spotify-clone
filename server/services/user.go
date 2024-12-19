package services

import (
	"errors"
	"server/common/xerror"
	"server/models/dao"
	"server/models/entity"
	"server/types"
	"server/utils"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserService struct {
	Query *dao.Query
}

func NewUserService(query *dao.Query) *UserService {
	return &UserService{
		Query: query,
	}
}

func (u *UserService) Signup(signupReq *types.SignupReq) (*types.SignupRes, error) {
	// 根据邮箱查询用户
	_, err := u.Query.User.Where(u.Query.User.Email.Eq(signupReq.Email)).First()
	// 查询出错
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerror.NewErrCodeMsg(xerror.DB_SEARCH_ERR, err.Error())
	}
	// 用户已存在
	if err == nil {
		return nil, xerror.NewErrCodeMsg(400, "user already exist")
	}
	// 创建新用户
	newUser := entity.User{
		ID:       uuid.New().String(),
		Name:     signupReq.Name,
		Email:    signupReq.Email,
		Password: utils.Md5ByString(signupReq.Password),
	}
	dao.User.Create(&newUser)
	var signupRes types.SignupRes
	copier.Copy(&signupRes, newUser)
	return &signupRes, nil
}

func (u *UserService) Login(loginReq *types.LoginReq) (*types.LoginRes, error) {
	// 根据邮箱查询用户
	user, err := u.Query.User.Where(u.Query.User.Email.Eq(loginReq.Email)).First()
	// 查询出错
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerror.NewErrCodeMsg(xerror.DB_SEARCH_ERR, err.Error())
	}
	// 用户不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerror.NewErrCodeMsg(400, "user not exist")
	}
	// 密码错误
	if user.Password != utils.Md5ByString(loginReq.Password) {
		return nil, xerror.NewErrCodeMsg(400, "password error")
	}
	// 生成token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, xerror.NewErrCodeMsg(xerror.SERVER_ERROR, err.Error())
	}
	var loginRes types.LoginRes
	loginRes.Token = token
	copier.Copy(&loginRes.User, user)
	return &loginRes, nil
}

func (u *UserService) UserInfo(userId string) (*types.UserInfoRes, error) {
	// 根据id查询用户
	user, err := u.Query.User.Where(u.Query.User.ID.Eq(userId)).First()
	// 查询出错
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerror.NewErrCodeMsg(xerror.DB_SEARCH_ERR, err.Error())
	}
	// 用户不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerror.NewErrCodeMsg(xerror.AUTH_CHECK_FAILURE, "user not exist")
	}
	var userInfoRes types.UserInfoRes
	copier.Copy(&userInfoRes, user)
	return &userInfoRes, nil
}
