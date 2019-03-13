package main

import (
	_ "github.com/go-sql-driver/mysql"

	"MPPL-Modul-2/manage_product/delivery/http"
	"MPPL-Modul-2/manage_product/repository"
	"MPPL-Modul-2/manage_product/usecase"
	. "MPPL-Modul-2/models/manage_product"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
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
	dbUser := viper.GetString(`database.user`)
	dbPassword := viper.GetString(`database.password`)
	dbName := viper.GetString(`database.name`)

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

	dbConn.AutoMigrate(&Product{}, &Brand{}, &Category{}, &Color{}, &Image{}, &Review{}, &Size{}, &Variant{})

	br := repository.NewBrandRepository(dbConn)
	ctr := repository.NewCategoryRepository(dbConn)
	clr := repository.NewColorRepository(dbConn)
	ir := repository.NewImageRepository(dbConn)
	pr := repository.NewProductRepository(dbConn)
	rr := repository.NewReviewRepository(dbConn)
	sr := repository.NewSizeRepository(dbConn)
	vr := repository.NewVariantRepository(dbConn)

	pu := usecase.NewProductUsecase(br, ctr, clr, ir, pr, rr, sr, vr)

	e := echo.New()
	http.NewBrandHandler(e, pu)
	http.NewCategoryHandler(e, pu)
	http.NewColorHandler(e, pu)
	http.NewImageHandler(e, pu)
	http.NewProductHandler(e, pu)
	http.NewReviewHandler(e, pu)
	http.NewSizeHandler(e, pu)
	http.NewVariantHandler(e, pu)

	_ = e.Start(viper.GetString("server.address"))
}