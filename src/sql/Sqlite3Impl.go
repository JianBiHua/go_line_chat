package sql

import (
	"GoWorkspace/go_line_chat/src/toolunit"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite3Impl struct {
	db *sql.DB
}

func NewSqlImpl() *Sqlite3Impl {
	var s = new(Sqlite3Impl)
	return s
}

// Open 关闭数据库
func (s *Sqlite3Impl) Open() {
	// 打开数据库
	var path, _ = toolunit.GetPathInstance().GetDBPath("foo.db")
	fmt.Println(path)
	var err error
	s.db, err = sql.Open("sqlite3", path)
	fmt.Println(err)

	// 创建表
	s.createTable()
}

// Close 关闭数据库
func (s *Sqlite3Impl) Close() {
	s.db.Close()
}

// 创建数据表
func (s *Sqlite3Impl) createTable() {
	// 联合主键设置
	//CREATE TABLE IF NOT EXISTS "UserTable" (
	//	"id" INTEGER KEY AUTOINCREMENT,
	//	"userName" varchar(30) unique,
	//	"password" varchar(30),
	//	"sex" int(2) NULL,
	//	"name" varchar(20) NULL,
	//	"age" TIMESTAMP default (datetime('now', 'localtime')) ,
	//"icon" varchar (200) NULL,
	//"signature" varchar (500) NULL,
	//"createDate" TIMESTAMP default (datetime('now', 'localtime')) ,
	//"lastDate" TIMESTAMP default (datetime('now', 'localtime'),
	// ,primary key (id,userName));
	//
	//
	// 创建索引
	// CREATE INDEX userId ON UserTable1(userName);
	//
	// unique
	// 唯一性;唯一约束
	sql_table := `
		CREATE TABLE IF NOT EXISTS "UserTable1" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		 	"userName" varchar(30) unique,
			"password" varchar(30),
			"sex" int(2) NULL,
		 	"name" varchar(20) NULL,
			"age" TIMESTAMP default (datetime('now', 'localtime')) ,
			"icon" varchar (200) NULL,
			"signature" varchar (500) NULL,
			"createDate" TIMESTAMP default (datetime('now', 'localtime')) ,
			"lastDate" TIMESTAMP default (datetime('now', 'localtime'))
		);
		
		CREATE INDEX IF NOT EXISTS userId ON UserTable1(userName);
		`

	_, err := s.db.Exec(sql_table)
	fmt.Println("error: ", err)
}

// 插入用户信息
func (s *Sqlite3Impl) Insert(user *User) int64 {
	// 设置插入类型
	stmt, err := s.db.Prepare("INSERT INTO UserTable1(userName,password) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	// 插入数据
	result, err := stmt.Exec(user.UserName, user.Password)
	if err != nil {
		fmt.Printf("add error: %v", err)
		return -1
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted id is ", lastID)
	// 获取插入的ID号
	user.id = lastID

	return user.id
}

func (s *Sqlite3Impl) Remove(user User) {
	stmt, err := s.db.Prepare("DELETE FROM UserTable1 WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(user.id)
	if err != nil {
		log.Fatal(err)
	}
	affectNum, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete affect rows is ", affectNum)
}

func (s *Sqlite3Impl) Update(user User) {
	if user.id == -1 {
		return
	}

	stmt, err := s.db.Prepare("UPDATE UserTable1 SET userName=?,password=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(user.UserName, user.Password, user.id)
	if err != nil {
		log.Fatal(err)
	}
	affectNum, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update affect rows is ", affectNum)
}

func (s *Sqlite3Impl) Search(user *User) {
	var cmd = fmt.Sprintf("SELECT id,userName,password FROM UserTable1 WHERE id=%d", user.id)
	fmt.Println(cmd)
	rows, err := s.db.Query(cmd)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		//rows.Columns()
		err := rows.Scan(&user.id, &user.UserName, &user.Password)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(user.id, user.UserName, user.Password)
	}
}
