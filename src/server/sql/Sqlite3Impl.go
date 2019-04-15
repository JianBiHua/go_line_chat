package sql

import (
	"GoWorkspace/go_line_chat/src/server/chatlog"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Sqlite3Impl is sqlite3 impl
type Sqlite3Impl struct {
	db *sql.DB
}

var instance *Sqlite3Impl

// GetInstance is 模拟单例实现功能
func GetInstance() *Sqlite3Impl {
	if instance == nil {
		instance = new(Sqlite3Impl)
	}

	return instance
}

// NewSqlite3Impl is static create sqlite3 impl
func NewSqlite3Impl() *Sqlite3Impl {
	var s = new(Sqlite3Impl)
	return s
}

// log is show log when Chat mode is SqlMode
func (s *Sqlite3Impl) log(msg string) {
	chatlog.Append(chatlog.LOGSQL, msg)
}

// Open 打开数据库，并创建表
//
// out:
// error: 打开数据库时的错误。
func (s *Sqlite3Impl) Open(path string) error {
	// 打开数据库
	var err error
	s.db, err = sql.Open("sqlite3", path)
	if err == nil {
		s.log("Sqlite3 Is Open")
	} else {
		s.log(fmt.Sprintf("Open sqlite3 error: %v", err))
	}

	return err
}

// Close 关闭数据库
func (s *Sqlite3Impl) Close() {
	s.log("Sqlite3 Is Closed")

	s.db.Close()
}

// CreateTable is 创建数据表
// func (s *Sqlite3Impl) createTable() {
// 	var err error
// 	err = createTable(SQLCommandUser)
// 	s.log(fmt.Sprintf("error: %v, %s", err, SQLTableUser))
// 	err = createTable(SQLCommandUserChatGroup)
// 	s.log(fmt.Sprintf("error: %v, %s", err, SQLTableChatGroup))
// 	err = createTable(SQLCommandUserChatGroupMember)
// 	s.log(fmt.Sprintf("error: %v, %s", err, SQLTableChatGroupMember))
// 	err = createTable(SQLCommandUserChatMsg)
// 	s.log(fmt.Sprintf("error: %v, %s", err, SQLTableChatMsg))
// 	err = createTable(SQLCommandFriends)
// 	s.log(fmt.Sprintf("error: %v, %s", err, SQLTableFriends))
// }
func (s *Sqlite3Impl) CreateTable(cmd string) error {
	_, err := s.db.Exec(cmd)
	return err
}

// UpdateOrDelete is run sql command
//
// Example:
//
// 	delete:
//	err := UpdateOrDelete(fmt.Spintf("DELETE FROM UserTable1 WHERE id=1"))
//
//	update:
//	err := UpdateOrDelete(fmt.Spintf("UPDATE UserTable1 SET userName=\"user2\",password=\"password2\" WHERE id=1")
//
func (s *Sqlite3Impl) UpdateOrDelete(prepare string) (int64, error) {
	//插入数据
	// stmt, err := s.db.Prepare(prepare)
	// if err != nil {
	// 	return -1, err
	// }
	// defer stmt.Close()

	res, err := s.db.Exec(prepare)
	if err != nil {
		return -1, err
	}

	id, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	return id, err
}

// Insert is insert sql object to table
//
// Example:
//
// Insert(fmt.Sprintf("INSERT INTO %s(userName, password) values(%s,%s)",
//		sql.SQLTableUser, "user", "password"))
func (s *Sqlite3Impl) Insert(prepare string) (int64, error) {
	//插入数据
	// stmt, err := s.db.Prepare(prepare)
	// if err != nil {
	// 	return -1, err
	// }
	// defer stmt.Close()

	res, err := s.db.Exec(prepare)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, err
}

// Get is get data base
//
// Example:
// var cmd = fmt.Sprintf("SELECT id,userName,password FROM UserTable1 WHERE id=%d", user.id)
// err := Get(cmd, func(p ...interface{}){
// 	for key, value := range param {
// 		fmt.Print("", key, value)
// 	}
//}, &userName, &password)
//
func (s *Sqlite3Impl) Get(cmd string, f func(...interface{}), param ...interface{}) error {
	rows, err := s.db.Query(cmd)
	if err != nil {
		s.log(fmt.Sprintf("Get error : %v", err))
		return err
	}

	defer rows.Close()

	var success = false
	for rows.Next() {
		err := rows.Scan(param)
		if err != nil {
			s.log(fmt.Sprintf("Get error : %v", err))
			return err
		}

		f(param)
		success = true
	}

	if success {
		return nil
	}

	return errors.New("not found")
}
