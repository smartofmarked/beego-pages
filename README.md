# beego-pages
简单的beego分页
# 使用
#### 控制器

 user.TplName = "user/user.html"
 page.SetCon(&user.Controller) // 传入*beego.Controller 可以放在控制器 Prepare 方法中，防止每次设置
 page.SetPageCount(2)          //设置分页数量 默认15
 user.Data["pageHtml"] = page.Page(count) //传入总条数
  
  
#### 模板
![image](https://user-images.githubusercontent.com/40753219/117231802-8a8d4f00-ae52-11eb-93d7-8a6fe3e58c85.png)

![image](https://user-images.githubusercontent.com/40753219/117231822-9547e400-ae52-11eb-9854-b5cb175e7707.png)
