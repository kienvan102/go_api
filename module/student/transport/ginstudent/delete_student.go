package ginstudent

import (
	studentbiz "first-apis/module/student/biz"
	studentstorage "first-apis/module/student/storage"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func DeleteStudentByID(db *gorm.DB) func(c *gin.Context){
	return func(c *gin.Context){
		
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		
		store := studentstorage.NewSqlStore(db)
		biz := studentbiz.NewDeleteStudentBiz(store)
		if err := biz.DeleteStudent(c.Request.Context(), id); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data" : "success",
		})
	}
}