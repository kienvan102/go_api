package ginstudent

import (
	studentbiz "first-apis/module/student/biz"
	studentmodel "first-apis/module/student/model"
	studentstorage "first-apis/module/student/storage"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateStudent(db *gorm.DB) func(c *gin.Context){
	return func(c *gin.Context){
		
		var data studentmodel.Student
		if err := c.ShouldBind(&data); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			
			return
		}

		// db.Create(&data)
		store := studentstorage.NewSqlStore(db)
		biz := studentbiz.NewCreateStudentBiz(store)
		if err := biz.CreateStudent(c.Request.Context(), &data); err!= nil{
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
