package models

import (
	"fmt"
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/zhufuyi/mongo"
)

var server = "mongodb://collectdata:123456@192.168.101.88:27018/collectdata"

func init() {
	err := mongo.InitializeMongodb(server)
	if err != nil {
		panic(err)
	}
}

func TestUserMsg_Insert(t *testing.T) {
	// 插入数据，逐条插入
	for i := 1; i <= 100; i++ {
		e := &UserMsg{
			UserID:     "1146955134308192251",
			FileID:     "36dbff15-1d39-4bf6-8487-4476a56f3cea",
			TaskID:     fmt.Sprintf("42207061225347891%d", i),
			PolicyName: fmt.Sprintf("测试策略%d", i),
			Message:    fmt.Sprintf("运行策略程序%d", i),
		}

		if err := e.Insert(); err != nil {
			t.Error(err)
		}
	}
}

func TestFindUserMsg(t *testing.T) {
	taskID := "422070612253478911"

	e, err := FindUserMsg(bson.M{"taskID": taskID}, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if e.TaskID != taskID {
		t.Errorf("got: %s, expect:%s", e.TaskID, taskID)
	}
}

func TestCountUserMsgs(t *testing.T) {
	query := bson.M{"fileID": "36dbff15-1d39-4bf6-8487-4476a56f3cea"}
	total, err := CountUserMsgs(query)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("一共有%d条记录\n", total)
}

func TestFindUserMsgs(t *testing.T) {
	query := bson.M{"fileID": "36dbff15-1d39-4bf6-8487-4476a56f3cea"}

	// 获取满足条件的记录数量
	total, _ := CountUserMsgs(query)
	// 每页多少条记录
	limit := 10
	// 计算一共多少页
	totalPage := total / limit
	fmt.Printf("一共有%d页\n", totalPage)
	// 查看第几页记录，从第1也开始
	page := 5

	es, err := FindUserMsgs(query, nil, page, limit)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("获取到%d条记录\n", len(es))

	if len(es) > 0 {
		if es[0].FileID != query["fileID"] {
			t.Errorf("got: %s, expect:%s", es[0].FileID, query["fileID"])
		}
	}
}

func TestUpdateUserMsg(t *testing.T) {
	expect := "36dbff15-1d39-4bf6-8487-4476a56f3ceb"

	query := bson.M{"fileID": "36dbff15-1d39-4bf6-8487-4476a56f3cea"}
	update := bson.M{"$set": bson.M{"fileID": expect}}

	err := UpdateUserMsg(query, update)
	if err != nil {
		t.Error(err)
		return
	}

	e, err := FindUserMsg(bson.M{"fileID": expect}, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if e.FileID != expect {
		t.Errorf("got: %s, expect:%s", e.FileID, expect)
	}
}

func TestUpdateUserMsgs(t *testing.T) {
	expect := "36dbff15-1d39-4bf6-8487-4476a56f3ceb"
	query := bson.M{"fileID": "36dbff15-1d39-4bf6-8487-4476a56f3cea"}
	update := bson.M{"$set": bson.M{"fileID": expect}}

	n, err := UpdateUserMsgs(query, update)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("更新成功个数 %d\n", n)
}

func TestFindAndModifyUserMsg(t *testing.T) {
	query := bson.M{"taskID": "422070612253478911"}
	update := bson.M{"$set": bson.M{"fileID": "36dbff15-1d39-4bf6-8487-4476a56f3ceG"}}

	UserMsg, err := FindAndModifyUserMsg(query, update)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%+v", *UserMsg)
}

func TestDeleteUserMsg(t *testing.T) {
	taskID := "422070612253478911"

	n, err := DeleteUserMsg(bson.M{"taskID": taskID})
	if err != nil {
		t.Error(err)
		return
	}

	if n > 0 {
		fmt.Printf("delete taskId=%s success. n=%d\n", taskID, n)
	}
}
