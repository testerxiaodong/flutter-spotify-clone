package services

import (
	"server/common/xerror"
	"server/models/dao"
	"server/models/entity"
	"server/types"
	"server/utils"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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
	user, err := u.Query.User.Where(u.Query.User.Email.Eq(signupReq.Email)).Find()
	if err != nil {
		return nil, xerror.NewErrCodeMsg(xerror.DB_SEARCH_ERR, err.Error())
	}
	if len(user) != 0 {
		return nil, xerror.NewErrCodeMsg(400, "user already exist")
	}
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
