package main

import (
	common "SpiderShop-Restfull-API/common"
	_ "SpiderShop-Restfull-API/docs"
	"SpiderShop-Restfull-API/router"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const FILE_CONTEXT = "config.json"

// @title SpiderShop-Restfull-API yeah
// @version 1.0
// @description Spider shop Restfull API using Gin Framework, Gorm orm and other libraries
// @host localhost:8080
// @BasePath /api/v1

func main() {
	// declare
	var r *gin.Engine
	var appctx common.AppConext
	var file *os.File
	var content string
	var err error
	// get file context
	if file, err = common.GetFile(FILE_CONTEXT); err != nil {
		fmt.Println("Read file failed " + err.Error())
		return
	}
	// get content in file context
	if content, err = common.GetFileContent(file); err != nil {
		fmt.Println("Get content in file context failed " + err.Error())
		return
	}
	// convert json to struct
	if err := json.Unmarshal([]byte(content), &appctx); err != nil {
		fmt.Println("Convert json to struct failed " + err.Error())
		return
	}
	// connect to mysql
	if appctx.GormDB, err = ConnectSqlServerGorm(appctx.MySQL); err != nil {
		fmt.Println("Connect to mysql failed " + err.Error())
		return
	}
	fmt.Println("Connect to mysql successfully ")
	// setting gin
	gin.SetMode(gin.ReleaseMode)
	// init gin
	r = gin.Default()

	// setting
	r.Use(cors.New(cors.Config{
		AllowHeaders:     appctx.AllowHeaders,
		AllowAllOrigins:  appctx.AllowAllOrigins,
		AllowMethods:     appctx.AllowMethods,
		AllowCredentials: appctx.AllowCredentials,
	}))
	//
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})
	// middleware recover
	// router
	router.V1Router(r, &appctx)
	// run
	r.Run()
}

func ConnectSqlServerGorm(MySQL common.MySQL) (*gorm.DB, error) {
	// create connection string
	connString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MySQL.User, MySQL.Pass, MySQL.DB)
	//declare
	var db *gorm.DB
	var err error
	// connect to database
	if db, err = gorm.Open(mysql.Open(connString), &gorm.Config{}); err != nil {
		// getting error and handle it
		fmt.Println("Connecting my sql server failure")
		return nil, err
	}
	// connect to database successfully and print success
	fmt.Println("Connecting my sql server successfully")
	return db, nil
}
