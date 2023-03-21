package handler

import (
	"fmt"
	"im/common/res"
	resmodel "im/common/resModel"
	"im/dao"
	"im/global"
	"im/log"
	"im/model"
	"im/utils"
	"io"
	"math/rand"
	"os"

	"github.com/gin-gonic/gin"
)

func UserHead(c *gin.Context) {
	c.File("./head/" + c.GetString("name") + ".png")
}

// Session
// UserLogin
// @Tags 用户模块
// @Summary 用户会话
// @Success 200 {string} json{"code", "msg", "data"}
// @Param id formData string false "会话id"
// @Param name formData string false "对方名称"
// @Router /user/session [get]
func Session(c *gin.Context) {
	var id uint
	err := utils.QuerySwitch(c.PostForm("id"), id)
	if err != nil {
		log.Warn.Println("sessionId类型非法")
		res.Err(c, "sessionId类型非法")
		return
	}

	list, _ := dao.QueryUserMessageByIds(id, c.GetUint("id"))
	res.OkWithData(c, "查询成功", list)
}

// GetUserList
// UserLogin
// @Tags 用户模块
// @Summary 用户列表
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /user/list [get]
func GetUserList(c *gin.Context) {
	res.OkWithData(c, "查询成功", dao.GetUserList())
}

// UserLogin
// @Tags 用户模块
// @Summary 用户登录
// @Success 200 {string} json{"code", "msg", "data"}
// @Param name formData string false "名称"
// @Param pwd formData string false "密码"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	log.Info.Println("user login")
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	identity := &model.UserIdentity{
		Name: name,
	}
	dao.QueryUserIdentity(identity)

	if ok := utils.CheckPwd(pwd, identity.Salt, identity.Pwd); ok {
		
		user := &model.User{
			Name: name,
		}
		
		dao.GetUserByName(user)
		if user.ID == 0 {
			// 未查询到用户
			res.Err(c, "该用户不存在")
			return
		}
		token, err := utils.GenerateToken(user.ID, name)
		if err != nil {
			log.Error.Println("generate token err: ", err)
			res.Err(c, "登录失败")
			return
		}

		log.Info.Println("generate token: ", token)
		data := resmodel.LoginData {
			ID: user.ID,
			Name: user.Name,
			Token: token,
		}
		res.OkWithData(c, "登录成功", data)
	} else {
		res.Err(c, "密码错误")
	}
}

// UserRegist
// @Tags 用户模块
// @Summary 用户注册
// @Param name formData string false "名称"
// @Param pwd formData string false "密码"
// @Param check formData string false "检查密码"
// @Param head formData file false "选择头像"
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /user/regist [post]
func UserRegist(c *gin.Context) {
	pwd := c.PostForm("pwd")
	if pwd != c.PostForm("check") {
		res.Err(c, "两次密码不一致")
		log.Warn.Println("两次密码不一致")
		return
	}

	name := c.PostForm("name")
	headname := name + ".png"
	salt := fmt.Sprintf("%10d", rand.Int31())
	identity := &model.UserIdentity{
		Name: name,
		Pwd: utils.MakePwd(pwd, salt),
		Salt: salt,
	}
	user := &model.User {
		Name: identity.Name,
	}

	// 先将用户信息存入数据库
	global.DB.Begin()

	db := dao.CreateUserIdentity(identity)
	if db.RowsAffected != 1 {
		global.DB.Rollback()
		res.Err(c, "用户标记信息新增失败")
		log.Warn.Println("用户标记信息新增失败")
		return
	}

	db = dao.CreateUser(user)
	if db.RowsAffected != 1 {
		global.DB.Rollback()
		res.Err(c, "用户新增失败")
		log.Warn.Println("用户新增失败")
		return
	}
	global.DB.Commit()

	// 存储用户头像, 若出现错误,则用户使用默认头像
	file, _, err := c.Request.FormFile("head")
	if err != nil {
		log.Info.Println("头像获取失败: " + err.Error())
		res.Ok(c, "注册成功,头像获取失败,请登录后切换")
		return
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err!=nil {
		res.Ok(c, "注册成功,头像读取失败,请登录后切换")
		return
	}

	filepath := "./head/" + headname

	tmp, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0666)
	if err!=nil {
		res.Ok(c, "注册成功,头像创建失败,请登录后切换")
		return
	}
	defer tmp.Close()
	_, err = tmp.Write(content)
	if err != nil {
		os.Remove(filepath)
		res.Ok(c, "注册成功,头像内容填充是失败,请登录后切换")
		return
	}
	
	res.Ok(c, "注册成功")
}

// DeleteUser
// @Tags 用户模块
// @Summary 用户注销
// @Param pwd formData string false "密码"
// @Param check formData string false "检查密码"
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /user/delete [delete]
func DeleteUser(c *gin.Context) {
	pwd := c.PostForm("pwd")

	if pwd != c.PostForm("check") {
		res.Err(c, "两次密码不一致")
		return
	}

	name := c.GetString("name")
	identity := &model.UserIdentity{
		Name: name,
	}
	dao.QueryUserIdentity(identity)
	if ok := utils.CheckPwd(pwd, identity.Salt, identity.Pwd); ok {
		user := &model.User{

		}
		db := dao.DeleteUser(user)
		if db.RowsAffected != 1 {
			res.Err(c, "删除失败")
			return
		}
		db = dao.DeleteUserIdentity(identity)
		if db.RowsAffected != 1 {
			res.Err(c, "删除失败")
			return
		}
	}
	res.Ok(c, "删除成功")
}

// UpdateUser
// @Tags 用户模块
// @Summary 用户信息更新
// @Param sex formData string false "性别"
// @Param introduce formData string false "个人介绍"
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /user/update [post]
func UpdateUser(c *gin.Context) {
	user := &model.User {
		Introduce: c.PostForm("introduce"),
	}
	dao.UpdateUser(user)
	res.Ok(c, "修改成功")
}

// QueryUserByName
// @Tags 用户模块
// @Summary 根据用户名获取用户信息
// @Param name query string false "名称"
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /user/query [get]
func QueryUserByName(c *gin.Context) {
	name := c.Query("name")
	log.Info.Println(name)
	if name == "" {
		// name为空
		res.Err(c, "用户名称为空")
		return
	}
	user := &model.User{
		Name: name,
	}
	db := dao.GetUserByName(user)
	if db.RowsAffected != 1 {
		res.Err(c, "未查询到用户")
	} else {
		res.OkWithData(c, "查询成功", user)
	}
}