package db

import (
	"beego_demo/models"
	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic("failed to connect database driver")
	}
	if err = orm.RegisterDataBase("default", "mysql", "root:rootroot@tcp(127.0.0.1:3306)/beego_demo?charset=utf8", 30); err != nil {
		panic("failed to connect database")
	}
	//if err = orm.RunSyncdb("default", false, true); err != nil {
	//	panic("failed to sync database")
	//}
	orm.DefaultTimeLoc = time.UTC
	orm.Debug = true

	orm.RegisterModel(new(models.Book))
}
