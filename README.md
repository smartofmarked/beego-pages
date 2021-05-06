# beego-pages
简单的beego分页
# 使用
#### 控制器

import (
 "github.com/smartofmarked/beego-pages/pages"
)

	user.TplName = "admin/user.html"
	pages.SetCon(&user.Controller) // 传入*beego.Controller  可以放在控制器 Prepare 初始化方法中，每次请求自动传入
	pages.SetPageCount(2)          // 分页数量(每页多少条)
	user.Data["pageHtml"] = pages.Page(count) // 需要分页的条数
  
  
#### 模板
![image](https://user-images.githubusercontent.com/40753219/117231802-8a8d4f00-ae52-11eb-93d7-8a6fe3e58c85.png)

#### 页面

![image](https://user-images.githubusercontent.com/40753219/117231822-9547e400-ae52-11eb-9854-b5cb175e7707.png)
