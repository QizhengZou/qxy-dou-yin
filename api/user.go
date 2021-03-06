package api

import (
	"fmt"
	"net/http"
	"qxy-dy/middleware"
	"qxy-dy/model"
	"qxy-dy/serializer"
	"time"

	"github.com/gin-gonic/gin"
)

type UserService struct{}

func (UserService) isFollow(myId string, hisId string) (isFollow bool, err error) {
	var count int64
	model.DB.Model(&model.Follow{}).Where("user_id = ? AND to_user_id = ?", myId, hisId).Count(&count)
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (UserService) isFollowByUint(myId string, hisId uint) (isFollow bool, err error) {
	var count int64
	model.DB.Model(&model.Follow{}).Where("user_id = ? AND to_user_id = ?", myId, hisId).Count(&count)
	if count > 0 {
		return true, nil
	}
	return false, nil
}

var MyUserService UserService

type UserLoginResponse struct {
	serializer.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	serializer.Response
	User serializer.User `json:"user"`
}

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var user model.User
	username := c.Query("username")
	password := c.Query("password")
	fmt.Printf("username:%#v\n", username)
	fmt.Printf("password:%#v\n", password)
	// 查询是否注册过用户
	var count int64 = 0
	model.DB.Model(&model.User{}).Where("username = ?  ", username).Count(&count)
	fmt.Printf("user:%#v", user)
	if count > 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: serializer.Response{StatusCode: 1, StatusMsg: "用户名已存在"},
		})
		return
	}

	// 创建用户
	u := model.User{Username: username, Password: password, FollowCount: 0, FollowerCount: 0}
	result := model.DB.Model(&model.User{}).Create(&u)
	fmt.Printf("result:%#v\n", result)

	// 根据用户信息生成token
	claims := &middleware.JWTClaims{
		Id:       uint64(user.ID),
		Username: username,
		Password: password,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(middleware.ExpireTime)).Unix()
	singedToken, err := middleware.GenToken(claims)
	// 判断是否生成token成功
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: serializer.Response{StatusCode: 0, StatusMsg: "注册成功"},
		UserId:   int64(u.ID),
		Token:    singedToken,
	})

}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var user model.User
	username := c.Query("username")
	password := c.Query("password")
	fmt.Printf("username:%#v\n", username)
	fmt.Printf("password:%#v\n", password)

	if err := model.DB.Model(&model.User{}).Where("username = ? AND password = ?", username, password).Find(&user).Error; err == nil {
		if user.Username != "" {
			// 根据用户信息生成token
			claims := &middleware.JWTClaims{
				Id:       uint64(user.ID),
				Username: username,
				Password: password,
			}
			claims.IssuedAt = time.Now().Unix()
			claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(middleware.ExpireTime)).Unix()
			singedToken, err := middleware.GenToken(claims)
			// 判断是否生成token成功
			if err != nil {
				c.String(http.StatusNotFound, err.Error())
				return
			}
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: serializer.Response{StatusCode: 0, StatusMsg: "登陆成功"},
				UserId:   int64(user.ID),
				Token:    singedToken,
			})
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: serializer.Response{StatusCode: 1, StatusMsg: "用户名或密码错误"},
			})
		}
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: serializer.Response{StatusCode: 1, StatusMsg: "用户名或密码错误"},
		})
	}

}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	var user model.User
	id := c.Query("user_id")
	// 调用方法获取用户自己的id
	myId := CurrentUser(c)
	if err := model.DB.Model(&model.User{}).Where("id = ?", id).Find(&user).Error; err == nil {

		if isFollow, err := MyUserService.isFollow(myId, id); err == nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: serializer.Response{StatusCode: 0, StatusMsg: "获取用户信息成功"},
				User: serializer.User{
					Id:            int64(user.ID),
					Name:          user.Username,
					FollowCount:   user.FollowCount,
					FollowerCount: user.FollowerCount,
					IsFollow:      isFollow,
				},
			})
		}
	} else {
		c.JSON(200, serializer.Response{
			StatusCode: 1,
			StatusMsg:  "获取信息失败",
		})
	}
}
