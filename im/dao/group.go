package dao

import (
	"im/global"
	"im/model"

	"gorm.io/gorm"
)

// 创建群聊
func CreateGroup(group *model.Group) *gorm.DB {
	return global.DB.Create(group)
}

// 删除群聊
func DeleteGroup(group *model.Group) *gorm.DB {
	return global.DB.Delete(group)
}

func GetGroupById(group *model.Group) *gorm.DB {
	return global.DB.Where("id = ?", group.ID).First(group)
}

// 修改群聊信息
func UpdateGroup(group *model.Group) *gorm.DB {
	return global.DB.Model(group).Updates(model.Group{
		Name: group.Name,
		Introduce: group.Introduce,
	})
}

// 通过群主名称查找群聊名称
func QueryGroupIdsByMasterId(masterId uint) ([]*model.Group, *gorm.DB) {
	var list []*model.Group
	db := global.DB.Select("id").Where("master_id = ?", masterId).Find(list)
	return list, db
}

func GetFuzzyGroupByGroupName(name string) (groups []*model.Group, db *gorm.DB){
	db = global.DB.Where("name like ?", "%"+name+"%").Find(&groups)
	return
}

func QueryGroupByGroupId(group *model.Group) *gorm.DB {
	return global.DB.First(group)
}