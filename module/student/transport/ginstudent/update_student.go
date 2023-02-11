package ginstudent

import (
	studentbiz "first-apis/module/student/biz"
	studentmodel "first-apis/module/student/model"
	studentstorage "first-apis/module/student/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func UpdateStudent(db *gorm.DB) func(c *gin.Context){
	return func(c *gin.Context){
		
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		var data studentmodel.StudentUpdate
		if err := c.ShouldBind(&data); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		store := studentstorage.NewSqlStore(db)
		biz := studentbiz.NewUpdateStudentBiz(store)
		if err := biz.UpdateStudent(c.Request.Context(), id, &data); err!= nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
		})

		return
	}

		c.JSON(http.StatusOK, gin.H{
			"data" : data,
		})
	}
}