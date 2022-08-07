package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GenPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	}

	return true
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *gorm.DB

func init() {
	dsn := "root:@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{})
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	b, _ := GenPwd(password)

	user := User{
		Username: username,
		Password: string(b),
	}
	db.Create(&user)
	c.Redirect(301, "/")
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Printf("表单密码： %v\n", password)

	var u User
	db.Where("username=?", username).First(&u)

	if u.Username == "" {
		c.HTML(200, "login.html", "用户名不存在！")
		fmt.Println("用户名不存在！")
	} else {
		if ComparePwd(u.Password, password) {
			fmt.Println("登录成功")
			c.Redirect(301, "/")
		} else {
			fmt.Println("密码错误")
			c.HTML(200, "login.html", "密码错误")
		}
	}
}
func main() {
	e := gin.Default()
	e.LoadHTMLGlob("*.html")

	e.GET("/login", GoLogin)
	e.POST("login", Login)

	e.GET("/", Index)
	e.POST("register", Register)
	e.GET("register", GoRegister)
	e.Run()
}
