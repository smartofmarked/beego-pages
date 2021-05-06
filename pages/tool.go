package pages

import (
	"bytes"
	"github.com/astaxie/beego"
	"runtime"
	"strconv"
)

//SetCon 设置当前的请求对象
func SetCon(Con *beego.Controller){
	requestMap[GetGID()] = Con
}

// GetCon 获取当前请求的对象
func GetCon()*beego.Controller{
	return requestMap[GetGID()]
}

//DelCon 删除当前的请求对象
func DelCon(){
	delete(requestMap,GetGID())
}

// GetGID 获取当前协程号
func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// SetPageCount 设置分页总数
func SetPageCount(count int){
	myLimit = count
}