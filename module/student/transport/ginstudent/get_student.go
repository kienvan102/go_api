package ginstudent

import (
	studentbiz "first-apis/module/student/biz"
	studentstorage "first-apis/module/student/storage"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func GetStudentByID(db *gorm.DB) func(c *gin.Context){
	return func(c *gin.Context){
		
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		
		store := studentstorage.NewSqlStore(db)
		biz := studentbiz.NewGetStudentBiz(store)
		student,err := biz.GetStudent(c.Request.Context(), id)
		if  err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data" : student,
		})
	}
}