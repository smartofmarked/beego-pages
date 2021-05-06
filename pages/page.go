package pages

import (
	"fmt"
	"github.com/astaxie/beego"
)

// MyLimit 分页总数 默认为15
var myLimit = 15

// Admin 当前请求的con
var requestMap = map[uint64] *beego.Controller{}

type myPage struct{
	page 		int 				//当前页
	maxPage 	int					//最大页数
	href 		string 				//链接
	result 		string 				//结果
	beegoCon 	*beego.Controller	//当前请求的控制器
	start 		int 				//中间几页开始位置
	end			int 				//中间几页结尾位置
	before 		int 				//当前页之前有几页
	after		int 				//当前页之后有几页
	pageStr 	string 				//url中 page 值
}

// Page 分页方法
func Page(pageCount int) (result string){
	if pageCount < 1 {
		return
	}
	//当前最大页
	maxPage := pageCount / myLimit
	if pageCount % myLimit != 0{
		//go 整形 除以 整形 会向下取整
		maxPage ++
	}
	if maxPage < 2 {
		//不需要分页
		return
	}
	//获取当前请求协程的controller
	beegoCon := GetCon()
	page,_ := beegoCon.GetInt("page")
	if page < 1 {
		page = 1
	}
	if page > maxPage {
		page = maxPage
	}
	p := myPage{
		page 		: page,
		maxPage 	: maxPage,
		beegoCon	: beegoCon,
		before  	: 2,
		after		: 2,
		pageStr 	: "page",
	}
	return p.getResult()
}


//解析当前url和参数
func (myPage *myPage) decodeParam(){
	con,fun := myPage.beegoCon.GetControllerAndAction()
	myPage.href = myPage.beegoCon.URLFor(con + "." + fun) + "?"
	for k,v :=  range myPage.beegoCon.Input(){
		if k == myPage.pageStr {
			continue
		}
		myPage.href += k + "=" + v[0] + "&"
	}
	return
}

//拼接分页html
func (myPage *myPage) andHtml(){
	myPage.result = "<div class=\"layui-box layui-laypage layui-laypage-default\" >"
	//上一页
	if myPage.page >1 {
		myPage.result += fmt.Sprintf("<a href=\"" + myPage.href +  myPage.pageStr +"=%v\">上一页</a>",myPage.page -1 )
	}else{
		myPage.result += "<a href=\"javascript:;\" class=\"layui-laypage-prev layui-disabled\" >上一页</a>"
	}
	//中间几页
	for i := myPage.start ; i <= myPage.end ; i ++ {
		if myPage.page != i {
			myPage.result += fmt.Sprintf("<a href=\"" + myPage.href + myPage.pageStr +"=%v\" >%v</a>",i,i)
		}else{
			//当前页
			myPage.result += fmt.Sprintf("<span class=\"layui-laypage-curr\"><em class=\"layui-laypage-em\"></em><em>%v</em></span>",i)
		}
	}
	//下一页
	if myPage.page < myPage.maxPage {
		myPage.result += fmt.Sprintf("<a href=\"" + myPage.href + myPage.pageStr + "=%v\">下一页</a>",myPage.page +1 )
	}else{
		myPage.result += "<a href=\"javascript:;\" class=\"layui-laypage-prev layui-disabled\" >下一页</a>"
	}
	myPage.result += "</div>"
	return
}

//计算中间几页
func (myPage *myPage) cal() {
	myPage.start = 1
	myPage.end = myPage.maxPage
	beforeSub := myPage.page - myPage.before
	afterAdd  := myPage.page + myPage.after
	//当前页之前
	if beforeSub > 1 {
		myPage.start = beforeSub
	}
	//当前页之后
	if afterAdd < myPage.maxPage {
		myPage.end = afterAdd
	}
}

//执行并响应结果
func (myPage *myPage) getResult() string{
	//解析参数
	myPage.decodeParam()
	//计算分页
	myPage.cal()
	//拼接html
	myPage.andHtml()
	//一般分页只会用一次，用过即删除当前请求对象
	DelCon()
	return myPage.result
}