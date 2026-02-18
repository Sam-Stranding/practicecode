package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

var students = []student{
	{1, "张三", 18, 1},
	{2, "李四", 19, 2},
	{3, "王五", 20, 3},
	{4, "赵六", 21, 4},
}

func getStudent(c *gin.Context) {
	idStr := c.Query("id")
	stdID, _ := strconv.Atoi(idStr) //将 idStr转换为int类型

	for _, std := range students { //模拟
		if std.ID == stdID {
			c.JSON(200, gin.H{
				"code": 200,
				"data": std,
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"code": 404,
		"data": "未找到该学生",
	})
}

func addStudent(c *gin.Context) {
	var newStudent student
	if err := c.BindJSON(&newStudent); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"data": "参数错误",
		})
		return
	}
	students = append(students, newStudent)
	c.JSON(200, gin.H{
		"code": 200,
		"data": "添加成功",
	})
}

func updateStudent(c *gin.Context) {
	var newStudent student
	if err := c.BindJSON(&newStudent); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"data": "参数错误",
		})
		return
	}
	for i, std := range students {
		if std.ID == newStudent.ID {
			students[i] = newStudent
			c.JSON(200, gin.H{
				"code": 200,
				"data": "更新成功",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"code:": 404,
		"data":  "未找到该学生",
	})
}
func ginRouterDemo() {
	router := gin.Default()

	apiV1 := router.Group("/api/v1") // http://localhost:8080/api/v1
	{
		stuGroup := apiV1.Group("/students") // http://localhost:8080/api/v1/students
		{
			stuGroup.GET("/info", getStudent)     // /api/v1/students/info?id=1
			stuGroup.POST("add", addStudent)      // /api/v1/students/add
			stuGroup.PUT("update", updateStudent) // /api/v1/students/update
		}
	}

	fmt.Println("Gin路由启动，8080")
	router.Run(":8080")
}

func main() {
	ginRouterDemo()
}
