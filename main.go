package main

import (
	// "fmt"
	_ "app/routers"
	// _ "app/models"
	"github.com/astaxie/beego"
	// "app/db"
	// "github.com/jinzhu/gorm"
)

// func init() {
//   _ = godotenv.Load()
// 	// db.Connect()
//   // defer db.Close()
//   beego.Run()
// }

func main() {
	beego.Run()
}

// func main() {
//   conn := getDBConnection()
//   defer conn.Close()
// }

// func getDBConnection() *gorm.DB {
//   dbHost := beego.AppConfig.String("dbHost")
//   dbUser := beego.AppConfig.String("dbUser")
//   dbName := beego.AppConfig.String("dbName")
//   dbPass := beego.AppConfig.String("dbPass")
//   conn, err := db.Connect(dbHost, dbName, dbUser, dbPass)
//   if err != nil {
//     panic(err.Error())
//   }
//   db.AutoMigrate()
//   return conn
// }
