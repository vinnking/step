# step
使用Beego框架构建的Go语言简单博客，内容包括文章、友情链接、标签等。

模板使用为 WEB 艺术家创造的PHP框架[Laravel官方网站](http://www.golaravel.com)首页样式。

文章原文均在文件夹views/stores目录中，博客使用tree命令显示如下：
```bash
[zhgxun-pro:step zhgxun$ tree
 .
 ├── LICENSE
 ├── README.md
 ├── conf
 │   └── app.conf
 ├── controllers
 │   ├── auth.go
 │   ├── banner.go
 │   ├── default.go
 │   ├── label.go
 │   ├── link.go
 │   ├── post.go
 │   ├── quote.go
 │   └── user.go
 ├── main.go
 ├── models
 │   ├── banner.go
 │   ├── label.go
 │   ├── link.go
 │   ├── menu.go
 │   ├── post.go
 │   ├── post_label.go
 │   ├── quote.go
 │   └── user.go
 ├── routers
 │   └── router.go
 ├── static
 │   ├── css
 │   │   ├── common.css
 │   │   ├── markdown.css
 │   │   └── screen.css
 │   ├── img
 │   │   └── favicon.ico
 │   └── js
 │       └── markdown.js
 ├── step
 ├── templates
 │   └── func.go
 └── views
     ├── base.html
     ├── index.html
     ├── label
     │   ├── create.html
     │   ├── from.html
     │   ├── index.html
     │   ├── update.html
     │   └── view.html
     ├── layout.html
     ├── link
     │   ├── create.html
     │   ├── from.html
     │   ├── index.html
     │   ├── update.html
     │   └── view.html
     ├── login.html
     ├── post
     │   ├── create.html
     │   ├── from.html
     │   ├── index.html
     │   ├── update.html
     │   └── view.html
     ├── quote
     │   ├── create.html
     │   ├── from.html
     │   ├── index.html
     │   ├── update.html
     │   └── view.html
     ├── stores
     │   ├── programming
     │   │   └── 2017
     │   │       ├── beego_link.html
     │   │       └── beego_set_cookie.html
     │   ├── reading
     │   │   └── 2017
     │   │       ├── a_friend.html
     │   │       └── the_go_programming_language.html
     │   └── travel
     │       └── 2017
     │           └── haituo.html
     └── user
         ├── create.html
         ├── from.html
         ├── index.html
         ├── update.html
         └── view.html

 22 directories, 62 files
 zhgxun-pro:step zhgxun$
```

由MacDown软件编写后转化为html格式后再次选择而成，博客中已经使用了MacDown生成Html的样式表和JavaScript文件。数据库中直接存储该源文件路径，即为显示的TplName模板名。
