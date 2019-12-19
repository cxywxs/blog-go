package dao

import (
	"blog-go/mapper"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"log"
)

//连接数据库
func UserDao() *gorm.DB {
	application := mapper.AnalysisApplication()
	db, err := gorm.Open("mysql", application.Databases.Root+":"+application.Databases.Password+"@tcp("+application.Databases.Server+":"+application.Databases.Port+")/"+application.Databases.Database+"?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	log.Println("mysql connect succeed!")
	db.SingularTable(true)
	return db
}
