package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/zhufuyi/mongo"
)

const UserMsgCollectionName = "userMsg"

// UserMsg 用户消息
type UserMsg struct {
	PublicFields `bson:",inline"`

	UserID     string      `bson:"userID" json:"userID"`         // 用户id
	FileID     string      `bson:"fileID" json:"fileID"`         // 文件id
	TaskID     string      `bson:"taskID" json:"taskID"`         // 策略运行id，需要建立索引
	PolicyName string      `bson:"policyName" json:"policyName"` // 用户命名的策略名称
	Message    interface{} `bson:"message" json:"message"`       // 信息内容
}

// Insert 插入一条新的记录
func (object *UserMsg) Insert() (err error) {
	mconn := mongo.GetSession()
	defer mconn.Close()
	object.SetFieldsValue()
	return mconn.Insert(UserMsgCollectionName, object)
}

// FindUserMsg 获取单条记录
func FindUserMsg(selector bson.M, field bson.M) (*UserMsg, error) {
	mconn := mongo.GetSession()
	defer mconn.Close()

	object := &UserMsg{}
	return object, mconn.FindOne(UserMsgCollectionName, object, selector, field)
}

// FindUserMsgs 获取多条记录
func FindUserMsgs(selector bson.M, field bson.M, page int, limit int, sort ...string) ([]*UserMsg, error) {
	mconn := mongo.GetSession()
	defer mconn.Close()

	// 默认从第一页开始
	if page < 1 {
		page = 1
	}

	objects := []*UserMsg{}
	return objects, mconn.FindAll(UserMsgCollectionName, &objects, selector, field, page-1, limit, sort...)
}

// UpdateUserMsg 更新单条记录
func UpdateUserMsg(selector, update bson.M) (err error) {
	mconn := mongo.GetSession()
	defer mconn.Close()

	return mconn.UpdateOne(UserMsgCollectionName, selector, updatedTime(update))
}

// UpdateUserMsgs 更新多条记录
func UpdateUserMsgs(selector, update bson.M) (n int, err error) {
	mconn := mongo.GetSession()
	defer mconn.Close()

	return mconn.UpdateAll(UserMsgCollectionName, selector, updatedTime(update))
}

// FindAndModifyUserMsg 更新并返回最新记录
func FindAndModifyUserMsg(selector bson.M, update bson.M) (*UserMsg, error) {
	mconn := mongo.GetSession()
	defer mconn.Close()

	object := &UserMsg{}
	return object, mconn.FindAndModify(UserMsgCollectionName, object, selector, updatedTime(update))
}

// CountUserMsgs 统计数量，不包括删除记录
func CountUserMsgs(selector bson.M) (n int, err error) {
	mconn := mongo.GetSession()
	defer mconn.Close()

	return mconn.Count(UserMsgCollectionName, excludeDeleted(selector))
}

// DeleteUserMsg 删除记录
func DeleteUserMsg(selector bson.M) (int, error) {
	mconn := mongo.GetSession()
	defer mconn.Close()

	return mconn.UpdateAll(UserMsgCollectionName, selector, deletedTime(bson.M{}))
}
