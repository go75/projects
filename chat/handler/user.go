package handler

import (
	"chat/common/res"
	"chat/dao"
	"chat/global"
	"chat/log"
	"chat/model"
	"chat/utils"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func UserRegistPage(c *gin.Context) {
    c.HTML(http.StatusOK, "regist.html", nil)
}

// UserRegist
// @Tags 用户模块
// @Summary 用户注册
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /user/regist/:name/:pwd/:check [post]
func UserRegist(c *gin.Context) {
	mobile := c.PostForm("mobile")
	passwd :=  c.PostForm("passwd")
	salt := fmt.Sprintf("%10d", rand.Int31())
	log.Info.Println(mobile)
	log.Info.Println(passwd)
	log.Info.Println(salt)
	identity := &model.UserIdentity{
		Mobile:  mobile,
        Pwd:  utils.MakePwd(passwd, salt),
        Salt: salt,
	}

	global.DB.Begin()
	db := dao.CreateUserIdentity(identity)
	if db.RowsAffected != 1 {
		log.Info.Printf("注册失败:%s\n", db.Error.Error())
		global.DB.Rollback()
		res.Err(c, "注册失败")
		return
	}

	user := &model.User {
		Mobile: mobile,
		Avatar   :"",
		Sex      :model.SEX_UNKNOW,
		Nickname :fmt.Sprintf("user%10d", rand.Int31()),
	}

	db = dao.CreateUser(user)
	if db.RowsAffected != 1 {
		global.DB.Rollback()
		res.Err(c, "注册失败")
		return
	}
	global.DB.Commit()
	res.Ok(c, "注册成功", nil)
}

// UserLogin
// @Tags 用户模块
// @Summary 用户登录
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /user/login/:name/:pwd [post]
func UserLogin(c *gin.Context) {
	log.Info.Println("user login")
	mobile := c.PostForm("mobile")
	passwd := c.PostForm("passwd")
	log.Info.Println(mobile)
	log.Info.Println(passwd)
	identity := &model.UserIdentity {
		Mobile: mobile,
	}
	dao.GetUserIdentity(identity)
	if ok := utils.CheckPwd(passwd, identity.Salt, identity.Pwd); ok {
		log.Info.Println(ok)
		user := &model.User{
			Mobile: mobile,
		}
		dao.GetUserByMobile(user)
		token, err := utils.GenerateToken(user.Id, user.Mobile, user.Nickname)
		if err != nil {
			log.Error.Println("generate token err: ", err)
			res.Err(c, "登录失败")
			return
		}

		log.Info.Println("generate token: ", token)
		res.Ok(c, "登录成功", token)
	} else {
		res.Err(c, "密码错误")
	}
}