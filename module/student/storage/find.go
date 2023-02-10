package studentstorage

import (
	"context"
	studentmodel "first-apis/module/student/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, 
	condition map[string]interface{},
	moreKeys ...string,
	) (*studentmodel.Student, error){
		
		var data studentmodel.Student
		if err := s.db.Where(condition).First(&data).Error; err != nil{
			return nil,err
		}
	return &data,nil
}