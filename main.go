package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySql() (err error) {
	dsn := "root:qwer1234@tcp(127.0.0.1:8993)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	// 连接数据库
	err := initMySql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	DB.AutoMigrate(&Todo{})

	engine := gin.Default()

	engine.LoadHTMLGlob("templates/*")
	engine.Static("/static", "static")

	engine.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := engine.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			c.BindJSON(&todo)
			err := DB.Create(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "success",
				"data": todo,
			})
		})

		// 查看所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			err := DB.Find(&todoList).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
				return
			}
			c.JSON(http.StatusOK, todoList)
		})

		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id := c.Params.ByName("id")
			var todo Todo
			err := DB.Where("id=?", id).First(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
				return
			}
			c.BindJSON(&todo)
			err = DB.Save(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
				return
			}
			c.JSON(http.StatusOK, todo)
		})

		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id := c.Params.ByName("id")
			var todo Todo
			err := DB.Where("id=?", id).First(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
				return
			}
			err = DB.Delete(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
				return
			}
			c.JSON(http.StatusOK, &todo)
		})
	}

	engine.Run()
}
