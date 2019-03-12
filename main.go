package main

import (
	. "MPPL-Modul-2/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"net/url"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	//dbHost := viper.GetString(`database.host`)
	//dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPassword := viper.GetString(`database.password`)
	dbName := viper.GetString(`database.name`)

	//connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	//val := url.Values{}
	//val.Add("parseTime", "1")
	//val.Add("loc", "Asia/Jakarta")
	//dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	//dbConn, err := sql.Open(`mysql`, dsn)

	connection := fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName)
	val := url.Values{}
	val.Add("charset", "utf8")
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
	}

	defer dbConn.Close()

	dbConn.AutoMigrate(&Product{}, &Brand{}, &Category{})



}