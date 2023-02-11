package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"first-apis/module/student/transport/ginstudent"
)
func getDBString() string{
	err := godotenv.Load(".env")

	if err != nil{
		log.Fatalf("Error loading .env file")
	}

	db_user := os.Getenv("DB_USER")
	db_pwd := os.Getenv("DB_PWD")
	db_Add := os.Getenv("DB_ADDRESS")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	dbStr := db_user + ":" + db_pwd + "@tcp(" + db_Add + ":" + db_port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dbStr
  }




func main() {

	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = getDBString();
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println(db,err)
	
	r := gin.Default()

	v1 := r.Group("/v1")
	students := v1.Group("/students")

	/*  [POST]  create a student /v1/students */
	students.POST("/", ginstudent.CreateStudent(db))

	/*[DELETE] Delete by id */
	students.DELETE("/:id",ginstudent.DeleteStudentByID(db) )

	/* GET one /v1/students /:id */
	students.GET("/:id", ginstudent.GetStudentByID(db))

	/*[PATH] Update a student */
	students.PATCH("/:id", ginstudent.UpdateStudent(db))

	// /* GET list */
	// students.GET("/", func(c *gin.Context){
		
	// 	type Paging struct{
	// 		Page int `json:"page" form:"page"`
	// 		Limit int `json:"limit" form:"limit"`
	// 	}
	// 	var data []Student
	// 	var pagingData Paging
		
	// 	if err := c.ShouldBind(&pagingData); err != nil{
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	if pagingData.Page <= 0{
	// 		pagingData.Page = 1
	// 	}

	// 	if pagingData.Limit <= 0{
	// 		pagingData.Page = 5
	// 	}

	// 	db.
	// 	Offset((pagingData.Page - 1)*pagingData.Limit).
	// 	Order("id desc").
	// 	Limit(pagingData.Limit).
	// 	Find(&data)

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data" : data,
	// 	})

	// })

	r.Run()
}

  