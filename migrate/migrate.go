package main

import (
	"api-go/db"
	"api-go/model"
	"fmt"
)

func main() {
	dbConn:=db.NewDB()
	defer fmt.Println("successfully migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{},&model.Task{})
}