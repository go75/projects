package handler

import (
	"im/common/res"
	"im/dao"
	"im/model"
	"im/utils"

	"github.com/gin-gonic/gin"
)

// GroupRegist
// @Tags 群聊模块
// @Summary 创建群聊
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /group/regist [post]
func GroupRegist(c *gin.Context) {
	groupName := c.PostForm("groupName")
	if groupName == "" {		
		res.Err(c, "群名不能为空")
        return
    }
	masterId := c.GetUint("id")
	group := &model.Group {
		Name: groupName,
		MasterId: masterId,
	}
	db := dao.CreateGroup(group)
	if db.RowsAffected != 1 {
		res.Err(c, "创建失败")
		return
	}
	res.Ok(c, "创建成功")
}

// GroupSession
// @Tags 群聊模块
// @Summary 群聊会话
// @Param groupid formData string false "群聊id"
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /group/session [get]
func GroupSession(c *gin.Context) {
	id := c.PostForm("groupId")
	var idnum uint 
	err := utils.QuerySwitch(id, idnum)
	if err != nil {
		res.Err(c, "群聊号格式错误")
		return
	}

	list, _ := dao.QueryGroupMessage(idnum)
	res.OkWithData(c, "查询成功", list)
}

// GroupList
// @Tags 群聊模块
// @Summary 群聊列表
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /group/list [get]
func GroupList(c *gin.Context) {
	list, _ := dao.QueryGroupSessionsByUserId(c.GetUint("id"))
	res.OkWithData(c, "查询成功", list)
}

// NewGroupFriendsInfo
// @Tags 群聊模块
// @Summary 新群友信息
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /group/new [get]
func NewGroupFriendsInfo(c *gin.Context) {
	list, _ := dao.QueryGroupIdsByMasterId(c.GetUint("id"))
	var msgs []*model.AddGroupMessage
	for _, v := range list {
		ls, _ := dao.GetAddGroupMessagesByGroupId(v.ID)
		msgs = append(msgs, ls...)
	}
	res.OkWithData(c, "获取成功", msgs)
}

// CreateAddGroupMessage
// @Tags 群聊模块
// @Summary 新群友信息
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /group/new [get]
func CreateAddGroupMessage(c *gin.Context) {

}