package main

import (
	"bipubipu/service"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type LoginJSON struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Shop struct {
	Id           string `json:"id" form:"id"`
	Name         string `json:"name" form:"name"`
	SellerId     string `json:"sellerId" form:"seller_id"`
	ImgUrl       string `json:"imgUrl" form:"img_url"`
	RecommandStr string `json:"recommandStr" form:"recommand_str"`
	Carousel     string `json:"carousel" form:"carousel_str"`
}

func main() {

	name, num := service.ShopService(435)
	fmt.Println(name, num)
	
	gin.SetMode(gin.DebugMode)
	
	r := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))



	// ------------------------------------------------
	
	r.GET("/volume", func(c *gin.Context) {
		// TODO 调用service
		volumes := service.GetVolumeAll()
		
		c.JSON(200, gin.H{"volumes": volumes})
	})

	// ------------------------------------------------




	r.GET("/ping", func(c *gin.Context) {

		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		// 初始化session里的键
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count += 1
		}
		session.Set("count", count)
		session.Save()

		c.String(200, "pong"+strconv.Itoa(count))
	})
	
	r.GET("/ping/:test", func(c *gin.Context) {
		test := c.Params.ByName("test")
		message := "name1:" + test
		c.String(200, message)
	})

	r.GET("/user/:name1/two/:name2", func(c *gin.Context) {
		name1 := c.Params.ByName("name1")
		name2 := c.Params.ByName("name2")
		message := "name1:" + name1 + ", name2:" + name2
		c.String(200, message)
	})

	r.GET("/db", func(c *gin.Context) {

		shops := make([]Shop, 0)

		db := getDB()
		result, _ := db.Query("SELECT * FROM tab_shop")
		defer result.Close()

		for result.Next() {
			var shop Shop
			result.Scan(&shop.Id, &shop.Name, &shop.SellerId, &shop.ImgUrl, &shop.RecommandStr, &shop.Carousel)
			shops = append(shops, shop)
		}
		c.JSON(200, gin.H{"shops": shops})
	})

	r.POST("/login", func(c *gin.Context) {
		var json LoginJSON

		c.Bind(&json) // This will infer what binder to use depending on the content-type header.

		if json.User == "manu" && json.Password == "123" {
			c.JSON(200, gin.H{"status": "you are logged in", "name": json.User})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized", "name": json.User})
		}
	})

	r.Run(":8080")
}

func getDB() (db *sql.DB) {
	db, _ = sql.Open("mysql", "rache:21431@tcp(localhost:3306)/bipubipu?collation=utf8_general_ci&parseTime=true")
	return
}
