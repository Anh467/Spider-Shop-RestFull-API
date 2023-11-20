package main

import (
	common "SpiderShop-Restfull-API/common"
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const FILE_CONTEXT = "config.json"

func main() {
	// declare
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
	if err := ConnectSqlServerGorm(appctx.MySQL, &appctx.GormDB); err != nil {
		fmt.Println("Connect to mysql failed " + err.Error())
		return
	}
	fmt.Println("Connect to mysql successfully ")
}

func ConnectSqlServerGorm(MySQL common.MySQL, db *gorm.DB) error {
	// create connection string
	connString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MySQL.User, MySQL.Pass, MySQL.DB)
	var err error
	// connect to database
	if db, err = gorm.Open(mysql.Open(connString), &gorm.Config{}); err != nil {
		// getting error and handle it
		fmt.Println("Connecting my sql server failure")
		return err
	}
	// connect to database successfully and print success
	fmt.Println("Connecting my sql server successfully")
	return nil
}
