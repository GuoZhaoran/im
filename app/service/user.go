package service

import (
	"IMtest/app/model"
	"IMtest/app/util"
	"errors"
	"time"
)

type UserService struct {
}

//用户注册
func (s *UserService) Register(mobile, plainpwd, nickname, avatar, sex string) (user model.User, err error) {
	//检测手机号是否存在
	regUser := model.User{}
	_, err = DbEngine.Where("mobile=?", mobile).Get(&regUser)
	if err != nil {
		return regUser, err
	}
	//如果存在返回提示已经注册
	if regUser.Id > 0 {
		return regUser, errors.New("该手机号已经注册")
	}
	//否则拼接插入数据
	regUser.Mobile = mobile
	regUser.Avatar = avatar
	regUser.Nickname = nickname
	regUser.Sex = sex
	regUser.Salt = util.GenRandomStr(6)
	regUser.Passwd = util.MakePasswd(plainpwd, regUser.Salt)
	regUser.Createat = time.Now()
	//token	可以是一个随机数
	regUser.Token = util.GenRandomStr(32)

	//插入 InsertOne
	_, err = DbEngine.InsertOne(&regUser)

	return regUser, nil
}

//用户登录
func (s *UserService) Login(mobile, plainpwd string) (user model.User, err error) {
	//数据库操作
	loginUser := model.User{}
	DbEngine.Where("mobile = ?", mobile).Get(&loginUser)
	if loginUser.Id == 0 {
		return loginUser, errors.New("用户不存在")
	}
	//判断密码是否正确
	if !util.ValidatePasswd(plainpwd, loginUser.Salt, loginUser.Passwd) {
		return loginUser, errors.New("密码不正确")
	}
	//刷新用户登录的token值
	token := util.GenRandomStr(32)
	loginUser.Token = token
	DbEngine.ID(loginUser.Id).Cols("token").Update(&loginUser)

	//返回新用户信息
	return loginUser, nil
}

//查找某个用户
func (s *UserService) Find(userId int64) (user model.User) {
	findUser := model.User{}
	DbEngine.ID(userId).Get(&findUser)

	return findUser
}
