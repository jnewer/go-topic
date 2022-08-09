package main

import (
	"encoding/json"
	// "encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species")) //blue gopher

	type User struct {
		Name      string    `json:"name"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"createdAt"`
	}

	u := &User{
		Name:      "kang",
		Password:  "666",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
	// {
	// 	"name": "kang",
	// 	"password": "666",
	// 	"createdAt": "2022-08-10T06:45:25.8077985+08:00"
	// }

	// type Address struct {
	// 	City, State string
	// }

	// type Person struct {
	// 	XMLName   xml.Name `xml:"person"`
	// 	Id        int      `xml:"id,attr"`
	// 	FirstName string   `xml:"name>first"`
	// 	LastName  string   `xml:"name>last"`
	// 	Age       int      `xml:"age"`
	// 	Height    float32  `xml:"height,omitempty"`
	// 	Married   bool
	// 	Address
	// 	Comment string `xml:",comment"`
	// }

	// v := &Person{
	// 	Id:        13,
	// 	FirstName: "Kong",
	// 	LastName:  "Kong",
	// 	Age:       20,
	// }

	// v.Comment = "comment"
	// v.Address = Address{"北京", "北京"}
	// output, err := xml.MarshalIndent(v, "", "  ")
	// if err != nil {
	// 	fmt.Printf("error: %v", err)
	// }

	// os.Stdout.Write(output)
	// 	<person id="13">
	//   <name>
	//     <first>Kong</first>
	//     <last>Kong</last>
	//   </name>
	//   <age>20</age>
	//   <Married>false</Married>
	//   <City>北京</City>
	//   <State>北京</State>
	//   <!--comment-->
	//  </person>

	// type Channel struct {
	// 	Id      uint64 `form:"id" gorm:"primaryKey"`
	// 	Title   string `form:"title" gorm:"title"`
	// 	Slug    string `form:"slug" gorm:"slug"`
	// 	Content string `form:"content" gorm:"content"`
	// 	Status  int    `form:"status" gorm:"status"`
	// 	Weight  int    `form:"weight" gorm:"weight"`
	// }

	router := gin.Default()

	router.POST("loginJSON", func(ctx *gin.Context) {
		var json Login
		if err := ctx.ShouldBindJSON(&json); err == nil {
			if json.User == "kong" && json.Password == "123" {
				ctx.JSON(http.StatusOK, gin.H{
					"status": "登录成功",
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"status": "未授权",
				})
			}
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	router.POST("/loginForm", func(ctx *gin.Context) {
		var form Login
		if err := ctx.ShouldBind(&form); err == nil {
			if form.User == "kong" && form.Password == "123" {
				ctx.JSON(http.StatusOK, gin.H{
					"status": "登录成功",
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"status": "未授权",
				})
			}
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	router.Run()
	// curl -H "Content-Type:application/json" -X POST -d '{"user":"kong","password":"123"}' http://localhost:8080/loginJSON
	// curl -X POST -d "user=kong&password=123" http://localhost:8080/loginForm

}
