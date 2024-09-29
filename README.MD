# 配置

```表结构生成模型
# 控制生成内容的级别，1 表示只生成模型，2 表示生成模型和控制器，3 表示生成模型、控制器和路由。
bee generate appcode -tables=表名 -driver=mysql -conn="root:rootroot@tcp(127.0.0.1:3306)/beego_demo" -level=1
```