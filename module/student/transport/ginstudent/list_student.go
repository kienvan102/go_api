package ginstudent



import (
	studentmodel "first-apis/module/student/model"
	studentbiz "first-apis/module/student/biz"
	studentstorage "first-apis/module/student/storage"
	// "strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func ListStudent(db *gorm.DB) func(c *gin.Context){
	return func(c *gin.Context){
		
		var paging studentmodel.Paging

		if err := c.ShouldBind(&paging); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		var data []studentmodel.Student

		store := studentstorage.NewSqlStore(db)
		biz := studentbiz.NewListStudentBiz(store)
		data, err := biz.ListStudent(c.Request.Context(), paging);
		if  err != nil{
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
