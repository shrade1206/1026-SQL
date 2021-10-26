package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gocv.io/x/gocv"
)

type Pic struct {
	Id      int `gorm:"primary_key"`
	Name    string
	Picture []byte `gorm:"type:blob"`
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	fmt.Println("Go WebSocket")
	mySql()
	r := gin.New()
	r.POST("/create", func(c *gin.Context) {
		fmt.Println("Create Data to DB")
		c.JSON(200, "OK")
	})
	r.GET("/ws", wsEndpoint)
	r.GET("/select", selectName)
	// server.GET("/cam", selectAll)
	// server.GET("/cam/:id", selectOne)
	// server.POST("/cam/insert", insert)
	// server.PUT("/cam/update/:id", update)
	// server.DELETE("/cam/:id", del)

	r.NoRoute(gin.WrapH(http.FileServer(http.Dir("./public"))), func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		fmt.Println(path)
		fmt.Println(method)
		//檢查path的開頭使是否為"/"
		if strings.HasPrefix(path, "/") {
			fmt.Println("ok")
		}
	})
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("8080 err : ", err.Error())
	}
}

var (
	db  *gorm.DB
	err error
	pic Pic
)

func mySql() {
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/camdb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Failed To Connect : ", err.Error())
	}
	db.LogMode(true)
	db.AutoMigrate(&Pic{})
}

// func selectAll(c *gin.Context) {
// 	err = db.Find(&pic).Error
// 	c.JSON(http.StatusOK, &pic)
// }
// func selectOne(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	err := db.First(&pic, id).Error
// 	if err != nil {
// 		c.AbortWithStatus(404)
// 		log.Printf("查詢失敗 : %s", err.Error())
// 	} else {
// 		c.JSON(http.StatusOK, &pic)
// 	}
// }

// func insert(c *gin.Context) {

// }
// func update(c *gin.Context) {

// }
// func del(c *gin.Context) {

// }

func selectName(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket 連線異常 : %s", err.Error())
		return
	}
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		var Aa Pic
		db.Where("Name = ? ", string(p)).Find(&Aa)
		log.Println("訊息 :" + string(p))
		log.Println(Aa)
		// if err := ws.WriteMessage(messageType, []byte()); err != nil {
		// 	log.Println(err)
		// 	return
		// }

	}
}

func wsEndpoint(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("使用者已連線")

	var newImg []byte
	var data string
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		if string(p) == "run" {
			func() {
				webcam, err := gocv.VideoCaptureDevice(0)
				if err != nil {
					log.Println(err)
				}
				time.Sleep(time.Second)
				img := gocv.NewMat()
				defer img.Close()

				webcam.Read(&img)
				defer webcam.Close()
				buf, err := gocv.IMEncode(".jpg", img)
				if err != nil {
					log.Fatal(err)
				}
				defer buf.Close() //nolint
				newImg = buf.GetBytes()
				// d, _ := os.ReadFile(a)
				data = base64.StdEncoding.EncodeToString(newImg)
				if err := ws.WriteMessage(messageType, []byte(data)); err != nil {
					log.Println(err)
					return
				}
			}()
		}
		if string(p) == "save" {
			a := []byte(data)
			pic = Pic{Name: "GGGG", Picture: a}
			result := db.Debug().Create(&pic)
			if result.Error != nil {
				log.Printf("Data Error : %s", result.Error)
			}
		}
		log.Println("使用者訊息: " + string(p))

		log.Println(string(p)[1])
		log.Println(string(p)[0])
	}
}
