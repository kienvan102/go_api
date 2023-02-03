package main

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	// "encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	fmt.Println(dbStr)
	return dbStr
  }


type Student struct{
	Id int `json:"id" gorm:"column:id"`
	FullName string `json:"fullname" gorm:"column:fname"`
	Sex string `json:"sex" gorm:"column:sex"`
	Phone int `json:"phone" gorm:"column:phone"`
	Email string `json:"email" gorm:"column:email"`
}

type StudentUpdate struct{
	Phone *int `json:"phone" gorm:"column:phone"`
	Email *string `json:"email" gorm:"column:email"`
}

func (Student) TableName() string { return "students" }
func (StudentUpdate) TableName() string { return Student{}.TableName()}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

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
	students.POST("/", func(c *gin.Context){
		
		var data Student
		if err := c.ShouldBind(&data); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}

		db.Create(&data)

		c.JSON(http.StatusOK, gin.H{
			"data" : data,
		})
	})

	/* GET one /v1/students /:id */
	students.GET("/:id", func(c *gin.Context){
		id, err := strconv.Atoi(c.Param("id"))
		
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		var data Student
		db.Where("id=?", id).First(&data)
		c.JSON(http.StatusOK, gin.H{
			"data" : data,
		})

	})

	/* GET list */
	students.GET("/", func(c *gin.Context){
		
		type Paging struct{
			Page int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}
		var data []Student
		var pagingData Paging
		
		if err := c.ShouldBind(&pagingData); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if pagingData.Page <= 0{
			pagingData.Page = 1
		}

		if pagingData.Limit <= 0{
			pagingData.Page = 5
		}

		db.
		Offset((pagingData.Page - 1)*pagingData.Limit).
		Order("id desc").
		Limit(pagingData.Limit).
		Find(&data)

		c.JSON(http.StatusOK, gin.H{
			"data" : data,
		})

	})

	/*[PATH] Update a student */
	students.PATCH("/:id", func(c *gin.Context){
		id, err := strconv.Atoi(c.Param("id"))
		
		
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		var data StudentUpdate
		if err := c.ShouldBind(&data); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := db.Where("id=?", id).Updates(&data).Error; err != nil{
			log.Println(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"data" : data,
		})

	})

	/*[DELETE] Delete by id */
	students.DELETE("/:id", func(c *gin.Context){
		id, err := strconv.Atoi(c.Param("id"))
		
		
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		
		if err := db.Table(Student{}.TableName()).Where("id=?", id).Delete(nil).Error; err != nil{
			log.Println(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"data" : "success",
		})

	})

	r.Run()
}

  