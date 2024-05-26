package config

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

const (
	host     = "localhost"
	port     = 3307
	dbname   = "upload"
	username = "root"
	password = ""
)

var dsn = fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

const (
    colorGreen = "\033[1;32m"
    colorRed   = "\033[1;31m"
    colorReset = "\033[0m"
)

func DB() *gorm.DB {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        fmt.Printf("\n----------------\n|%sNOT CONNECTED!%s|\n----------------\n", colorRed, colorReset)
        panic(err)
    }
	printGraffiti()
    return db
}

func printGraffiti() {
	graffiti := `
  
  ___  ____   ___   ___  _   _ _   _ _____ ____ _____ 
 |_ _|/ ___| / __| / _ \| \ | | \ | | ____/ ___|_   _|
  | | \___ \| /   | | | |  \| |  \| |  _|| /     | |  
  | |  ___) | \__ | |_| | |\  | |\  | |__| \___  | |  
 |___||____/ \ __| \___/|_| \_|_| \_|_____|____| |_|   
  ____________________________________________________
 \_\___\___\___\___\___\___\___\______________________\
                                                      `

	fmt.Printf("%s%s%s\n", colorGreen, graffiti, colorReset)
}