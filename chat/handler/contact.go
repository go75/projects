package handler

import (
	"chat/common/res"
	"chat/dao"
	"chat/global"
	"chat/log"
	"chat/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AddFriend(c *gin.Context) {
	ownid := c.GetUint("id")
	dstid := c.PostForm("dstid")
	if strconv.Itoa(int(ownid)) == dstid {
		res.Err(c, "不能添加自己为好友")
		return
	}

	dstidint, err := strconv.Atoi(dstid)
	if err != nil {
		res.Err(c, "对方id格式错误")
		return
	}
	tmp := &model.Contact{
		Ownerid: ownid,
		Dstid: uint(dstidint),
		Cate: model.CONCAT_CATE_USER,
	}
	global.DB.Where("ownerid = ? and dstid = ? and cate = ?", ownid, dstid, model.CONCAT_CATE_USER).First(tmp)
	if tmp.Id > 0 {
		res.Err(c, "该用户已添加")
        return
    }
	global.DB.Begin()
	contact0 := &model.Contact {
		Ownerid: ownid,
        Dstid: uint(dstidint),
        Cate: model.CONCAT_CATE_USER,
		Createat: time.Now(),
	}
	contact1 := &model.Contact {
		Ownerid: uint(dstidint),
        Dstid: ownid,
        Cate: model.CONCAT_CATE_USER,
        Createat: time.Now(),
	}
	e0 := dao.CreateContact(contact0).Error
	e1 := dao.CreateContact(contact1).Error
	if e0 != nil || e1 != nil {
		global.DB.Rollback()
		log.Info.Println(e0)
		log.Info.Panicln(e1)
		res.Err(c, "添加失败")
		return
	}
	global.DB.Commit()
	res.Ok(c, "已成功发起请求", nil)
}

func LoadFriend(c *gin.Context) {
	id := c.GetUint("id")
	if id == 0 {
		res.Err(c, "身份错误")
		return
	}

	contact := &model.Contact {
		Ownerid: id,
		Cate: model.CONCAT_CATE_USER,
	}

	list, _ := dao.GetContactsByOwneridAndCate(contact)
	friendIds :=make([]uint,len(list))
	for i:=0;i<len(friendIds);i++ {
		friendIds[i] = list[i].Dstid
	}
	friends := make([]model.User,0)
	global.DB.Find(&friends, friendIds)
	log.Info.Printf("%v\n", friends)
	res.OkList(c, friends, len(friends))
}

func LoadCommunity(c *gin.Context) {
	id := c.GetUint("id")
	if id == 0 {
		res.Err(c, "身份错误")
		return
	}

	contact := &model.Contact {
        Ownerid: id,
        Cate: model.CONCAT_CATE_COMUNITY,
    }

	list, _ := dao.GetContactsByOwneridAndCate(contact)
	comIds :=make([]uint, len(list))
	for i:=0;i<len(comIds);i++ {
        comIds[i] = list[i].Dstid
    }
	coms := make([]model.Community,0)
	global.DB.Find(&coms, comIds)
	log.Info.Printf("%v\n", coms)
	res.OkList(c, coms, len(coms))	
}