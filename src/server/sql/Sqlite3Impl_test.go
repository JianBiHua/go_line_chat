package sql

import (
	"GoWorkspace/go_line_chat/src/server/toolunit"
	"fmt"
	"testing"
)

func TestSqlite3Impl(t *testing.T) {
	//var impl SqlInterface
	impl := NewSqlite3Impl()

	// 打开并且链接数据库
	path, err := toolunit.GetPathInstance().GetDBPath(SQLDataBaseName)
	if err != nil {
		t.Errorf("sql path error: %v", err)
	}
	fmt.Println(path)
	impl.Open(path)

	// 创建表
	// _, err = impl.Exec(SQLCommandUser)
	// if err != nil {
	// 	t.Errorf("create %s: %v", SQLTableUser, err)
	// }

	// _, err = impl.Exec(SQLCommandUserChatGroup)
	// if err != nil {
	// 	t.Errorf("create %s: %v", SQLTableChatGroup, err)
	// }

	// _, err = impl.Exec(SQLCommandUserChatGroupMember)
	// if err != nil {
	// 	t.Errorf("create %s: %v", SQLTableChatGroupMember, err)
	// }

	// _, err = impl.Exec(SQLCommandUserChatMsg)
	// if err != nil {
	// 	t.Errorf("create %s: %v", SQLTableChatMsg, err)
	// }

	// _, err = impl.Exec(SQLCommandFriends)
	// if err != nil {
	// 	t.Errorf("create %s: %v", SQLTableFriends, err)
	// }

	// // test insert  and return object id
	// //
	// //	var result interface
	// result, err1 := impl.Exec(fmt.Sprintf("INSERT INTO %s (userName, password) values(\"%s\",\"%s\")", SQLTableUser, "User", "password"))
	// if err1 != nil {
	// 	t.Errorf("error: %v, %v", result, err1)
	// }

	// // test update
	// var userName string
	// var password string
	// err = impl.Get(fmt.Sprintf("SELECT userName,password FROM "+SQLTableUser+" where id= %d", id), func(param ...interface{}) {
	// 	for key, value := range param {
	// 		fmt.Print("", key, value)
	// 	}
	// }, id, userName, password)

	// if err != nil {
	// 	t.Errorf("error: %v", err)
	// }

	// // test update
	// err = impl.Update(fmt.Sprintf("UPDATE "+SQLTableUser+" SET userName=?,password=? WHERE id=%d", id), "User1", "pasw1")
	// fmt.Print("Update: ", err)
	// if err != nil {
	// 	t.Errorf("error: %v", err)
	// }

	// // test del
	// err = impl.Delete("DELETE FROM "+SQLTableUser+" WHERE id=?", id)
	// fmt.Print("Delete: ", err)
	// if err != nil {
	// 	t.Errorf("error: %v", err)
	// }

	impl.Close()
}

// func ExampleNewSqlite3Impl() {
// 	//var impl SqlInterface
// 	impl := NewSqlite3Impl()
// 	// impl.Open()

// 	// // test insert  and return object id
// 	// //
// 	// // id, err := impl.Insert("INSERT INTO "+SQLTableUser+"(userName,password) values(?,?)", "User", "password")
// 	// // fmt.Print("insert: ", id, err)

// 	// // // test update
// 	// // var userName string
// 	// // var password string
// 	// // impl.Get(fmt.Sprintf("SELECT userName,password FROM "+SQLTableUser+"where id= %d", id), func(param ...interface{}) {
// 	// // 	for key, value := range param {
// 	// // 		fmt.Print("", key, value)
// 	// // 	}
// 	// // }, id, userName, password)

// 	// // // test update
// 	// // err = impl.Update(fmt.Sprintf("UPDATE "+SQLTableUser+" SET userName=?,password=? WHERE id=%d", id), "User1", "pasw1")
// 	// // fmt.Print("Update: ", err)

// 	// // // test del
// 	// // err = impl.Delete("DELETE FROM "+SQLTableUser+" WHERE id=?", id)
// 	// // fmt.Print("Delete: ", err)

// 	impl.Close()
// }
