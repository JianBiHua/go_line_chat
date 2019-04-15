package toolunit

import (
	"os"
	"path/filepath"
)

// Path is path tool
type Path struct {
}

// GetPathInstance is get static Path pointer
// 返回:
// *Path  a path pointer
func GetPathInstance() *Path {
	var path = new(Path)
	path.initPath()
	return path
}

// mkdir is create dir of path
// params:
// 	path: will create dir of path
func (p *Path) mkdir(path string) {
	_, err := os.Stat(path)
	if !os.IsExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

func (p *Path) initPath() {
	//
	// 创建数据库目录
	var path = p.GetDBPath("")
	p.mkdir(path)

	// 创建日志目录。
	path = p.GetLogPath("")
	p.mkdir(path)
}

// GetRootPath 获取运行程序所在的路径
// 返回：
// string: 返回运行程序所在的路径
// error: 错误信息
func (p *Path) GetRootPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

// GetExecPath2 获取运行程序所在的路径+name
// 返回：
// string: 返回运行程序所在的路径+name
// error: 错误信息
func (p *Path) GetExecPath2(name string) (string, error) {
	path, err := p.GetRootPath()
	return path + name, err
}

// GetLogPath 获取运行程序所在的路径+name
// 参数:
// name: 日志文件目录下的文件，或者文件夹名称。
// 返回:
// string: 返回运行程序所在的路径+name
// error: 错误信息
func (p *Path) GetLogPath(name string) string {
	path, _ := p.GetRootPath()
	return path + "/log/" + name
}

// GetDBPath is Get DataBase Dir Path
// params:
//	name: is sub dir name
// return:
//	string: is DataBase Dir sub dir path
//  error: is error
func (p *Path) GetDBPath(name string) string {
	path, _ := p.GetRootPath()
	return path + "/databases/" + name
}
